// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"math/big"
	"metatech-api-graphql/internal/repository"
	"metatech-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MtcBurnedTotal resolves total amount of burned MTC tokens in WEI units.
func (rs *rootResolver) MtcBurnedTotal() hexutil.Big {
	val, err := repository.R().MtcBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return hexutil.Big{}
	}
	return hexutil.Big(*new(big.Int).Mul(big.NewInt(val), types.BurnDecimalsCorrection))
}

// MtcBurnedTotalAmount resolves total amount of burned MTC tokens in MTC units.
func (rs *rootResolver) MtcBurnedTotalAmount() float64 {
	val, err := repository.R().MtcBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return 0
	}
	return float64(val) / types.BurnMTCDecimalsCorrection
}

// MtcLatestBlockBurnList resolves a list of the latest block burns.
func (rs *rootResolver) MtcLatestBlockBurnList(args struct{ Count int32 }) ([]types.MtcBurn, error) {
	if args.Count < 1 || args.Count > 50 {
		args.Count = 25
	}
	return repository.R().MtcBurnList(int64(args.Count))
}
