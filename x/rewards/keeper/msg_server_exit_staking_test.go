package keeper_test

import (
	"fmt"
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestMsgExitStaking_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.ExitStaking(goCtx, nil)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_MissingStakingReward() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	msg := types.MsgExitStaking{
		Creator:  "",
		RewardId: "a",
	}

	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "reward with provided id not found")
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_MissingStakingRewardParticipant() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "0",
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	msg := types.MsgExitStaking{
		Creator:  "aaa",
		RewardId: sr.RewardId,
	}
	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "you are not a participant in this staking reward")
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_Success_OngoingStakingReward_NoClaim() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "72",
		DistributedStake: "5",
		Lock:             100,
		StakingDenom:     denomBze,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	untouchedSrp := types.StakingRewardParticipant{
		Address:  "asadsadasda",
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}
	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "22",
		JoinedAt: "5",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)
	suite.k.SetStakingRewardParticipant(suite.ctx, untouchedSrp)

	//tests and asserts
	msg := types.MsgExitStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
	}
	suite.epochMock.EXPECT().GetEpochCountByIdentifier(goCtx, "hour").Return(int64(1))
	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//check the unlock was created
	lockKey := types.CreatePendingUnlockParticipantKey(int64(sr.Lock*24+1), fmt.Sprintf("%s/%s", sr.RewardId, srp.Address))
	unlockList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().NotEmpty(unlockList)
	suite.Require().EqualValues(unlockList[0].Index, lockKey)
	suite.Require().EqualValues(unlockList[0].Address, srp.Address)
	suite.Require().EqualValues(unlockList[0].Amount, "22")
	suite.Require().EqualValues(unlockList[0].Denom, sr.StakingDenom)

	//check the staking reward participant was deleted
	_, f := suite.k.GetStakingRewardParticipant(suite.ctx, srp.Address, sr.RewardId)
	suite.Require().False(f)

	//check the staking reward was updated
	newSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(newSr.StakedAmount, "50")

	//check that the dummy srp is not touched since it wasn't belonging to the message creator
	untouchedSrpStorage, f := suite.k.GetStakingRewardParticipant(suite.ctx, untouchedSrp.Address, untouchedSrp.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(untouchedSrpStorage, untouchedSrp)
}

// tests that upon staking exit the user will also receive his unclaimed rewards
func (suite *IntegrationTestSuite) TestMsgExitStaking_Success_OngoingStakingReward_WithPendingToClaim() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "4",
		Lock:             100,
		StakingDenom:     denomBze,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	untouchedSrp := types.StakingRewardParticipant{
		Address:  "asadsadasda",
		RewardId: sr.RewardId,
		Amount:   "25",
		JoinedAt: "0",
	}
	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "25",
		JoinedAt: "0",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)
	suite.k.SetStakingRewardParticipant(suite.ctx, untouchedSrp)

	//tests and asserts
	msg := types.MsgExitStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
	}
	suite.epochMock.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 100))).
		Times(1)

	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//check the unlock was created
	lockKey := types.CreatePendingUnlockParticipantKey(int64(sr.Lock*24+1), fmt.Sprintf("%s/%s", sr.RewardId, srp.Address))
	unlockList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().NotEmpty(unlockList)
	suite.Require().EqualValues(unlockList[0].Index, lockKey)
	suite.Require().EqualValues(unlockList[0].Address, srp.Address)
	suite.Require().EqualValues(unlockList[0].Amount, "25")
	suite.Require().EqualValues(unlockList[0].Denom, sr.StakingDenom)

	//check the staking reward participant was deleted
	_, f := suite.k.GetStakingRewardParticipant(suite.ctx, srp.Address, sr.RewardId)
	suite.Require().False(f)

	//check the staking reward was updated
	newSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(newSr.StakedAmount, "25")

	//check that the dummy srp is not touched since it wasn't belonging to the message creator
	untouchedSrpStorage, f := suite.k.GetStakingRewardParticipant(suite.ctx, untouchedSrp.Address, untouchedSrp.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(untouchedSrpStorage, untouchedSrp)
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_Success_EmptyingStakingReward() {
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
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	untouchedSrp := types.StakingRewardParticipant{
		Address:  "asadsadasda",
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}

	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)
	suite.k.SetStakingRewardParticipant(suite.ctx, untouchedSrp)

	//tests and asserts
	msg := types.MsgExitStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
	}
	suite.epochMock.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 250))).
		Times(1)
	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//check the unlock was created
	lockKey := types.CreatePendingUnlockParticipantKey(int64(sr.Lock*24+1), fmt.Sprintf("%s/%s", sr.RewardId, srp.Address))
	unlockList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().NotEmpty(unlockList)
	suite.Require().EqualValues(unlockList[0].Index, lockKey)
	suite.Require().EqualValues(unlockList[0].Address, srp.Address)
	suite.Require().EqualValues(unlockList[0].Amount, "50")
	suite.Require().EqualValues(unlockList[0].Denom, sr.StakingDenom)

	//check the staking reward participant was deleted
	_, f := suite.k.GetStakingRewardParticipant(suite.ctx, srp.Address, sr.RewardId)
	suite.Require().False(f)

	//check the staking reward was updated
	newSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(newSr.StakedAmount, "0")

	//check that the dummy srp is not touched since it wasn't belonging to the message creator
	untouchedSrpStorage, f := suite.k.GetStakingRewardParticipant(suite.ctx, untouchedSrp.Address, untouchedSrp.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(untouchedSrpStorage, untouchedSrp)
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_Success_RemovingStakingReward() {
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
		Payouts:          100,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)

	//tests and asserts
	msg := types.MsgExitStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
	}
	suite.epochMock.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 250))).
		Times(1)
	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//check the unlock was created
	lockKey := types.CreatePendingUnlockParticipantKey(int64(sr.Lock*24+1), fmt.Sprintf("%s/%s", sr.RewardId, srp.Address))
	unlockList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().NotEmpty(unlockList)
	suite.Require().EqualValues(unlockList[0].Index, lockKey)
	suite.Require().EqualValues(unlockList[0].Address, srp.Address)
	suite.Require().EqualValues(unlockList[0].Amount, "50")
	suite.Require().EqualValues(unlockList[0].Denom, sr.StakingDenom)

	//check the staking reward participant was deleted
	_, f := suite.k.GetStakingRewardParticipant(suite.ctx, srp.Address, sr.RewardId)
	suite.Require().False(f)

	//check the staking reward was deleted
	_, f = suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().False(f)
}

func (suite *IntegrationTestSuite) TestMsgExitStaking_Success_RemovingStakingReward_WithoutLock() {
	//dependencies
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeDenom:       denomBze,
		StakedAmount:     "50",
		DistributedStake: "5",
		Lock:             0,
		StakingDenom:     denomBze,
		Duration:         100,
		Payouts:          100,
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	srp := types.StakingRewardParticipant{
		Address:  addr1.String(),
		RewardId: sr.RewardId,
		Amount:   "50",
		JoinedAt: "0",
	}
	suite.k.SetStakingRewardParticipant(suite.ctx, srp)

	//tests and asserts
	msg := types.MsgExitStaking{
		Creator:  addr1.String(),
		RewardId: sr.RewardId,
	}
	suite.epochMock.EXPECT().GetEpochCountByIdentifier(gomock.Any(), "hour").Times(1).Return(int64(1))
	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 250))).
		Times(1)
	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewInt64Coin(denomBze, 50))).
		Times(1)
	_, err := suite.msgServer.ExitStaking(goCtx, &msg)
	suite.Require().NoError(err)

	//check the unlock was NOT created since the funds should be released immediately
	unlockList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().Empty(unlockList)

	//check the staking reward participant was deleted
	_, f := suite.k.GetStakingRewardParticipant(suite.ctx, srp.Address, sr.RewardId)
	suite.Require().False(f)

	//check the staking reward was deleted
	_, f = suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().False(f)
}
