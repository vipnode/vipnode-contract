const VipnodePool = artifacts.require("VipnodePool");

// expectThrow borrowed from https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/test/helpers/expectThrow.js
async function expectThrow (promise, message, expected) {
  try {
    await promise;
  } catch (error) {
    // Message is an optional parameter here
    if (message) {
      assert(
        error.message.search(message) >= 0,
        'Expected \'' + message + '\', got \'' + error + '\' instead',
      );
      return;
    } else {
      // TODO: Check jump destination to destinguish between a throw
      //       and an actual invalid jump.
      const invalidOpcode = error.message.search('invalid opcode') >= 0;
      // TODO: When we contract A calls contract B, and B throws, instead
      //       of an 'invalid jump', we get an 'out of gas' error. How do
      //       we distinguish this from an actual out of gas event? (The
      //       ganache log actually show an 'invalid jump' event.)
      const outOfGas = error.message.search('out of gas') >= 0;
      const revert = error.message.search('revert') >= 0;
      assert(
        invalidOpcode || outOfGas || revert,
        'Expected throw, got \'' + error + '\' instead',
      );
      return;
    }
  }
  assert.ok(false, 'Expected throw not received: ' + expected);
}


contract('Vipnode Pool', async (accounts) => {

  let idx = 0;

  it("operator should be able to withdraw", async () => {
    const opWallet = accounts[idx++];
    const accountWallet = accounts[idx++];
    const instance = await VipnodePool.new(opWallet);
    const value = web3.toWei(42, "ether");

    let res = await instance.addBalance({ value: value, from: accountWallet });
    assert.equal(res.logs[0].event, "Balance");
    let evt = res.logs[0].args;
    assert.equal(evt.account, accountWallet);
    assert.equal(evt.balance, value);

    assert.equal(web3.eth.getBalance(instance.address), value);

    await expectThrow(
      instance.opWithdraw(value, { from: accountWallet }),
      false,
      "failed to throw on bad opWithdraw",
    );

    assert.equal(web3.eth.getBalance(instance.address), value);
    const startBalance = web3.eth.getBalance(opWallet);

    await instance.opWithdraw(value, { from: opWallet });
    assert.equal(web3.eth.getBalance(instance.address), 0);
    assert.equal(
      web3.eth.getBalance(opWallet).toPrecision(5),
      startBalance.plus(value).minus(res.receipt.cumulativeGasUsed).toPrecision(5), // This math is slightly off for some reason? Is the receipt a lie?
    );
  })

  it("account should be able to force settlement", async () => {
    const opWallet = accounts[0];
    const accountWallet = accounts[idx++];
    const instance = await VipnodePool.new(opWallet);
    const value = web3.toWei(42, "ether");

    function ethbalance(wallet) {
      return web3.fromWei(web3.eth.getBalance(wallet), "ether").toNumber();
    }

    assert.ok(ethbalance(accountWallet) > 90, "start: balance is close to full:" + ethbalance(accountWallet));
    await instance.addBalance({ value: value, from: accountWallet });
    assert.ok(ethbalance(accountWallet) < 90, "start: balance is not near full:" + ethbalance(accountWallet));

    await expectThrow(
      // forceWithdraw should fail because forceSettle hasn't completed.
      instance.forceWithdraw({ from: accountWallet }),
      false,
      "premature forceWithdraw failed to revert",
    );

    let account;
    account = await instance.accounts.call(accountWallet);
    assert.equal(account[1], 0);

    await instance.forceSettle({ from: accountWallet });
    account = await instance.accounts.call(accountWallet);
    assert.notEqual(account[1], 0);

    await expectThrow(
      // forceWithdraw should fail because timeLocked has not passed yet
      instance.forceWithdraw({ from: accountWallet }),
      false,
      "premature forceWithdraw after forceSettle failed to revert",
    );

    const timeJump = await instance.withdrawInterval();
    let res = await web3.currentProvider.send({
      id: new Date().getSeconds(),
      jsonrpc: "2.0",
      method: "evm_increaseTime",
      params: [timeJump.toNumber()+42],
    });
    assert.notOk(res.error!==undefined, "failed to increase time");

    assert.ok(ethbalance(accountWallet) < 90, "end: balance is not near full:" + ethbalance(accountWallet));
    try {
      await instance.forceWithdraw({ from: accountWallet });
    } catch(err) {
      assert.ok(false, "forceWithdraw reverted: " + err);
    }
    assert.ok(ethbalance(accountWallet) > 90, "end: balance is close to full:" + ethbalance(accountWallet));
  })

  it("op should be able to update balances", async () => {
    const opWallet = accounts[0];
    const accountWallet = accounts[idx++];
    const instance = await VipnodePool.new(opWallet);
    const value = web3.toWei(42, "ether");

    let res = await instance.opSettle(accountWallet, 0, value, { from: opWallet });

    assert.equal(res.logs[0].event, "Balance");
    let evt = res.logs[0].args;
    assert.equal(evt.account, accountWallet);
    assert.equal(evt.balance, value);

    // Contract balance is still 0 despite the account having credit
    assert.equal(web3.eth.getBalance(instance.address), 0);

    let account = await instance.accounts.call(accountWallet);
    let accountBalance = account[0];
    assert.equal(accountBalance, value);
  })

});
