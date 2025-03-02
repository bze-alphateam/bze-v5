package keeper_test

import (
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestGetDistributeAllStakingRewardsHook_NoEpoch() {
	hook := suite.k.GetDistributeAllStakingRewardsHook()
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeAmount:      "100",
		PrizeDenom:       denomBze,
		StakingDenom:     denomStake,
		Duration:         10,
		Payouts:          0,
		MinStake:         0,
		Lock:             10,
		StakedAmount:     "10",
		DistributedStake: "0",
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	err := hook.AfterEpochEnd(suite.ctx, "not_an_epoch", 321)
	suite.Require().NoError(err)

	storageSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().EqualValues(storageSr, sr)
}

func (suite *IntegrationTestSuite) TestGetDistributeAllStakingRewardsHook_Success() {
	hook := suite.k.GetDistributeAllStakingRewardsHook()
	sr := types.StakingReward{
		RewardId:         "01",
		PrizeAmount:      "100",
		PrizeDenom:       denomBze,
		StakingDenom:     denomStake,
		Duration:         10,
		Payouts:          0,
		MinStake:         0,
		Lock:             10,
		StakedAmount:     "10",
		DistributedStake: "0",
	}
	suite.k.SetStakingReward(suite.ctx, sr)

	err := hook.AfterEpochEnd(suite.ctx, "day", 101)
	suite.Require().NoError(err)

	storageSr, f := suite.k.GetStakingReward(suite.ctx, sr.RewardId)
	suite.Require().True(f)
	suite.Require().NotEqualValues(storageSr, sr)
}

func (suite *IntegrationTestSuite) TestGetUnlockPendingUnlockParticipantsHook_NoEpoch() {
	hook := suite.k.GetUnlockPendingUnlockParticipantsHook()

	pup := types.PendingUnlockParticipant{Index: types.CreatePendingUnlockParticipantKey(int64(321), "something")}
	suite.k.SetPendingUnlockParticipant(suite.ctx, pup)

	err := hook.AfterEpochEnd(suite.ctx, "not_an_epoch", 321)
	suite.Require().NoError(err)

	storageList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().NotEmpty(storageList)
	suite.Require().EqualValues(pup, storageList[0])
}

func (suite *IntegrationTestSuite) TestGetUnlockPendingUnlockParticipantsHook_Success() {
	hook := suite.k.GetUnlockPendingUnlockParticipantsHook()
	addr1 := sdk.AccAddress("addr1_______________")
	pup := types.PendingUnlockParticipant{
		Index:   types.CreatePendingUnlockParticipantKey(int64(321), "something"),
		Address: addr1.String(),
		Amount:  "1",
		Denom:   denomBze,
	}
	suite.k.SetPendingUnlockParticipant(suite.ctx, pup)

	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(1)))).
		Times(1).Return(nil)

	err := hook.AfterEpochEnd(suite.ctx, "hour", 321)
	suite.Require().NoError(err)

	storageList := suite.k.GetAllPendingUnlockParticipant(suite.ctx)
	suite.Require().Empty(storageList)
}

func (suite *IntegrationTestSuite) TestGetRemoveExpiredPendingTradingRewardsHook_NoEpoch() {
	hook := suite.k.GetRemoveExpiredPendingTradingRewardsHook()

	ptre := types.TradingRewardExpiration{RewardId: "01", ExpireAt: uint32(1)}
	suite.k.SetPendingTradingRewardExpiration(suite.ctx, ptre)

	suite.bankMock.EXPECT().BurnCoins(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	err := hook.AfterEpochEnd(suite.ctx, "not_an_epoch", 1)
	suite.Require().NoError(err)

	storageList := suite.k.GetAllPendingTradingRewardExpiration(suite.ctx)
	suite.Require().NotEmpty(storageList)
}

func (suite *IntegrationTestSuite) TestGetRemoveExpiredPendingTradingRewardsHook_Success() {
	hook := suite.k.GetRemoveExpiredPendingTradingRewardsHook()
	tr := types.TradingReward{
		RewardId:    "01",
		PrizeDenom:  denomBze,
		PrizeAmount: "100",
		Slots:       1,
	}
	suite.k.SetPendingTradingReward(suite.ctx, tr)
	ptre := types.TradingRewardExpiration{RewardId: tr.RewardId, ExpireAt: uint32(1)}
	suite.k.SetPendingTradingRewardExpiration(suite.ctx, ptre)

	suite.bankMock.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(100)))).Times(1)
	err := hook.AfterEpochEnd(suite.ctx, "hour", 1)
	suite.Require().NoError(err)

	_, f := suite.k.GetPendingTradingReward(suite.ctx, tr.RewardId)
	suite.Require().False(f)

	expList := suite.k.GetAllPendingTradingRewardExpiration(suite.ctx)
	suite.Require().Empty(expList)
}

func (suite *IntegrationTestSuite) TestGetTradingRewardsDistributionHook_NoEpoch() {
	hook := suite.k.GetTradingRewardsDistributionHook()

	tr := types.TradingReward{
		RewardId:    "01",
		PrizeDenom:  denomBze,
		PrizeAmount: "100",
		Slots:       1,
	}
	suite.k.SetActiveTradingReward(suite.ctx, tr)
	ptre := types.TradingRewardExpiration{RewardId: tr.RewardId, ExpireAt: uint32(1)}
	suite.k.SetActiveTradingRewardExpiration(suite.ctx, ptre)

	suite.bankMock.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, gomock.Any(), gomock.Any()).
		Times(0)

	err := hook.AfterEpochEnd(suite.ctx, "not_an_epoch", 1)
	suite.Require().NoError(err)
}
