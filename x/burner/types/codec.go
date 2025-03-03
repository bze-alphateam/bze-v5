package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&BurnCoinsProposal{}, "burner/BurnCoinsProposal", nil)
	cdc.RegisterConcrete(&MsgFundBurner{}, "burner/FundBurner", nil)
	cdc.RegisterConcrete(&MsgStartRaffle{}, "burner/StartRaffle", nil)
	cdc.RegisterConcrete(&MsgJoinRaffle{}, "burner/JoinRaffle", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFundBurner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartRaffle{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinRaffle{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&BurnCoinsProposal{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
}
