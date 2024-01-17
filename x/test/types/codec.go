package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDeposit{}, "test/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "test/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "test/DeleteDeposit", nil)
	cdc.RegisterConcrete(&MsgCreateWithdraw{}, "test/CreateWithdraw", nil)
	cdc.RegisterConcrete(&MsgUpdateWithdraw{}, "test/UpdateWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeleteWithdraw{}, "test/DeleteWithdraw", nil)
	cdc.RegisterConcrete(&MsgCreateCallerGroup{}, "test/CreateCallerGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateCallerGroup{}, "test/UpdateCallerGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteCallerGroup{}, "test/DeleteCallerGroup", nil)
	cdc.RegisterConcrete(&MsgCreateSignerGroup{}, "test/CreateSignerGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateSignerGroup{}, "test/UpdateSignerGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteSignerGroup{}, "test/DeleteSignerGroup", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeposit{},
		&MsgUpdateDeposit{},
		&MsgDeleteDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWithdraw{},
		&MsgUpdateWithdraw{},
		&MsgDeleteWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCallerGroup{},
		&MsgUpdateCallerGroup{},
		&MsgDeleteCallerGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignerGroup{},
		&MsgUpdateSignerGroup{},
		&MsgDeleteSignerGroup{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
