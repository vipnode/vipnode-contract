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
  // TODO: ...
}
