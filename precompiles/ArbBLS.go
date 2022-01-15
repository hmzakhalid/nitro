//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package precompiles

// Provides a registry of BLS public keys for accounts.
type ArbBLS struct {
	Address addr // 0x67
}

// Retrieves the BLS public key for the account provided
func (con ArbBLS) GetPublicKey(c ctx, evm mech, address addr) (huge, huge, huge, huge, error) {
	return c.state.BLSTable().GetPublicKey(address)
}

// Sets the caller's BLS public key
func (con ArbBLS) Register(c ctx, evm mech, x0, x1, y0, y1 huge) error {
	return c.state.BLSTable().Register(c.caller, x0, x1, y0, y1)
}
