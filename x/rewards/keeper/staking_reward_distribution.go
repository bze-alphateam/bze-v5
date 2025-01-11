package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DistributeAllStakingRewards(ctx sdk.Context) {
	k.IterateAllStakingRewards(ctx, k.getDistributeRewardHandler())
}

func (k Keeper) getDistributeRewardHandler() func(ctx sdk.Context, reward types.StakingReward) (stop bool) {
	return func(ctx sdk.Context, sr types.StakingReward) (stop bool) {
		logger := k.Logger(ctx).With("staking_reward", sr)

		logger.Debug("preparing to distribute staking reward")

		if sr.StakedAmount == "0" {
			logger.Debug("staking reward has no staked coins. skipping distribution")
			stop = false
			return
		}

		if sr.Payouts >= sr.Duration {
			logger.Debug("staking reward finished. skipping distribution")
			stop = false
			return
		}

		err := k.distributeStakingRewards(&sr, sr.PrizeAmount)
		if err != nil {
			logger.Error(err.Error())
			stop = false
			return
		}

		//increment payouts to know when the reward finished (a.k.a. all payouts calculated)
		sr.Payouts++
		k.SetStakingReward(ctx, sr)

		err = ctx.EventManager().EmitTypedEvent(
			&types.StakingRewardDistributionEvent{
				RewardId: sr.RewardId,
				Amount:   sr.PrizeAmount,
			},
		)

		if err != nil {
			k.Logger(ctx).Error(err.Error())
		}

		return
	}
}

func (k Keeper) distributeStakingRewards(sr *types.StakingReward, rewardAmount string) error {
	stakedAmount, err := math.LegacyNewDecFromStr(sr.StakedAmount)
	if err != nil {
		return fmt.Errorf("could not transform staked amount from storage into int: %w", err)
	}

	if !stakedAmount.IsPositive() {
		return fmt.Errorf("no stakers found")
	}

	reward, err := math.LegacyNewDecFromStr(rewardAmount)
	if err != nil {
		return fmt.Errorf("could not transform reward amount to int: %w", err)
	}

	if !reward.IsPositive() {
		return fmt.Errorf("reward amount should be positive")
	}

	sFloat, err := math.LegacyNewDecFromStr(sr.DistributedStake)
	if err != nil {
		return err
	}

	//S = S + r / T;
	sFloat = sFloat.Add(reward.Quo(stakedAmount))
	sr.DistributedStake = sFloat.String()

	return nil
}
