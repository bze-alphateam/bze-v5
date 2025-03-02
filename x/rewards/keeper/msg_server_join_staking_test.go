package keeper_test

import (
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestMsgJoinStaking_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.JoinStaking(goCtx, nil)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_InvalidCreator() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	msg := types.MsgJoinStaking{
		Creator:  "",
		RewardId: "",
		Amount:   "",
	}

	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_MissingStakingReward() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: "0001",
		Amount:   "",
	}

	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "reward with provided id not found")
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_AmountLowerThanMinStake() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "5",
		Lock:             100,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          5,
		MinStake:         1000,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "1",
	}

	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))
	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "amount is smaller than staking reward min stake")
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_AllowedAmountLowerThanMinStake() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "5",
		Lock:             100,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          5,
		MinStake:         1000,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "1000",
	}
	//first stake the min amount allowed
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1000))).
		Times(1)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))
	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//try to stake an amount lower than min stake
	//it should be allowed since we already have a stake greater than/equal to min stake
	msg.Amount = "50"
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 50))).
		Times(1)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))
	_, err = suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().NoError(err)
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_NotEnoughBalance() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "5",
		Lock:             100,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          5,
		MinStake:         1,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	//test and assertions
	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "10",
	}
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 0)))
	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "user balance is too low")
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_Success_NewParticipant() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "0",
		DistributedStake: "0",
		Lock:             100,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          5,
		MinStake:         1,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	//test and assertions
	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "10",
	}
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 10))).
		Times(1)

	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))

	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().NoError(err)

	part, f := suite.k.GetStakingRewardParticipant(suite.ctx, msg.Creator, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(part.JoinedAt, sr.DistributedStake)
	suite.Require().EqualValues(part.Address, msg.Creator)
	suite.Require().EqualValues(part.Amount, msg.Amount)
	suite.Require().EqualValues(part.RewardId, msg.RewardId)

	storageSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(storageSr.StakedAmount, "10")
}

func (suite *IntegrationTestSuite) TestMsgJoinStaking_Success_ExistingParticipant() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "0",
		Lock:             100,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          0,
		MinStake:         10,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)

	//test and assertions
	msg := types.MsgJoinStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "10",
	}
	suite.bankMock.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 10))).
		Times(1)
	suite.bankMock.EXPECT().SpendableCoins(gomock.Any(), addr1).Times(1).Return(sdk.NewCoins(sdk.NewInt64Coin(denomBze, 1001)))

	_, err := suite.msgServer.JoinStaking(goCtx, &msg)
	suite.Require().NoError(err)

	part, f := suite.k.GetStakingRewardParticipant(suite.ctx, msg.Creator, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(part.JoinedAt, sr.DistributedStake)
	suite.Require().EqualValues(part.Address, msg.Creator)
	suite.Require().EqualValues(part.Amount, "60")
	suite.Require().EqualValues(part.RewardId, msg.RewardId)

	storageSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(storageSr.StakedAmount, "60")
}
