# VipnodePool manages the payment component of a Vipnode Pool. (https:#vipnode.org)
#
# The payment mechanism can evolve in the future, but the current version holds
# a maximum balance for accounts, and the pool operator is trusted to settle
# that balance periodically based on internal accounting.
#
# Pool operators may impose their own fees during settlement.
#
# Future improvements ideas are outlined at: https:#github.com/vipnode/vipnode-contract


  # Account represents a participant in the pool. It can be a client or a host.
struct Account:
    accountBalance: wei_value
    timeLocked: timestamp

  # ForceSettle is emitted when a account requests the pool operator to settle
  # the balance and unlock it for withdrawing.
ForceSettle: event({account: address, timeLocked: timestamp})

  # Balance is emitted whenever a account balance is updated.
Balance_pool: event({account: address, balance_pool: wei_value})

  # operator is the pool operator who has permission to manage the account balances.
operator: public(address)

  # Amount of time the pool manager has to settle the balance before the
  # account can force a balance withdraw.
withdrawInterval: constant(timedelta) = 604800 

  # Mapping of owner -> account.
accounts: public(map(address, Account))


@public
def __init__(_operator: address):
    #assert
    self.operator = _operator


# addBalance adds the msg.sender as a Account.
@public
@payable
def addBalance() :
    self.accounts[msg.sender].accountBalance += msg.value

    log.Balance_pool(msg.sender, self.accounts[msg.sender].accountBalance)


  # forceSettle emits a ForceSettle event for the msg.sender if it's a valid
  # account, this prompts the pool operator to initiate balance settlement
  # and release the balance funds.
  # The account's timeLocked value will initialize to +withdrawInterval, after
  # which the account will be allowed to forceWithdraw.
  # Requires the account's balance to be >0.
@public    
def ForceSettle():
    assert self.accounts[msg.sender].accountBalance > 0

    self.accounts[msg.sender].timeLocked = withdrawInterval + block.timestamp
    log.ForceSettle(msg.sender, self.accounts[msg.sender].timeLocked)

  # forceWithdraw allows the msg.sender to withdraw any available account
  # balance iff timeLocked is set and older than now and balance > 0.
@public    
def forceWithdraw():
    assert self.accounts[msg.sender].accountBalance > 0 
    assert self.accounts[msg.sender].timeLocked > 0 
    assert self.accounts[msg.sender].timeLocked <= block.timestamp

    # Reset account
    amount: wei_value = self.accounts[msg.sender].accountBalance
    self.accounts[msg.sender].accountBalance = 0 
    self.accounts[msg.sender].timeLocked = 0 

    send(msg.sender, amount)
    log.Balance_pool(msg.sender, self.accounts[msg.sender].accountBalance)


  # -------------------
  # Operator Functions:
  # -------------------

  # Only callable by operator: Reset _account's balance and timelock and
  # transfer _amount worth of funds to _account.
  #
  # This function can be used to answer a ForceSettle request, or a normal
  # off-chain withdraw request.
  #
  # Note: The account's credit total is stored off-chain, so _releaseAmount
  # transferred can differ from the balance amount in the contract.
@public    
def opSettle(_account: address, _releaseAmount: uint256, _newBalance: uint256):
    assert msg.sender == self.operator

    self.accounts[msg.sender].accountBalance = _newBalance
    self.accounts[msg.sender].timeLocked = 0 

    if _releaseAmount > 0:
        send(_account, _releaseAmount)
        log.Balance_pool(msg.sender, self.accounts[msg.sender].accountBalance)

  # Only callable by operator: Withdraw some balance from the contract.
  #
  # This can be used to consume pool operational fees, or to pay hosts
  # off-contact. The pool operator is trusted to not abuse the balance of the
  # contract.
@public
def opWithdraw(_amount: uint256):
    assert msg.sender == self.operator

    send(msg.sender, _amount)
