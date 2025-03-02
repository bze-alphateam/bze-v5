package keeper_test

import (
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestCreateStakingReward_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.CreateStakingReward(goCtx, nil)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestUpdateStakingReward_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.UpdateStakingReward(goCtx, nil)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestCreateStakingReward_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	req := &types.MsgCreateStakingReward{Creator: ""}

	_, err := suite.msgServer.CreateStakingReward(goCtx, req)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestUpdateStakingReward_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	req := &types.MsgUpdateStakingReward{Creator: ""}

	_, err := suite.msgServer.UpdateStakingReward(goCtx, req)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateStakingReward_InvalidStakingReward() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	tests := []struct {
		name string
		msg  types.MsgCreateStakingReward
	}{
		{
			name: "empty prize amount",
			msg: types.MsgCreateStakingReward{
				Creator:     addr1.String(),
				PrizeAmount: "",
			},
		},
		{
			name: "zero prize amount",
			msg: types.MsgCreateStakingReward{
				Creator:     addr1.String(),
				PrizeAmount: "0",
			},
		},
		{
			name: "negative prize amount",
			msg: types.MsgCreateStakingReward{
				Creator:     addr1.String(),
				PrizeAmount: "-10",
			},
		},
		{
			name: "empty prize denom",
			msg: types.MsgCreateStakingReward{
				Creator:     addr1.String(),
				PrizeAmount: "10",
				PrizeDenom:  "",
			},
		},
		{
			name: "empty staking denom",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "",
			},
		},
		{
			name: "invalid min stake",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "",
			},
		},
		{
			name: "negative min stake",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "-10",
			},
		},
		{
			name: "invalid duration",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "",
			},
		},
		{
			name: "duration too low",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "0",
			},
		},
		{
			name: "duration too high",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "3213132131231",
			},
		},
		{
			name: "invalid lock",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "100",
				Lock:         "asd",
			},
		},
		{
			name: "negative lock",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "100",
				Lock:         "-33",
			},
		},
		{
			name: "lock too high",
			msg: types.MsgCreateStakingReward{
				Creator:      addr1.String(),
				PrizeAmount:  "10",
				PrizeDenom:   "ubze",
				StakingDenom: "ubze",
				MinStake:     "10",
				Duration:     "100",
				Lock:         "7899",
			},
		},
	}
	for _, tt := range tests {
		_, err := suite.msgServer.CreateStakingReward(goCtx, &tt.msg)
		suite.Require().NotNil(err)
	}
}

func (suite *IntegrationTestSuite) TestUpdateStakingReward_InvalidStakingReward() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	tests := []struct {
		name             string
		msg              types.MsgUpdateStakingReward
		addStakingReward bool
	}{
		{
			name: "invalid duration",
			msg: types.MsgUpdateStakingReward{
				Creator:  addr1.String(),
				Duration: "",
			},
		},
		{
			name: "zero duration",
			msg: types.MsgUpdateStakingReward{
				Creator:  addr1.String(),
				Duration: "0",
			},
		},
		{
			name: "negative duration",
			msg: types.MsgUpdateStakingReward{
				Creator:  addr1.String(),
				Duration: "-220",
			},
		},
		{
			name: "missing staking reward",
			msg: types.MsgUpdateStakingReward{
				Creator:  addr1.String(),
				Duration: "10",
				RewardId: "notastakingrewardid",
			},
		},
		{
			name: "not enough balance",
			msg: types.MsgUpdateStakingReward{
				Creator:  addr1.String(),
				Duration: "10",
				RewardId: "001",
			},
			addStakingReward: true,
		},
	}
	stakingReward := types.StakingReward{RewardId: "001", PrizeDenom: "ubze", PrizeAmount: "10"}
	suite.k.SetStakingReward(suite.ctx, stakingReward)

	for _, tt := range tests {
		if tt.addStakingReward {
			suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Return(sdk.NewCoins())
		}

		_, err := suite.msgServer.UpdateStakingReward(goCtx, &tt.msg)
		suite.Require().NotNil(err)
	}
}

func (suite *IntegrationTestSuite) TestCreateStakingReward_MissingSupply() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateStakingReward{
		Creator:      addr1.String(),
		PrizeAmount:  "10",
		PrizeDenom:   denomBze,
		StakingDenom: denomBze,
		MinStake:     "10",
		Duration:     "100",
		Lock:         "1",
	}

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(1).Return(false)
	_, err := suite.msgServer.CreateStakingReward(goCtx, &msg)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateStakingReward_NotEnoughBalance() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateStakingReward{
		Creator:      addr1.String(),
		PrizeAmount:  "10",
		PrizeDenom:   denomBze,
		StakingDenom: denomBze,
		MinStake:     "10",
		Duration:     "100",
		Lock:         "1",
	}

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(2).Return(true)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 10*99)))
	_, err := suite.msgServer.CreateStakingReward(goCtx, &msg)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestCreateStakingReward_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	msg := types.MsgCreateStakingReward{
		Creator:      addr1.String(),
		PrizeAmount:  "10",
		PrizeDenom:   denomBze,
		StakingDenom: denomBze,
		MinStake:     "10",
		Duration:     "100",
		Lock:         "1",
	}

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Times(2).Return(true)
	//user has balance for prize and for staking reward creation fee too
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 30_000_000_000_000)))
	//expect user to pay for the rewards (duration * prize amount)
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1000))).
		Times(1)
	//expect user to pay for staking reward creation
	suite.distrMock.EXPECT().FundCommunityPool(goCtx, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 25000_000000)), addr1).Times(1).Return(nil)

	res, err := suite.msgServer.CreateStakingReward(goCtx, &msg)
	suite.Require().NoError(err)

	storeStakingReward, ok := suite.k.GetStakingReward(suite.ctx, res.RewardId)
	suite.Require().True(ok)

	suite.Require().EqualValues(msg.PrizeAmount, storeStakingReward.PrizeAmount)
	suite.Require().EqualValues(msg.PrizeDenom, storeStakingReward.PrizeDenom)
	suite.Require().EqualValues(msg.StakingDenom, storeStakingReward.StakingDenom)
	suite.Require().EqualValues(10, storeStakingReward.MinStake)
	suite.Require().EqualValues(100, storeStakingReward.Duration)
	suite.Require().EqualValues(1, storeStakingReward.Lock)
	suite.Require().EqualValues(0, storeStakingReward.Payouts)
}

func (suite *IntegrationTestSuite) TestUpdateStakingReward_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")

	stakingReward := types.StakingReward{
		RewardId:     "001",
		PrizeAmount:  "10",
		PrizeDenom:   denomBze,
		StakingDenom: denomBze,
		MinStake:     10,
		Duration:     100,
		Lock:         1,
	}
	suite.k.SetStakingReward(suite.ctx, stakingReward)

	msg := types.MsgUpdateStakingReward{
		Creator:  addr1.String(),
		Duration: "200",
		RewardId: "001",
	}

	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 3000)))
	//expect user to pay for the rewards (duration * prize amount)
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 2000))).
		Times(1)

	_, err := suite.msgServer.UpdateStakingReward(goCtx, &msg)
	suite.Require().NoError(err)

	storeStakingReward, ok := suite.k.GetStakingReward(suite.ctx, stakingReward.RewardId)
	suite.Require().True(ok)

	//check duration was updated
	suite.Require().EqualValues(uint32(300), storeStakingReward.Duration)
	//check other fields were not changed
	suite.Require().EqualValues(stakingReward.PrizeAmount, storeStakingReward.PrizeAmount)
	suite.Require().EqualValues(stakingReward.PrizeDenom, storeStakingReward.PrizeDenom)
	suite.Require().EqualValues(stakingReward.StakingDenom, storeStakingReward.StakingDenom)
	suite.Require().EqualValues(stakingReward.MinStake, storeStakingReward.MinStake)
	suite.Require().EqualValues(stakingReward.Lock, storeStakingReward.Lock)
}
