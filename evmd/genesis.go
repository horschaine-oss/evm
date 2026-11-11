package evmd

import (
	"encoding/json"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	erc20types "github.com/cosmos/evm/x/erc20/types"
	feemarkettypes "github.com/cosmos/evm/x/feemarket/types"
	evmtypes "github.com/cosmos/evm/x/vm/types"
)

const (
	HORSCHAINBaseDenom    = "uhrs"
	HORSCHAINDisplayDenom = "HRS"
)

type GenesisState map[string]json.RawMessage

func NewBankGenesisState() *banktypes.GenesisState {
	bankGenState := banktypes.DefaultGenesisState()
	bankGenState.DenomMetadata = append(bankGenState.DenomMetadata, banktypes.Metadata{
		Description: "HORSCHAIN native coin",
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: HORSCHAINBaseDenom, Exponent: 0},
			{Denom: HORSCHAINDisplayDenom, Exponent: 18},
		},
		Base:    HORSCHAINBaseDenom,
		Display: HORSCHAINDisplayDenom,
		Name:    "HORSCHAIN",
		Symbol:  HORSCHAINDisplayDenom,
	})
	return bankGenState
}

func NewEVMGenesisState() *evmtypes.GenesisState {
	evmGenState := evmtypes.DefaultGenesisState()
	evmGenState.Params.EvmDenom = HORSCHAINBaseDenom
	evmGenState.Params.ExtendedDenomOptions = &evmtypes.ExtendedDenomOptions{
		ExtendedDenom: HORSCHAINBaseDenom,
	}
	evmGenState.Params.ActiveStaticPrecompiles = evmtypes.AvailableStaticPrecompiles
	evmGenState.Preinstalls = evmtypes.DefaultPreinstalls
	return evmGenState
}

func NewErc20GenesisState() *erc20types.GenesisState {
	erc20GenState := erc20types.DefaultGenesisState()
	return erc20GenState
}

func NewMintGenesisState() *minttypes.GenesisState {
	mintGenState := minttypes.DefaultGenesisState()
	mintGenState.Params.MintDenom = HORSCHAINBaseDenom
	return mintGenState
}

func NewFeeMarketGenesisState() *feemarkettypes.GenesisState {
	feeMarketGenState := feemarkettypes.DefaultGenesisState()
	feeMarketGenState.Params.NoBaseFee = true
	return feeMarketGenState
}

func NewStakingGenesisState() *stakingtypes.GenesisState {
	stakingGenState := stakingtypes.DefaultGenesisState()
	stakingGenState.Params.BondDenom = HORSCHAINBaseDenom
	return stakingGenState
}

func NewGovGenesisState() *govv1.GenesisState {
	govGenState := govv1.DefaultGenesisState()

	for i := range govGenState.Params.MinDeposit {
		govGenState.Params.MinDeposit[i].Denom = HORSCHAINBaseDenom
	}
	for i := range govGenState.Params.ExpeditedMinDeposit {
		govGenState.Params.ExpeditedMinDeposit[i].Denom = HORSCHAINBaseDenom
	}

	return govGenState
}
