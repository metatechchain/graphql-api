/*
Package rpc implements bridge to Meta full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an MetaTech/Meta node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Meta RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Meta RPC interface with connection limited to specified endpoints.

We strongly discourage opening Meta RPC interface for unrestricted Internet access.
*/
package rpc

//go:generate tools/abigen.sh --abi ./contracts/abi/defi-tokens-registry.abi --pkg contracts --type DefiFMintTokenRegistry --out ./contracts/fmint_tokens.go

import (
	"fmt"
	"math/big"
	"metatech-api-graphql/internal/repository/rpc/contracts"
	"metatech-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (mtc *MtcBridge) DefiTokens() ([]types.DefiToken, error) {
	// connect the contract
	contract, err := mtc.fMintCfg.tokenRegistryContract()
	if err != nil {
		return nil, err
	}

	return mtc.defiTokensList(contract)
}

// DefiTokenList creates a list of addresses / identifiers of all the ERC20 tokens
// involved with the fMint protocol.
func (mtc *MtcBridge) DefiTokenList() ([]common.Address, error) {
	// connect the contract
	contract, err := mtc.fMintCfg.tokenRegistryContract()
	if err != nil {
		return nil, err
	}

	return mtc.defiTokenAddressList(contract.TokensCount, contract.TokensList)
}

// DefiToken loads details of a single DeFi token by it's address.
func (mtc *MtcBridge) DefiToken(token *common.Address) (*types.DefiToken, error) {
	// connect the contract
	contract, err := mtc.fMintCfg.tokenRegistryContract()
	if err != nil {
		return nil, err
	}

	return mtc.defiTokenDetail(contract, token)
}

// defiTokenAddressList load list of addresses of tokens using given
// count function and item access function to do the loading.
func (mtc *MtcBridge) defiTokenAddressList(
	fCount func(*bind.CallOpts) (*big.Int, error),
	fItem func(*bind.CallOpts, *big.Int) (common.Address, error),
) ([]common.Address, error) {
	// get the number of tokens in the reference aggregator
	count, err := fCount(nil)
	if err != nil {
		mtc.log.Errorf("can not get tokens range; %s", err.Error())
		return nil, err
	}

	// make a container for tokens
	list := make([]common.Address, count.Uint64())
	index := new(big.Int)

	// load all the tokens in the contract
	for i := uint64(0); i < count.Uint64(); i++ {
		// read the indexed token from contract
		list[i], err = fItem(nil, index.SetUint64(i))
		if err != nil {
			mtc.log.Errorf("token %d address not found; %s", i, err.Error())
			return nil, err
		}
	}

	return list, nil
}

// defiTokenDetail loads details of a token specified by the token address.
func (mtc *MtcBridge) defiTokenDetail(contract *contracts.DefiFMintTokenRegistry, token *common.Address) (*types.DefiToken, error) {
	// get the token details
	tk, err := contract.Tokens(nil, *token)
	if err != nil {
		mtc.log.Errorf("token %s not found; %s", token.String(), err.Error())
		return nil, err
	}

	// decode token details
	dt, err := decodeToken(token, tk)
	if err != nil {
		mtc.log.Errorf("can not decode token %s; %s", token.String(), err.Error())
		return nil, err
	}

	return &dt, nil
}

// defiTokensList loads list of DeFi tokens from the smart contract.
func (mtc *MtcBridge) defiTokensList(contract *contracts.DefiFMintTokenRegistry) ([]types.DefiToken, error) {
	// get tge list of addresses
	al, err := mtc.defiTokenAddressList(contract.TokensCount, contract.TokensList)
	if err != nil {
		mtc.log.Errorf("tokens list not available; %s", err.Error())
		return nil, err
	}

	// make a container for tokens
	list := make([]types.DefiToken, 0)

	// load all the tokens in the contract
	for i, addr := range al {
		// decode the token
		tk, err := mtc.defiTokenDetail(contract, &addr)
		if err != nil {
			mtc.log.Errorf("invalid token #%d; %s", i, err.Error())
			return nil, err
		}

		// add the token if it's still active
		if tk.IsActive {
			list = append(list, *tk)
		}
	}

	return list, nil
}

// decodeToken decodes the contract internal token representation
// into the API structure.
func decodeToken(addr *common.Address, tk struct {
	Id            *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Logo          string
	Oracle        common.Address
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanMint       bool
}) (types.DefiToken, error) {
	// do we have a valid token? fail if not
	if tk.Id == nil || 0 == tk.Id.Uint64() {
		return types.DefiToken{}, fmt.Errorf("token undefined")
	}

	// decode and return
	return types.DefiToken{
		Address:       *addr,
		Index:         hexutil.Uint64(tk.Id.Uint64()),
		Name:          tk.Name,
		Symbol:        tk.Symbol,
		LogoUrl:       tk.Logo,
		Decimals:      int32(tk.Decimals),
		PriceDecimals: int32(tk.PriceDecimals),
		IsActive:      tk.IsActive,
		CanDeposit:    tk.CanDeposit,
		CanMint:       tk.CanMint,
		CanBorrow:     false,
		CanTrade:      false,
	}, nil
}
