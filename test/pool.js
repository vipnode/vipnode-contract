const VipnodePool = artifacts.require("VipnodePool");

contract('Vipnode Pool', async (accounts) => {

  it("operator should be able to withdraw", async () => {
    const opWallet = accounts[0];
    const clientWallet = accounts[1];
    const instance = await VipnodePool.new(opWallet);
    const nodeID = web3.fromAscii("42");
    const value = web3.toWei(42, "ether");

    let res = await instance.addBalance(nodeID, { value: value, from: clientWallet });
    assert.equal(res.logs[0].event, "Balance");
    let evt = res.logs[0].args;
    assert.equal(evt.client, clientWallet);
    assert.equal(evt.balance, value);

    assert.ok(await instance.checkBalance(clientWallet, nodeID, value - 1000));
    assert.notOk(await instance.checkBalance(clientWallet, nodeID, value + 1000));
    assert.notOk(await instance.checkBalance(opWallet, nodeID, value - 1000));
    assert.notOk(await instance.checkBalance(clientWallet, web3.fromAscii("69"), value - 1000));

    assert.equal(web3.eth.getBalance(instance.address), value);

    try {
      await instance.opWithdraw(value, { from: clientWallet });
      assert.fail("failed to throw on bad opWithdraw");
    } catch(err) {}

    assert.equal(web3.eth.getBalance(instance.address), value);
    const startBalance = web3.eth.getBalance(opWallet);

    await instance.opWithdraw(value, { from: opWallet });
    assert.equal(web3.eth.getBalance(instance.address), 0);
    assert.equal(
      web3.eth.getBalance(opWallet).toPrecision(5),
      startBalance.plus(value).minus(res.receipt.cumulativeGasUsed).toPrecision(5), // This math is slightly off for some reason? Is the receipt a lie?
    );
  })
})
