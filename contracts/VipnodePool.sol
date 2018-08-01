pragma solidity ^0.4.24;

// VipnodePool manages the payment component of a Vipnode Pool. (https://vipnode.org)
//
// The payment mechanism can evolve in the future, but the current version holds
// a maximum balance for clients, and the pool operator is trusted to draw from
// that balance periodically based on internal accounting.
//
// Future improvements could use state channels to do micropayments directly between
// the pool host and client, bypassing the trusted pool operator.
//
// The Client is managed by a wallet address, but the client whitelists any
// number of NodeIDs (public keys).
contract VipnodePool {

  // XXX: This is a work in progress. It is not fully tested yet.

  // TODO: Should host node balances be also managed on-chain?
  // TODO: Should the operator take a fee? Should the txn fee be taken from the
  // client balance?
  // TODO: Should the nodeID whitelist be handled entirely by the pool
  // offchain, and this should be a more generic timelock thing?

  // ForceSettle is emitted when a client requests the pool operator to settle
  // the balance and unlock it for withdrawing.
  event ForceSettle(
    address client,
    uint256 timeLocked
  );

  // Balance is emitted whenever a client balance is updated.
  event Balance(
    address client,
    uint256 balance
  );

  // Client represents a Pool client who pays to connect to a Pool host.
  struct Client {
    bytes32[] nodeIDs; // Whitelisted nodes by Client
    uint256 balance; // Client's balance
    uint256 timeLocked; // If not 0, then client will be allowed to withdraw balance after this time.
  }

  // operator is the pool operator who has permission to manage the client balances.
  address public operator;

  // Amount of time the pool manager has to settle the balance before the
  // client can force a balance withdraw.
  uint256 public withdrawInterval = 7 days;

  // Mapping of owner -> client.
  mapping (address => Client) public clients;

  constructor(address _operator) public {
    require(_operator != address(0));

    operator = _operator;
  }

  // [Pure] checkBalance confirms that a nodeID exists and there is a minimum
  // balance available for the client.
  function checkBalance(address _client, bytes32 _nodeID, uint256 _minBalance) public view returns (bool) {
    Client memory c = clients[_client];
    if (c.balance < _minBalance) return false;
    for(uint i=0; i<c.nodeIDs.length; i++) {
      if (c.nodeIDs[i] == _nodeID) return true;
    }
    return false;
  }

  // addBalance adds the msg.sender as a Client and adds the nodeID to the node
  // whitelist. If nodeID is 0, whitelisting is skipped.
  function addBalance(bytes32 _nodeID) public payable {
    Client storage c = clients[msg.sender];
    c.balance = c.balance + msg.value;

    if (_nodeID != bytes32(0)) {
      c.nodeIDs.push(_nodeID);
    }

    emit Balance(msg.sender, c.balance);
  }

  // forceSettle emits a ForceSettle event for the msg.sender if it's a valid
  // client, this prompts the pool operator to initiate balance settlement
  // and release the balance funds.
  // The client's timeLocked value will initialize to +withdrawInterval, after
  // which the client will be allowed to forceWithdraw.
  // Requires the client's balance to be >0.
  function forceSettle() public {
    Client storage c = clients[msg.sender];
    require(c.balance > 0);

    c.timeLocked = withdrawInterval + block.timestamp;

    emit ForceSettle(msg.sender, c.timeLocked);
  }

  // forceWithdraw allows the msg.sender to withdraw any available client
  // balance iff timeLocked is set and older than now and balance > 0.
  function forceWithdraw() public {
    Client storage c = clients[msg.sender];
    require(c.balance > 0);
    require(c.timeLocked > 0);
    require(c.timeLocked <= block.timestamp);

    // Reset client
    uint256 amount = c.balance;
    c.balance = 0;
    c.timeLocked = 0;

    msg.sender.transfer(amount);
    emit Balance(msg.sender, c.balance);
  }


  /*
  // removeClientNode removes nodeID from the msg.sender's client's node
  // whitelist.
  function removeClientNode(bytes32 _nodeID) {
    // XXX: TODO
  }
  */

  // -------------------
  // Operator Functions:
  // -------------------

  // Only callable by operator: Deduct _amount from _client's balance and send
  // it to _operator.
  // If _release, then send the remaining _balance to the _client and set
  // _balance to 0. To be called when ForceSettle is initiated.
  function opSettle(address _client, uint256 _amount, bool _release) public {
    require(msg.sender == operator);
    // FIXME: Confirm that there is no race conflict with addBalance.

    Client storage c = clients[_client];
    require(c.balance >= _amount);
    c.balance = c.balance - _amount;

    if (_release) {
      uint256 amount = c.balance;
      c.balance = 0;
      c.timeLocked = 0;

      _client.transfer(amount);
    }

    emit Balance(_client, c.balance);
  }

  // Only callable by operator: Withdraw some balance, this ability is required
  // to pay the hosts in a trusted manner. Would be nice for this contract to
  // be trustless, but that would require 2*N*M on-chain transactions to do
  // naively. A future version will improve upon this.
  function opWithdraw(uint256 _amount) public {
    require(msg.sender == operator);

    operator.transfer(_amount);
  }
}
