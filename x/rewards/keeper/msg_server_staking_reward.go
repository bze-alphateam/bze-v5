package keeper

import (
	"context"
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func (k msgServer) CreateStakingReward(goCtx context.Context, msg *types.MsgCreateStakingReward) (*types.MsgCreateStakingRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg == nil {
		return nil, sdkerrors.ErrInvalidRequest
	}

	acc, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	stakingReward, err := msg.ToStakingReward()
	if err != nil {
		return nil, err
	}

	//check denoms
	ok := k.bankKeeper.HasSupply(ctx, stakingReward.StakingDenom)
	if !ok {
		return nil, types.ErrInvalidStakingDenom
	}
	ok = k.bankKeeper.HasSupply(ctx, stakingReward.PrizeDenom)
	if !ok {
		return nil, types.ErrInvalidPrizeDenom
	}

	toCapture, err := k.getAmountToCapture(stakingReward.PrizeDenom, stakingReward.PrizeAmount, int64(stakingReward.Duration))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "could not calculate amount needed to create the reward")
	}
	fee := k.getRewardCreationFee(ctx, k.GetParams(ctx).CreateStakingRewardFee)

	neededBalance := toCapture
	if fee != nil {
		neededBalance = neededBalance.Add(fee...)
	}

	err = k.checkUserBalances(ctx, neededBalance, acc)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, acc, types.ModuleName, toCapture)
	if err != nil {
		return nil, err
	}

	if fee != nil {
		err = k.distrKeeper.FundCommunityPool(ctx, fee, acc)
		if err != nil {
			return nil, err
		}
	}

	//add ID
	stakingReward.RewardId = k.smallZeroFillId(k.GetStakingRewardsCounter(ctx))
	k.SetStakingReward(
		ctx,
		stakingReward,
	)

	err = ctx.EventManager().EmitTypedEvent(
		&types.StakingRewardCreateEvent{
			RewardId:     stakingReward.RewardId,
			PrizeAmount:  stakingReward.PrizeAmount,
			PrizeDenom:   stakingReward.PrizeDenom,
			StakingDenom: stakingReward.StakingDenom,
			Duration:     stakingReward.Duration,
			MinStake:     stakingReward.MinStake,
			Lock:         stakingReward.Lock,
		},
	)

	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	return &types.MsgCreateStakingRewardResponse{RewardId: stakingReward.RewardId}, nil
}

func (k msgServer) UpdateStakingReward(goCtx context.Context, msg *types.MsgUpdateStakingReward) (*types.MsgUpdateStakingRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg == nil {
		return nil, sdkerrors.ErrInvalidRequest
	}

	acc, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	durationInt, err := strconv.ParseInt(msg.Duration, 10, 32)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDuration, "could not convert duration to int: %s", err.Error())
	}

	if durationInt <= 0 {
		return nil, types.ErrInvalidDuration
	}

	stakingReward, isFound := k.GetStakingReward(ctx, msg.RewardId)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "staking reward not found")
	}

	toCapture, err := k.getAmountToCapture(stakingReward.PrizeDenom, stakingReward.PrizeAmount, durationInt)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "could not calculate amount needed to create the reward")
	}

	err = k.checkUserBalances(ctx, toCapture, acc)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, acc, types.ModuleName, toCapture)
	if err != nil {
		return nil, err
	}

	stakingReward.Duration += uint32(durationInt)
	k.SetStakingReward(ctx, stakingReward)

	err = ctx.EventManager().EmitTypedEvent(
		&types.StakingRewardUpdateEvent{
			RewardId: stakingReward.RewardId,
			Duration: stakingReward.Duration,
		},
	)

	if err != nil {
		k.Logger(ctx).Error(err.Error())
	}

	return &types.MsgUpdateStakingRewardResponse{}, nil
}

func (k msgServer) getRewardCreationFee(ctx sdk.Context, feeParam string) sdk.Coins {
	if feeParam == "" {
		return nil
	}

	fee, err := sdk.ParseCoinsNormalized(feeParam)
	if err != nil {
		k.Logger(ctx).Error("could not parse reward creation fee", "error", err.Error(), "feeParam", feeParam)

		return nil
	}

	if !fee.IsAllPositive() {
		return nil
	}
	//just avoid any accidental panic
	if !fee.IsValid() {
		k.Logger(ctx).Error("invalid reward creation fee", "feeParam", feeParam, "fee", fee)

		return nil
	}

	return fee
}
