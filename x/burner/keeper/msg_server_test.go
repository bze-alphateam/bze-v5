package keeper_test

import (
	"fmt"
	"github.com/bze-alphateam/bze/x/burner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"
	"go.uber.org/mock/gomock"
	"strconv"
)

func (suite *IntegrationTestSuite) TestFundBurner_InvalidAmount() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgFundBurner{
		Creator: "",
		Amount:  "-1.23",
	}
	_, err := suite.msgServer.FundBurner(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "coin")
}

func (suite *IntegrationTestSuite) TestFundBurner_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgFundBurner{
		Creator: "a",
		Amount:  "123ubze",
	}
	_, err := suite.msgServer.FundBurner(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "bech32")
}

func (suite *IntegrationTestSuite) TestFundBurner_NotEnoughFunds() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgFundBurner{
		Creator: addr1.String(),
		Amount:  "123ubze",
	}

	suite.bank.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ubze", sdk.NewInt(123)))).
		Times(1).
		Return(fmt.Errorf("insufficient funds"))

	_, err := suite.msgServer.FundBurner(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "insufficient funds")
}

func (suite *IntegrationTestSuite) TestFundBurner_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgFundBurner{
		Creator: addr1.String(),
		Amount:  "123ubze",
	}
	suite.bank.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ubze", sdk.NewInt(123)))).
		Times(1).
		Return(nil)

	_, err := suite.msgServer.FundBurner(goCtx, &msg)
	suite.Require().NoError(err)
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidDenom() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgStartRaffle{
		Creator:     "",
		Pot:         "",
		Duration:    "",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}
	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(false)

	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "denom aau does not exist")
}

func (suite *IntegrationTestSuite) TestStartRaffle_RaffleAlreadyExists() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	raffle := types.Raffle{
		Pot:         "",
		Duration:    0,
		Chances:     0,
		Ratio:       "",
		EndAt:       0,
		Winners:     0,
		TicketPrice: "",
		Denom:       "aau",
	}
	suite.k.SetRaffle(suite.ctx, raffle)

	msg := types.MsgStartRaffle{
		Creator:     "",
		Pot:         "",
		Duration:    "",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}
	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "raffle already running for this coin")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	msg := types.MsgStartRaffle{
		Creator:     "a",
		Pot:         "",
		Duration:    "",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}
	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)

	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "bech32")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidPot() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "wwqdsaca",
		Duration:    "",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "invalid pot")
}

func (suite *IntegrationTestSuite) TestStartRaffle_NotPositivePot() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "0",
		Duration:    "",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "provided pot is not positive")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidDuration() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "poweiqj",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "invalid duration")
}

func (suite *IntegrationTestSuite) TestStartRaffle_NotPositiveDuration() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "0",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(2).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "duration should be positive")

	msg2 := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "-3",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}
	_, err2 := suite.msgServer.StartRaffle(goCtx, &msg2)
	suite.Require().Error(err2)
	suite.Require().ErrorContains(err2, "duration should be positive")
}

func (suite *IntegrationTestSuite) TestStartRaffle_OutOfBoundDuration() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "220",
		Chances:     "",
		Ratio:       "",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "duration have a value between")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidRatio() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "",
		Ratio:       "nskadh",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "invalid ratio")
}

func (suite *IntegrationTestSuite) TestStartRaffle_NotPositiveRatio() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "",
		Ratio:       "0",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(2).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "ratio is not positive")

	msg = types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "",
		Ratio:       "-0.05",
		TicketPrice: "",
		Denom:       "aau",
	}
	_, err = suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "ratio is not positive")
}

func (suite *IntegrationTestSuite) TestStartRaffle_RatioBoundaries() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "",
		Ratio:       "0.001",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(2).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "ratio must have a value between")

	msg = types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "",
		Ratio:       "1.0001",
		TicketPrice: "",
		Denom:       "aau",
	}
	_, err = suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "ratio must have a value between")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidChances() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "asdfgh",
		Ratio:       "0.1",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "invalid chances provided")
}

func (suite *IntegrationTestSuite) TestStartRaffle_ChancesBoundaries() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "0",
		Ratio:       "0.1",
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(2).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "chances should have a value between")

	msg = types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "1000001",
		Ratio:       "0.1",
		TicketPrice: "",
		Denom:       "aau",
	}
	_, err = suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "chances should have a value between")
}

func (suite *IntegrationTestSuite) TestStartRaffle_InvalidTicketPrice() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "20",
		Ratio:       "0.1",
		TicketPrice: "sdadsa",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "invalid ticket price provided")
}

func (suite *IntegrationTestSuite) TestStartRaffle_NegativeTicketPrice() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "20",
		Ratio:       "0.1",
		TicketPrice: "-10000002310",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "provided ticket price is not positive")
}

func (suite *IntegrationTestSuite) TestStartRaffle_IbcDenomFailure() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "109999990",
		Duration:    "15",
		Chances:     "20",
		Ratio:       "0.1",
		TicketPrice: "1000000310",
		Denom:       "ibc/aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "ibc/aau").Times(1).Return(true)
	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "coin not allowed in raffles")
}

func (suite *IntegrationTestSuite) TestStartRaffle_NotEnoughBalance() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "20",
		Ratio:       "0.1",
		TicketPrice: "150000000",
		Denom:       "aau",
	}

	suite.bank.EXPECT().HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	suite.bank.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).
		Return(sdk.NewCoins())
	suite.epoch.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))

	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "balance")
}

func (suite *IntegrationTestSuite) TestStartRaffle_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgStartRaffle{
		Creator:     addr1.String(),
		Pot:         "100",
		Duration:    "15",
		Chances:     "20",
		Ratio:       "0.1",
		TicketPrice: "150000000",
		Denom:       "aau",
	}

	suite.bank.EXPECT().
		HasSupply(gomock.Any(), "aau").Times(1).Return(true)
	suite.bank.EXPECT().
		SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin("aau", 500)))
	suite.epoch.EXPECT().
		GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.acc.EXPECT().
		GetModuleAccount(gomock.Any(), types.RaffleModuleName).Times(5).
		Return(types2.NewEmptyModuleAccount(types.RaffleModuleName))
	suite.bank.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.RaffleModuleName, sdk.NewCoins(sdk.NewInt64Coin("aau", 100)))

	_, err := suite.msgServer.StartRaffle(goCtx, &msg)
	suite.Require().NoError(err)
	storageRaffle, ok := suite.k.GetRaffle(suite.ctx, "aau")
	suite.Require().True(ok)
	suite.Require().EqualValues(msg.Pot, storageRaffle.Pot)
	suite.Require().EqualValues(msg.Duration, strconv.Itoa(int(storageRaffle.Duration)))
	suite.Require().EqualValues(msg.Chances, strconv.Itoa(int(storageRaffle.Chances)))
	suite.Require().EqualValues(msg.Ratio, storageRaffle.Ratio)
	suite.Require().EqualValues(msg.Denom, storageRaffle.Denom)
	suite.Require().EqualValues(msg.TicketPrice, storageRaffle.TicketPrice)
	suite.Require().EqualValues(uint64(1+15*24), storageRaffle.EndAt)

	deleteHook := suite.k.GetRaffleDeleteHookByEndAtPrefix(suite.ctx, uint64(1+15*24))
	suite.Require().NotEmpty(deleteHook)
	suite.Require().Equal(deleteHook[0].Denom, "aau")
	suite.Require().EqualValues(deleteHook[0].EndAt, uint64(1+15*24))
}

func (suite *IntegrationTestSuite) TestJoinRaffle_InvalidDenom() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgJoinRaffle{
		Creator: "",
		Denom:   "dsa",
	}
	suite.bank.EXPECT().
		HasSupply(gomock.Any(), "dsa").Times(1).Return(false)

	_, err := suite.msgServer.JoinRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "denom")
}

func (suite *IntegrationTestSuite) TestJoinRaffle_RaffleNotFound() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgJoinRaffle{
		Creator: "",
		Denom:   "aau",
	}
	suite.bank.EXPECT().
		HasSupply(gomock.Any(), "aau").Times(1).Return(true)

	_, err := suite.msgServer.JoinRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "raffle")
}

func (suite *IntegrationTestSuite) TestJoinRaffle_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	raffle := types.Raffle{
		Pot:         "",
		Duration:    0,
		Chances:     0,
		Ratio:       "",
		EndAt:       10,
		Winners:     0,
		TicketPrice: "",
		Denom:       "aau",
	}

	suite.k.SetRaffle(suite.ctx, raffle)

	msg := types.MsgJoinRaffle{
		Creator: "aa",
		Denom:   "aau",
	}
	suite.epoch.EXPECT().
		GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.bank.EXPECT().
		HasSupply(gomock.Any(), "aau").Times(1).Return(true)

	_, err := suite.msgServer.JoinRaffle(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "bech32")
}
