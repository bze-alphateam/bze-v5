package keeper

import (
	"context"
	"fmt"
	"github.com/bze-alphateam/bze/x/burner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	raffleDelayHeight = 2
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) FundBurner(goCtx context.Context, msg *types.MsgFundBurner) (*types.MsgFundBurnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	creatorAccount, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAccount, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&types.FundBurnerEvent{From: msg.Creator, Amount: amount.String()})
	if err != nil {
		ctx.Logger().Error("failed to emit fund burner event", "error", err)
	}

	return &types.MsgFundBurnerResponse{}, nil
}

func (k msgServer) StartRaffle(goCtx context.Context, msg *types.MsgStartRaffle) (*types.MsgStartRaffleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.bankKeeper.HasSupply(ctx, msg.Denom) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "denom %s does not exist", msg.Denom)
	}

	_, alreadyStarted := k.GetRaffle(ctx, msg.Denom)
	if alreadyStarted {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "raffle already running for this coin")
	}

	creatorAcc, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	raffle, err := k.raffleFromMsgStartRaffle(ctx, msg)
	if err != nil {
		return nil, err
	}

	//do not check if OK because it is checked in basic validation and in method that converts message to storage struct
	potAmt, _ := sdk.NewIntFromString(raffle.Pot)
	pot := sdk.NewCoin(raffle.Denom, potAmt)
	if !k.accountHasCoins(ctx, pot, creatorAcc) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not enough balance")
	}

	if err = k.captureRafflePot(ctx, pot, creatorAcc); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "could not capture pot: (%s)", err.Error())
	}

	k.SetRaffle(ctx, raffle)
	k.SetRaffleDeleteHook(ctx, types.RaffleDeleteHook{
		Denom: raffle.Denom,
		EndAt: raffle.EndAt,
	})

	return &types.MsgStartRaffleResponse{}, nil
}

func (k Keeper) captureRafflePot(ctx sdk.Context, pot sdk.Coin, creator sdk.AccAddress) error {
	//call it to make sure the account is created
	raffleAcc := k.accKeeper.GetModuleAccount(ctx, types.RaffleModuleName)
	if raffleAcc == nil {

		return fmt.Errorf("could not get module account %s ", types.RaffleModuleName)
	}

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.RaffleModuleName, sdk.NewCoins(pot))
}

func (k Keeper) accountHasCoins(ctx sdk.Context, coinsNeeded sdk.Coin, account sdk.AccAddress) bool {
	balances := k.bankKeeper.SpendableCoins(ctx, account)

	return coinsNeeded.Amount.LTE(balances.AmountOf(coinsNeeded.Denom))
}

func (k Keeper) raffleFromMsgStartRaffle(ctx sdk.Context, msg *types.MsgStartRaffle) (types.Raffle, error) {
	raffle, err := msg.ToStorageRaffle()
	if err != nil {
		return types.Raffle{}, err
	}

	raffle.Winners = 0
	currentEpoch := k.GetRaffleCurrentEpoch(ctx)
	raffle.EndAt = currentEpoch + (raffle.Duration * 24)
	raffle.TotalWon = sdk.ZeroInt().String()

	return raffle, nil
}

func (k msgServer) JoinRaffle(goCtx context.Context, msg *types.MsgJoinRaffle) (*types.MsgJoinRaffleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check denom exists
	if !k.bankKeeper.HasSupply(ctx, msg.Denom) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "denom %s does not exist", msg.Denom)
	}

	//get raffle
	raffle, found := k.GetRaffle(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "raffle not found for provided denom")
	}

	//do not allow participants to join close to expiration
	re := k.GetRaffleCurrentEpoch(ctx)
	if re > 0 && raffle.EndAt <= (re-1) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "raffle has expired")
	}

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	//get ticket price to enter the raffle
	ticketPriceInt, ok := sdk.NewIntFromString(raffle.TicketPrice)
	if !ok {
		//should never happen
		return nil, fmt.Errorf("could not get raffle ticket price")
	}

	ticketPrice := sdk.NewCoin(raffle.Denom, ticketPriceInt)
	if !k.accountHasCoins(ctx, ticketPrice, creatorAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not enough balance")
	}

	//get raffle module account
	raffleAcc := k.accKeeper.GetModuleAccount(ctx, types.RaffleModuleName)
	if raffleAcc == nil {
		return nil, fmt.Errorf("could not get module account %s ", types.RaffleModuleName)
	}

	//get raffle module balance for the raffle denom before capturing the ticket price
	currentPot := k.bankKeeper.GetBalance(ctx, raffleAcc.GetAddress(), raffle.Denom)
	if !currentPot.IsPositive() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no pot to participate to")
	}

	//capture ticket price from user account to raffle module name
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.RaffleModuleName, sdk.NewCoins(ticketPrice))
	if err != nil {
		return nil, err
	}

	participant := types.RaffleParticipant{
		Index:       k.GetParticipantCounter(ctx),
		Denom:       raffle.Denom,
		Participant: creatorAddr.String(),
		ExecuteAt:   ctx.BlockHeight() + raffleDelayHeight,
	}

	k.SetRaffleParticipant(ctx, participant)

	return &types.MsgJoinRaffleResponse{}, nil
}
