package keeper_test

import (
	"fmt"
	"github.com/bze-alphateam/bze/x/burner/keeper"
	"github.com/bze-alphateam/bze/x/burner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"
	"go.uber.org/mock/gomock"
)

func (suite *IntegrationTestSuite) TestGetBurnerRaffleCleanupHook() {
	hook := suite.k.GetBurnerRaffleCleanupHook()
	suite.Require().NotNil(hook)

	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "day", 1321))
	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "week", 132))
	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "hour", 11))
}

func (suite *IntegrationTestSuite) TestBurnerRaffleCleanupHook_MultipleRaffles_DifferentEndAt() {
	hook := suite.k.GetBurnerRaffleCleanupHook()
	suite.Require().NotNil(hook)

	//add to store some random data to delete
	for i := 1; i <= 5; i++ {
		denom := fmt.Sprintf("burner%d", i)
		raffleDeleteHook := types.RaffleDeleteHook{
			Denom: denom,
			EndAt: uint64(i),
		}
		suite.k.SetRaffleDeleteHook(suite.ctx, raffleDeleteHook)

		raffle := types.Raffle{
			Pot:         "5000",
			Duration:    1,
			Chances:     1,
			Ratio:       "0.1",
			EndAt:       uint64(i),
			Winners:     1,
			TicketPrice: "2",
			Denom:       denom,
		}
		suite.k.SetRaffle(suite.ctx, raffle)

		w1 := types.RaffleWinner{
			Index:  "1",
			Denom:  denom,
			Amount: "231",
			Winner: "addr_1",
		}
		suite.k.SetRaffleWinner(suite.ctx, w1)

		w2 := types.RaffleWinner{
			Index:  "2",
			Denom:  denom,
			Amount: "231",
			Winner: "addr_1",
		}
		suite.k.SetRaffleWinner(suite.ctx, w2)
	}

	//minimal check that we have something in storage
	list := suite.k.GetAllRaffle(suite.ctx)
	suite.Require().Len(list, 5)
	for i := 1; i <= 5; i++ {
		denom := fmt.Sprintf("burner%d", i)
		//expect to get module account
		suite.acc.EXPECT().GetModuleAccount(gomock.Any(), types.RaffleModuleName).Times(1).
			Return(types2.NewEmptyModuleAccount(types.RaffleModuleName))

		//expect to get balance for current denom
		expectedBalance := sdk.NewCoin(denom, sdk.NewInt(1))
		suite.bank.EXPECT().GetBalance(gomock.Any(), gomock.Any(), denom).Times(1).
			Return(expectedBalance)
		//expect to burn coins
		suite.bank.EXPECT().BurnCoins(gomock.Any(), gomock.Any(), sdk.NewCoins(expectedBalance)).Times(1).
			Return(nil)

		suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "hour", int64(i)))

		//check raffle was deleted
		_, ok := suite.k.GetRaffle(suite.ctx, denom)
		suite.Require().False(ok)

		//check winners were deleted
		winners := suite.k.GetRaffleWinners(suite.ctx, denom)
		suite.Require().Len(winners, 0)

		//check delete hook was deleted
		delHooksStored := suite.k.GetRaffleDeleteHookByEndAtPrefix(suite.ctx, uint64(i))
		suite.Require().Len(delHooksStored, 0)
	}
}

func (suite *IntegrationTestSuite) TestBurnerRaffleCleanupHook_MultipleRaffles_SameEndAt() {
	hook := suite.k.GetBurnerRaffleCleanupHook()
	suite.Require().NotNil(hook)

	//add to store some random data to delete
	for i := 1; i <= 5; i++ {
		denom := fmt.Sprintf("burner%d", i)
		if i%2 == 0 {
			denom = fmt.Sprintf("factory/%s", denom)
		}
		raffleDeleteHook := types.RaffleDeleteHook{
			Denom: denom,
			EndAt: 1,
		}
		suite.k.SetRaffleDeleteHook(suite.ctx, raffleDeleteHook)

		raffle := types.Raffle{
			Pot:         "5000",
			Duration:    1,
			Chances:     1,
			Ratio:       "0.1",
			EndAt:       1,
			Winners:     1,
			TicketPrice: "2",
			Denom:       denom,
		}
		suite.k.SetRaffle(suite.ctx, raffle)

		w1 := types.RaffleWinner{
			Index:  "1",
			Denom:  denom,
			Amount: "231",
			Winner: "addr_1",
		}
		suite.k.SetRaffleWinner(suite.ctx, w1)

		w2 := types.RaffleWinner{
			Index:  "2",
			Denom:  denom,
			Amount: "231",
			Winner: "addr_1",
		}
		suite.k.SetRaffleWinner(suite.ctx, w2)
	}

	//minimal check that we have something in storage
	list := suite.k.GetAllRaffle(suite.ctx)
	suite.Require().Len(list, 5)

	suite.acc.EXPECT().GetModuleAccount(gomock.Any(), types.RaffleModuleName).Times(5).
		Return(types2.NewEmptyModuleAccount(types.RaffleModuleName))

	suite.bank.EXPECT().GetBalance(gomock.Any(), gomock.Any(), gomock.Any()).Times(5).
		DoAndReturn(func(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
			return sdk.NewCoin(denom, sdk.NewInt(1))
		})
	//expect to burn coins
	suite.bank.EXPECT().BurnCoins(gomock.Any(), gomock.Any(), gomock.Any()).Times(5).
		Return(nil)

	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "hour", int64(1)))
	for i := 1; i <= 5; i++ {
		isFactoryDenom := false
		denom := fmt.Sprintf("burner%d", i)
		if i%2 == 0 {
			denom = fmt.Sprintf("factory/%s", denom)
			isFactoryDenom = true
		}

		//check raffle was deleted
		_, ok := suite.k.GetRaffle(suite.ctx, denom)
		suite.Require().False(ok)

		//check winners were deleted
		winners := suite.k.GetRaffleWinners(suite.ctx, denom)
		suite.Require().Len(winners, 0)

		delHooksStored := suite.k.GetRaffleDeleteHookByEndAtPrefix(suite.ctx, uint64(i))
		suite.Require().Len(delHooksStored, 0)

		allBurns := suite.k.GetAllBurnedCoins(suite.ctx)
		if !isFactoryDenom {
			suite.Require().NotEmpty(allBurns)
		}
		for _, b := range allBurns {
			if isFactoryDenom {
				suite.Require().NotContains(b.Burned, denom)
			} else {
				suite.Require().Contains(b.Burned, denom)
			}
		}
	}
}

func (suite *IntegrationTestSuite) TestGetBurnerPeriodicBurnHook_BurnSuccess() {
	hook := suite.k.GetBurnerPeriodicBurnHook()
	suite.Require().NotNil(hook)

	suite.acc.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Times(1).
		Return(types2.NewEmptyModuleAccount(types.ModuleName))

	suite.bank.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
			return sdk.NewCoins(sdk.NewCoin("ubze", sdk.NewInt(1)))
		})
	//expect to burn coins
	suite.bank.EXPECT().BurnCoins(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
		Return(nil)

	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "week", int64(keeper.BurnInterval)))

	allBurns := suite.k.GetAllBurnedCoins(suite.ctx)
	suite.Require().Len(allBurns, 1)
}

func (suite *IntegrationTestSuite) TestGetBurnerPeriodicBurnHook_BurnSkipped() {
	hook := suite.k.GetBurnerPeriodicBurnHook()
	suite.Require().NotNil(hook)

	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "day", int64(0)))

	allBurns := suite.k.GetAllBurnedCoins(suite.ctx)
	suite.Require().Len(allBurns, 0)
}

func (suite *IntegrationTestSuite) TestGetBurnerPeriodicBurnHook_EmptyBalance() {
	hook := suite.k.GetBurnerPeriodicBurnHook()
	suite.Require().NotNil(hook)
	suite.acc.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Times(1).
		Return(types2.NewEmptyModuleAccount(types.ModuleName))

	suite.bank.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
			return sdk.NewCoins()
		})
	suite.Require().NoError(hook.AfterEpochEnd(suite.ctx, "week", int64(keeper.BurnInterval*2)))

	allBurns := suite.k.GetAllBurnedCoins(suite.ctx)
	suite.Require().Len(allBurns, 0)
}
