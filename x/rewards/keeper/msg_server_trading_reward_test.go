package keeper_test

import (
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestCreateTradingReward_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.CreateTradingReward(goCtx, nil)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestCreateTradingReward_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	req := &types.MsgCreateTradingReward{Creator: ""}

	_, err := suite.msgServer.CreateTradingReward(goCtx, req)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateTradingReward_InvalidTradingReward() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	tests := []struct {
		name string
		msg  types.MsgCreateTradingReward
	}{
		{
			name: "empty prize amount",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "",
			},
		},
		{
			name: "zero prize amount",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "0",
			},
		},
		{
			name: "negative prize amount",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "-10",
			},
		},
		{
			name: "empty prize denom",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "",
			},
		},
		{
			name: "missing market id",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				MarketId:    "not_a_market_id",
			},
		},
		{
			name: "invalid duration",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				MarketId:    "stake/ubze",
				Duration:    "",
			},
		},
		{
			name: "zero duration",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "0",
				MarketId:    "stake/ubze",
			},
		},
		{
			name: "negative duration",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "-11",
				MarketId:    "stake/ubze",
			},
		},
		{
			name: "duration too high",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "3213132131231",
				MarketId:    "stake/ubze",
			},
		},
		{
			name: "invalid slots",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "100",
				MarketId:    "stake/ubze",
				Slots:       "",
			},
		},
		{
			name: "zero slots",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "100",
				MarketId:    "stake/ubze",
				Slots:       "0",
			},
		},
		{
			name: "negative slots",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "100",
				MarketId:    "stake/ubze",
				Slots:       "-3",
			},
		},
		{
			name: "too many slots slots",
			msg: types.MsgCreateTradingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "ubze",
				Duration:    "100",
				MarketId:    "stake/ubze",
				Slots:       "101",
			},
		},
	}

	for _, tt := range tests {
		_, err := suite.msgServer.CreateTradingReward(goCtx, &tt.msg)
		suite.Require().NotNil(err, tt.name)
	}
}

func (suite *IntegrationTestSuite) TestCreateTradingReward_MissingSupply() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateTradingReward{
		Creator:     addr1.String(),
		PrizeAmount: "10",
		PrizeDenom:  denomBze,
		Duration:    "100",
		MarketId:    "stake/ubze",
		Slots:       "1",
	}

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(1).Return(false)
	_, err := suite.msgServer.CreateTradingReward(goCtx, &msg)
	suite.Require().ErrorIs(err, types.ErrInvalidPrizeDenom)
}

func (suite *IntegrationTestSuite) TestCreateTradingReward_NotEnoughBalance() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateTradingReward{
		Creator:     addr1.String(),
		PrizeAmount: "10",
		PrizeDenom:  "ubze",
		Duration:    "100",
		MarketId:    "stake/ubze",
		Slots:       "10",
	}

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(1).Return(true)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins())
	suite.tradeMock.EXPECT().MarketExists(gomock.Any(), "stake/ubze").Return(true).MinTimes(1)

	_, err := suite.msgServer.CreateTradingReward(goCtx, &msg)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestCreateTradingReward_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateTradingReward{
		Creator:     addr1.String(),
		PrizeAmount: "200",
		PrizeDenom:  denomBze,
		Duration:    "100",
		MarketId:    "stake/ubze",
		Slots:       "10",
	}

	suite.epochMock.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Return(int64(1)).Times(1)
	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(1).Return(true)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 51000_000000)))
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 2000))).
		Times(1)
	suite.tradeMock.EXPECT().MarketExists(gomock.Any(), "stake/ubze").Return(true).Times(1)
	suite.distrMock.EXPECT().FundCommunityPool(gomock.Any(), sdk.NewCoins(sdk.NewInt64Coin(denomBze, 50000_000000)), addr1).Times(1)

	res, err := suite.msgServer.CreateTradingReward(goCtx, &msg)
	suite.Require().NoError(err)

	storeTradingReward, ok := suite.k.GetPendingTradingReward(suite.ctx, res.RewardId)
	suite.Require().True(ok)

	suite.Require().EqualValues(msg.PrizeAmount, storeTradingReward.PrizeAmount)
	suite.Require().EqualValues(msg.PrizeDenom, storeTradingReward.PrizeDenom)
	suite.Require().EqualValues(uint32(100), storeTradingReward.Duration)
	suite.Require().EqualValues(msg.MarketId, storeTradingReward.MarketId)
	suite.Require().EqualValues(uint32(10), storeTradingReward.Slots)
}
