# vipnode-contract

Smart contract for the Vipnode Pool.

**Status**: Work in progress.

## Overview

### Version 1: [Vipnode.sol](https://github.com/vipnode/vipnode-contract/blob/bd0d788966cacd5924a39888bc56abe9b8a71e4f/contracts/Vipnode.sol) (Deprecated)

Minimum Viable Product implementation. Unlimited usage per interval billing,
entirely on-chain.


### Version 2: [VipnodePool.sol](https://github.com/vipnode/vipnode-contract/blob/master/contracts/VipnodePool.sol)

The goals are:

1. Support arbitrary per-usage billing
2. Simplest implementation
3. Minimize on-chain transactions

The compromises is that the pool is trusted using the same model as mining
pools: The pool operator can run away with the funds from the previous billing
interval.

### Version 3+ (Someday)

Reduce trust required in the pool. Ideally, the pool will only be required to
coordinate the initial connection between the host and client.

- Use state channels for direct payments between clients and hosts. This sounds
  good in theory, but in practice doing N:M state channels (many clients paying
  to many hosts) is incredibly expensive. Recall that any client can connect to
  any number of hosts over a given billing interval. To use state channels
  efficiently, we would need collateralized relays to aggregate N:M payments
  into fewer on-chain transactions, but this increases the total capitalization
  requirement.
- Use zero-knowledge proofs to aggregate transactions without collateralized
  relays. This is an active area of research.


## License

MIT
