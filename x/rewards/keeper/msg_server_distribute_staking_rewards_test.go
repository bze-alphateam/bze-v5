package keeper_test

import (
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, nil)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_InvalidCreator() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	msg := types.MsgDistributeStakingRewards{
		Creator: "",
	}

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().Equal(err.Error(), "empty address string is not allowed")
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_InvalidAmount() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgDistributeStakingRewards{
		Creator: addr1.String(),
	}

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "could not convert order amount")
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_NegativeAmount() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgDistributeStakingRewards{
		Creator: addr1.String(),
		Amount:  "-1",
	}

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "amount should be greater than 0")
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_MissingStakingReward() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgDistributeStakingRewards{
		Creator: addr1.String(),
		Amount:  "100",
	}

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "staking reward not found")
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_NoUserBalance() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "0",
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	msg := types.MsgDistributeStakingRewards{
		Creator:  addr1.String(),
		Amount:   "100",
		RewardId: sr.RewardId,
	}

	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins())

	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)
}

func (suite *IntegrationTestSuite) TestMsgDistributeStakingRewards_Success() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "0",
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	msg := types.MsgDistributeStakingRewards{
		Creator:  addr1.String(),
		Amount:   "1000",
		RewardId: sr.RewardId,
	}

	//user has balance in bank more than distributed reward
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(
		gomock.Any(),
		addr1,
		types.ModuleName,
		sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(1000))),
	).Times(1).Return(nil)

	//test & asserts
	_, err := suite.msgServer.DistributeStakingRewards(goCtx, &msg)
	suite.Require().NoError(err)

	//check sr was updated
	newSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().NotEqualValues(newSr, sr)
	suite.Require().NotEqualValues(newSr.DistributedStake, sr.DistributedStake)
}
