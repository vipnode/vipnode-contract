const VipnodePool = artifacts.require("VipnodePool");

contract('Vipnode Pool', async (accounts) => {

  it("operator should be able to withdraw", async () => {
    const opWallet = accounts[0];
    const clientWallet = accounts[1];
    const instance = await VipnodePool.new(opWallet);
    const nodeID = web3.fromAscii("42");

    let res = await instance.addBalance(nodeID, { value: 1234, from: clientWallet });
    assert.equal(res.logs[0].event, "Balance");
    let evt = res.logs[0].args;
    assert.equal(evt.client, clientWallet);
    assert.equal(evt.balance, 1234);

    assert.ok(await instance.checkBalance(clientWallet, nodeID, 1200));
    assert.notOk(await instance.checkBalance(clientWallet, nodeID, 1300));
    assert.notOk(await instance.checkBalance(opWallet, nodeID, 1200));
    assert.notOk(await instance.checkBalance(clientWallet, web3.fromAscii("69"), 1200));
  })
})
