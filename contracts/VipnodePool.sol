pragma solidity ^0.4.24;

// VipnodePool manages the payment component of a Vipnode Pool. (https://vipnode.org)
//
// The payment mechanism can evolve in the future, but the current version holds
// a maximum balance for accounts, and the pool operator is trusted to settle
// that balance periodically based on internal accounting.
//
// Pool operators may impose their own fees during settlement.
//
// Future improvements ideas are outlined at: https://github.com/vipnode/vipnode-contract
contract VipnodePool {

  // ForceSettle is emitted when a account requests the pool operator to settle
  // the balance and unlock it for withdrawing.
  event ForceSettle(
    address account,
    uint256 timeLocked
  );

  // Balance is emitted whenever a account balance is updated.
  event Balance(
    address account,
    uint256 balance
  );

  // Account represents a participant in the pool. It can be a client or a host.
  struct Account {
    uint256 balance; // Account's balance
    uint256 timeLocked; // If not 0, then account will be allowed to withdraw balance after this time.
  }

  // operator is the pool operator who has permission to manage the account balances.
  address public operator;

  // Amount of time the pool manager has to settle the balance before the
  // account can force a balance withdraw.
  uint256 public withdrawInterval = 7 days;

  // Mapping of owner -> account.
  mapping (address => Account) public accounts;

  constructor(address _operator) public {
    require(_operator != address(0));

    operator = _operator;
  }

  // ----------
  // Interface:
  // ----------

  // addBalance adds the msg.sender as a Account.
  function addBalance() public payable {
    Account storage a = accounts[msg.sender];
    a.balance = a.balance + msg.value;

    emit Balance(msg.sender, a.balance);
  }

  // forceSettle emits a ForceSettle event for the msg.sender if it's a valid
  // account, this prompts the pool operator to initiate balance settlement
  // and release the balance funds.
  // The account's timeLocked value will initialize to +withdrawInterval, after
  // which the account will be allowed to forceWithdraw.
  // Requires the account's balance to be >0.
  function forceSettle() public {
    Account storage a = accounts[msg.sender];
    require(a.balance > 0);

    a.timeLocked = withdrawInterval + block.timestamp;

    emit ForceSettle(msg.sender, a.timeLocked);
  }

  // forceWithdraw allows the msg.sender to withdraw any available account
  // balance iff timeLocked is set and older than now and balance > 0.
  function forceWithdraw() public {
    Account storage a = accounts[msg.sender];
    require(a.balance > 0);
    require(a.timeLocked > 0);
    require(a.timeLocked <= block.timestamp);

    // Reset account
    uint256 amount = a.balance;
    a.balance = 0;
    a.timeLocked = 0;

    msg.sender.transfer(amount);
    emit Balance(msg.sender, a.balance);
  }

  // -------------------
  // Operator Functions:
  // -------------------

  // Only callable by operator: Reset _account's balance and timelock and
  // transfer _amount worth of funds to _account.
  //
  // This function can be used to answer a ForceSettle request, or a normal
  // off-chain withdraw request.
  //
  // Note: The account's credit total is stored off-chain, so _releaseAmount
  // transferred can differ from the balance amount in the contract.
  function opSettle(address _account, uint256 _releaseAmount, uint256 _newBalance) public {
    require(msg.sender == operator);

    Account storage a = accounts[_account];
    a.balance = _newBalance;
    a.timeLocked = 0;

    if (_releaseAmount > 0 ) {
      _account.transfer(_releaseAmount);
    }

    emit Balance(_account, a.balance);
  }

  // Only callable by operator: Withdraw some balance from the contract.
  //
  // This can be used to consume pool operational fees, or to pay hosts
  // off-contact. The pool operator is trusted to not abuse the balance of the
  // contract.
  function opWithdraw(uint256 _amount) public {
    require(msg.sender == operator);

    operator.transfer(_amount);
  }
}
