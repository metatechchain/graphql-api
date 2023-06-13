/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access MetaTech/Meta full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"metatech-api-graphql/internal/types"
)

// StoreMtcBurn stores the given native MTC burn per block record into the persistent storage.
func (p *proxy) StoreMtcBurn(burn *types.MtcBurn) error {
	p.cache.MtcBurnUpdate(burn, p.db.BurnTotal)
	return p.db.StoreBurn(burn)
}

// MtcBurnTotal provides the total amount of burned native MTC.
func (p *proxy) MtcBurnTotal() (int64, error) {
	return p.cache.MtcBurnTotal(p.db.BurnTotal)
}

// MtcBurnList provides list of per-block burned native MTC tokens.
func (p *proxy) MtcBurnList(count int64) ([]types.MtcBurn, error) {
	return p.db.BurnList(count)
}
