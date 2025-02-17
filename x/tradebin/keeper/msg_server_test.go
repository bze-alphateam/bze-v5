package keeper_test

import (
	"fmt"
	"github.com/bze-alphateam/bze/x/tradebin/keeper"
	"github.com/bze-alphateam/bze/x/tradebin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/rand"
	"go.uber.org/mock/gomock"
	"strconv"
	"time"
)

func (suite *IntegrationTestSuite) TestCancelOrder_MarketNotFound() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CancelOrder(goCtx, &types.MsgCancelOrder{
		Creator:   "me",
		MarketId:  "me/me",
		OrderId:   "",
		OrderType: "",
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCancelOrder_OrderNotFound() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CancelOrder(goCtx, &types.MsgCancelOrder{
		Creator:   "me",
		MarketId:  getMarketId(),
		OrderId:   "123",
		OrderType: "dsa",
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCancelOrder_Unauthorized() {
	suite.k.SetMarket(suite.ctx, market)
	order := suite.k.NewOrder(suite.ctx, types.Order{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Amount:    "102",
		Price:     "1",
		Owner:     "me",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CancelOrder(goCtx, &types.MsgCancelOrder{
		Creator:   "not_me",
		MarketId:  getMarketId(),
		OrderId:   order.Id,
		OrderType: types.OrderTypeBuy,
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCancelOrder_CancelBuy_Success() {
	suite.k.SetMarket(suite.ctx, market)
	order := suite.k.NewOrder(suite.ctx, types.Order{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Amount:    "102000",
		Price:     "1",
		Owner:     "me",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CancelOrder(goCtx, &types.MsgCancelOrder{
		Creator:   "me",
		MarketId:  getMarketId(),
		OrderId:   order.Id,
		OrderType: types.OrderTypeBuy,
	})
	suite.Require().Nil(err)

	qms := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qms, 1)

	suite.Require().Equal(qms[0].MarketId, order.MarketId)
	suite.Require().Equal(qms[0].MessageType, types.MessageTypeCancel)
	suite.Require().Equal(qms[0].OrderId, order.Id)
	suite.Require().Equal(qms[0].OrderType, order.OrderType)
	suite.Require().Equal(qms[0].Owner, order.Owner)
}

func (suite *IntegrationTestSuite) TestCancelOrder_CancelSell_Success() {
	suite.k.SetMarket(suite.ctx, market)
	order := suite.k.NewOrder(suite.ctx, types.Order{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Amount:    "10200021",
		Price:     "1000",
		Owner:     "me",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CancelOrder(goCtx, &types.MsgCancelOrder{
		Creator:   "me",
		MarketId:  getMarketId(),
		OrderId:   order.Id,
		OrderType: types.OrderTypeSell,
	})
	suite.Require().Nil(err)

	qms := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qms, 1)

	suite.Require().Equal(qms[0].MarketId, order.MarketId)
	suite.Require().Equal(qms[0].MessageType, types.MessageTypeCancel)
	suite.Require().Equal(qms[0].OrderId, order.Id)
	suite.Require().Equal(qms[0].OrderType, order.OrderType)
	suite.Require().Equal(qms[0].Owner, order.Owner)
}

func (suite *IntegrationTestSuite) TestCreateMarket_InvalidDenom() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	//empty base denom
	_, err := suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    "",
		Quote:   denomBze,
	})
	suite.Require().NotNil(err)

	//empty quote denom
	_, err = suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    denomBze,
		Quote:   "",
	})
	suite.Require().NotNil(err)

	//same denom for both
	_, err = suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    denomBze,
		Quote:   denomBze,
	})
	suite.Require().NotNil(err)

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), gomock.AnyOf(denomStake, denomBze)).Return(false).Times(1)
	//denom has no supply
	_, err = suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    denomStake,
		Quote:   denomBze,
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateMarket_MarketAlreadyExist() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    denomStake,
		Quote:   denomBze,
	})
	suite.Require().NotNil(err)

	_, err = suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: "me",
		Base:    denomBze,
		Quote:   denomStake,
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateMarket_NotEnoughCoinsForFee() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomStake).Return(true).Times(1)
	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Return(true).Times(1)
	//expect market fee to be paid
	suite.distrMock.EXPECT().
		FundCommunityPool(gomock.Any(), sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(25000000000))), addr1).
		Times(1).
		Return(fmt.Errorf("not enough balance"))

	_, err := suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: addr1.String(),
		Base:    denomStake,
		Quote:   denomBze,
	})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestCreateMarket_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomStake).Return(true).Times(1)
	suite.bankMock.EXPECT().HasSupply(gomock.Any(), denomBze).Return(true).Times(1)
	//expect market fee to be paid
	suite.distrMock.EXPECT().
		FundCommunityPool(gomock.Any(), sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(25000000000))), addr1).
		Times(1).
		Return(nil)

	newMarket := types.Market{
		Creator: addr1.String(),
		Base:    denomStake,
		Quote:   denomBze,
	}
	_, err := suite.msgServer.CreateMarket(goCtx, &types.MsgCreateMarket{
		Creator: addr1.String(),
		Base:    denomStake,
		Quote:   denomBze,
	})

	suite.Require().Nil(err)
	storageMarket, ok := suite.k.GetMarket(suite.ctx, newMarket.Base, newMarket.Quote)
	suite.Require().True(ok)
	suite.Require().Equal(newMarket, storageMarket)
}

func (suite *IntegrationTestSuite) TestCreateOrder_InvalidAmount() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CreateOrder(goCtx, &types.MsgCreateOrder{
		Amount: "hdsihdshdshids",
		Price:  "1",
	})

	suite.Require().NotNil(err)
	suite.Require().Contains(err.Error(), "amount could not be converted to Int")
}

func (suite *IntegrationTestSuite) TestCreateOrder_AmountTooLow() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CreateOrder(goCtx, &types.MsgCreateOrder{
		Amount: "1",
		Price:  "1",
	})

	suite.Require().NotNil(err)
	//amount should be bigger than
	suite.Require().Contains(err.Error(), "amount should be bigger than")
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketNotFound() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CreateOrder(goCtx, &types.MsgCreateOrder{
		Amount:   "1000000",
		Price:    "1",
		MarketId: "notamarket/notatall",
	})

	suite.Require().NotNil(err)
	//market id
	suite.Require().Contains(err.Error(), "market id")
}

func (suite *IntegrationTestSuite) TestCreateOrder_InvalidOrderType() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	addr1 := sdk.AccAddress("addr1_______________")
	_, err := suite.msgServer.CreateOrder(goCtx, &types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "1",
		MarketId:  getMarketId(),
		OrderType: "notatype",
		Creator:   addr1.String(),
	})

	suite.Require().NotNil(err)
	suite.Require().Contains(err.Error(), "order type")
}

func (suite *IntegrationTestSuite) TestCreateOrder_InvalidCreator() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.CreateOrder(goCtx, &types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "1",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   "notanaddress",
	})

	suite.Require().NotNil(err)
	suite.Require().Contains(err.Error(), "bech32")
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketMaker_Buy_Success_ZeroDust() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	makerFee, err := sdk.ParseCoinsNormalized(params.MarketMakerFee)
	suite.Require().NoError(err)
	paidCoins := sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(2_000_000)))

	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.MakerFeeDestination, makerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)

	orderMsg := types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "2",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Nil(err)

	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, orderMsg.Amount)
	suite.Require().Equal(qmList[0].OrderType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].MessageType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].Price, orderMsg.Price)
	suite.Require().Equal(qmList[0].Owner, orderMsg.Creator)

	_, ok := suite.k.GetUserDust(suite.ctx, addr1.String(), market.Quote)
	suite.Require().False(ok)

	_, ok = suite.k.GetUserDust(suite.ctx, addr1.String(), market.Base)
	suite.Require().False(ok)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketTaker_Buy_Success_ZeroDust() {
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetAggregatedOrder(suite.ctx, types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Amount:    "10000",
		Price:     "2",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)
	paidCoins := sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(2_000_000)))

	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)

	orderMsg := types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "2",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)

	suite.Require().Nil(err)
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, orderMsg.Amount)
	suite.Require().Equal(qmList[0].OrderType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].MessageType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].Price, orderMsg.Price)
	suite.Require().Equal(qmList[0].Owner, orderMsg.Creator)

	_, ok := suite.k.GetUserDust(suite.ctx, addr1.String(), market.Quote)
	suite.Require().False(ok)

	_, ok = suite.k.GetUserDust(suite.ctx, addr1.String(), market.Base)
	suite.Require().False(ok)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketTaker_Buy_Success_WithDust() {
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetAggregatedOrder(suite.ctx, types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Amount:    "100000",
		Price:     "0.02331",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")
	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)
	paidCoins := sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(3)))

	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
	orderMsg := types.MsgCreateOrder{
		Amount:    "87",
		Price:     "0.02331",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Nil(err)
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, orderMsg.Amount)
	suite.Require().Equal(qmList[0].OrderType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].MessageType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].Price, orderMsg.Price)
	suite.Require().Equal(qmList[0].Owner, orderMsg.Creator)

	udQuote, ok := suite.k.GetUserDust(suite.ctx, addr1.String(), market.Quote)
	suite.Require().True(ok)
	suite.Require().Equal(udQuote.Denom, market.Quote)
	suite.Require().Equal(udQuote.Owner, addr1.String())
	suite.Require().EqualValues(udQuote.Amount, "0.972030000000000000")

	_, ok = suite.k.GetUserDust(suite.ctx, addr1.String(), market.Base)
	suite.Require().False(ok)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketMaker_Sell_Success() {
	suite.k.SetMarket(suite.ctx, market)
	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketMakerFee)
	suite.Require().NoError(err)
	paidCoins := sdk.NewCoins(sdk.NewCoin(denomStake, sdk.NewInt(1000000)))

	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
	orderMsg := types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "1",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)

	suite.Require().Nil(err)
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, orderMsg.Amount)
	suite.Require().Equal(qmList[0].OrderType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].MessageType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].Price, orderMsg.Price)
	suite.Require().Equal(qmList[0].Owner, orderMsg.Creator)

	//should never have dust on sell orders
	_, ok := suite.k.GetUserDust(suite.ctx, addr1.String(), market.Quote)
	suite.Require().False(ok)

	_, ok = suite.k.GetUserDust(suite.ctx, addr1.String(), market.Base)
	suite.Require().False(ok)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketTaker_Sell_Success() {
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetAggregatedOrder(suite.ctx, types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Amount:    "10000",
		Price:     "1",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)
	paidCoins := sdk.NewCoins(sdk.NewCoin(denomStake, sdk.NewInt(1000000)))

	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
	orderMsg := types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "1",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Nil(err)
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, orderMsg.Amount)
	suite.Require().Equal(qmList[0].OrderType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].MessageType, orderMsg.OrderType)
	suite.Require().Equal(qmList[0].Price, orderMsg.Price)
	suite.Require().Equal(qmList[0].Owner, orderMsg.Creator)

	//TODO: better testing in case fee destination is changed to community pool

	//should never have dust on sell orders
	_, ok := suite.k.GetUserDust(suite.ctx, addr1.String(), market.Quote)
	suite.Require().False(ok)

	_, ok = suite.k.GetUserDust(suite.ctx, addr1.String(), market.Base)
	suite.Require().False(ok)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketTaker_StressBalance() {
	suite.k.SetMarket(suite.ctx, market)
	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	//create initial random markets
	testDenom1 := "test1"
	testDenom2 := "test2"
	market1 := types.Market{
		Base:    testDenom1,
		Quote:   testDenom2,
		Creator: "addr1",
	}
	market2 := types.Market{
		Base:    denomStake,
		Quote:   testDenom2,
		Creator: "addr1",
	}
	market3 := types.Market{
		Base:    denomBze,
		Quote:   testDenom1,
		Creator: "addr1",
	}

	market4 := types.Market{
		Base:    denomBze,
		Quote:   testDenom2,
		Creator: "addr1",
	}
	marketsMap := make(map[string]types.Market)
	marketsMap[types.CreateMarketId(market.Base, market.Quote)] = market
	marketsMap[types.CreateMarketId(market1.Base, market1.Quote)] = market1
	marketsMap[types.CreateMarketId(market2.Base, market2.Quote)] = market2
	marketsMap[types.CreateMarketId(market3.Base, market3.Quote)] = market3
	marketsMap[types.CreateMarketId(market4.Base, market4.Quote)] = market4
	suite.k.SetMarket(suite.ctx, market1)
	suite.k.SetMarket(suite.ctx, market2)
	suite.k.SetMarket(suite.ctx, market3)
	suite.k.SetMarket(suite.ctx, market4)

	var creators []string
	for i := 0; i < 10; i++ {
		addr1 := sdk.AccAddress(fmt.Sprintf("addr%d_______________", i))
		creators = append(creators, addr1.String())
	}

	allPaid := sdk.NewCoins()
	for i := 0; i < 5; i++ {
		_, paid := suite.randomOrderCreateMessages(suite.randomNumber(1000), creators, market)
		allPaid = allPaid.Add(paid...)
		_, paid = suite.randomOrderCreateMessages(suite.randomNumber(1000), creators, market1)
		allPaid = allPaid.Add(paid...)
		_, paid = suite.randomOrderCreateMessages(suite.randomNumber(1000), creators, market2)
		allPaid = allPaid.Add(paid...)
		_, paid = suite.randomOrderCreateMessages(suite.randomNumber(1000), creators, market3)
		allPaid = allPaid.Add(paid...)
		_, paid = suite.randomOrderCreateMessages(suite.randomNumber(1000), creators, market4)
		allPaid = allPaid.Add(paid...)

		suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx sdk.Context, moduleName string, recipient sdk.AccAddress, coins sdk.Coins) error {
				allPaid = allPaid.Sub(coins...)
				return nil
			}).AnyTimes()
		//suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		engine.ProcessQueueMessages(suite.ctx)
	}

	allOrders := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(allOrders)
	amounts := sdk.NewCoins()
	//check module balance is equal to the coins found
	foundOrderTypesByPrice := make(map[string]string)
	aggregatedOrders := make(map[string]types.AggregatedOrder)
	for _, or := range allOrders {
		//search to see if an opposite order with the same price already exists
		//if so it means the engine failed to match them when they were processed
		oppositeKey := fmt.Sprintf("%s_%s_%s", or.MarketId, or.Price, types.TheOtherOrderType(or.OrderType))
		id, ok := foundOrderTypesByPrice[oppositeKey]
		suite.Require().False(ok, fmt.Sprintf("order [%s] has the same price [%s] as [%s]", id, or.Price, or.Id))

		//save locally to check later
		key := fmt.Sprintf("%s_%s_%s", or.MarketId, or.Price, or.OrderType)
		foundOrderTypesByPrice[key] = or.Id
		amtInt, ok := sdk.NewIntFromString(or.Amount)
		suite.Require().True(ok)
		pickMarket, ok := marketsMap[or.MarketId]
		suite.Require().True(ok)
		coins, _, err := suite.k.GetOrderSdkCoin(or.OrderType, or.Price, amtInt, &pickMarket)
		suite.Require().NoError(err)

		amounts = amounts.Add(coins)
		//save in aggregated orders map so we can check it later
		agg, found := aggregatedOrders[key]
		if !found {
			agg = types.AggregatedOrder{
				MarketId:  or.MarketId,
				OrderType: or.OrderType,
				Amount:    "0",
				Price:     or.Price,
			}
		}
		aggAmt, ok := sdk.NewIntFromString(agg.Amount)
		suite.Require().True(ok)
		aggAmt = aggAmt.Add(amtInt)
		agg.Amount = aggAmt.String()
		aggregatedOrders[key] = agg
	}

	for _, agg := range aggregatedOrders {
		foundAgg, ok := suite.k.GetAggregatedOrder(suite.ctx, agg.MarketId, agg.OrderType, agg.Price)
		suite.Require().True(ok)
		suite.Require().Equal(foundAgg.Amount, agg.Amount)
	}

	suite.Require().Equal(allPaid, amounts)
}

func (suite *IntegrationTestSuite) TestCreateOrder_MarketTaker_CheckPrice_Fail() {
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetAggregatedOrder(suite.ctx, types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Amount:    "10000",
		Price:     "1",
	})
	suite.k.SetAggregatedOrder(suite.ctx, types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Amount:    "10000",
		Price:     "5",
	})

	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 := sdk.AccAddress("addr1_______________")

	//check price error on sell order
	orderMsg := types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "0.5",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Creator:   addr1.String(),
	}
	_, err := suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Error(err)
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 0)

	//check price error on buy
	orderMsg = types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "5.1",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Error(err)

	qmList = suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 0)

	//fee should be captured
	params := suite.k.GetParams(suite.ctx)
	makerFee, err := sdk.ParseCoinsNormalized(params.MarketMakerFee)
	suite.Require().NoError(err)

	paidCoins := sdk.NewCoins(sdk.NewCoin(denomBze, sdk.NewInt(4000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.MakerFeeDestination, makerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
	//add 2 new orders in order to check message queue price validator
	orderMsg = types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "4",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().NoError(err)

	paidCoins = sdk.NewCoins(sdk.NewCoin(denomStake, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.MakerFeeDestination, makerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
	orderMsg = types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "4.5",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().NoError(err)

	orderMsg = types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "3.5",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Error(err)

	orderMsg = types.MsgCreateOrder{
		Amount:    "1000000",
		Price:     "4.55",
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Creator:   addr1.String(),
	}
	_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) randomOrderCreateMessages(count int, creators []string, market types.Market) ([]types.MsgCreateOrder, sdk.Coins) {
	var msgs []types.MsgCreateOrder
	orderTypes := []string{types.OrderTypeBuy, types.OrderTypeSell}
	goCtx := sdk.WrapSDKContext(suite.ctx)
	allPaid := sdk.NewCoins()
	const ExecPrice = 15
	for i := 0; i < count; i++ {
		randomOrderType := orderTypes[i%2]
		randomPrice := suite.randomNumber(24) + 1 //make sure it's always higher than 0
		if randomOrderType == types.OrderTypeBuy && randomPrice > ExecPrice {
			randomPrice = ExecPrice
		} else if randomOrderType == types.OrderTypeSell && randomPrice < ExecPrice {
			randomPrice = ExecPrice
		}
		randomPriceStr := strconv.Itoa(randomPrice)
		minAmount := keeper.CalculateMinAmount(randomPriceStr)
		orderAmount := minAmount.AddRaw(int64(suite.randomNumber(1000)))
		orderMsg := types.MsgCreateOrder{
			Amount:    orderAmount.String(),
			Price:     randomPriceStr,
			MarketId:  types.CreateMarketId(market.Base, market.Quote),
			OrderType: randomOrderType,
			Creator:   creators[suite.randomNumber(len(creators))],
		}

		//fee should be captured
		params := suite.k.GetParams(suite.ctx)
		takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
		suite.Require().NoError(err)
		makerFee, err := sdk.ParseCoinsNormalized(params.MarketMakerFee)
		suite.Require().NoError(err)
		var paidCoins sdk.Coins
		if orderMsg.OrderType == types.OrderTypeBuy {
			paidCoins = sdk.NewCoins(sdk.NewCoin(market.Quote, orderAmount.MulRaw(int64(randomPrice))))
		} else {
			paidCoins = sdk.NewCoins(sdk.NewCoin(market.Base, orderAmount))
		}
		allPaid = allPaid.Add(paidCoins...)

		creatorAcc, err := sdk.AccAddressFromBech32(orderMsg.Creator)
		suite.Require().NoError(err)
		suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), creatorAcc, gomock.AnyOf(params.TakerFeeDestination, params.MakerFeeDestination), gomock.AnyOf(takerFee, makerFee)).Return(nil).Times(1)
		suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), creatorAcc, types.ModuleName, paidCoins).Return(nil).Times(1)

		_, err = suite.msgServer.CreateOrder(goCtx, &orderMsg)
		suite.Require().NoError(err)
	}

	return msgs, allPaid
}

func (suite *IntegrationTestSuite) randomNumber(to int) int {
	// Seed the random number generator to get different results each run
	// Uses the current time as the seed
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and to
	num := rand.Intn(to)

	return num
}

func (suite *IntegrationTestSuite) msgOrderFillSetup(orderType string) (allPrices []string, addr1, addr2 sdk.AccAddress) {
	suite.k.SetMarket(suite.ctx, market)

	goCtx := sdk.WrapSDKContext(suite.ctx)
	addr1 = sdk.AccAddress("addr1_______________")
	addr2 = sdk.AccAddress("addr2_______________")
	params := suite.k.GetParams(suite.ctx)
	makerFee, err := sdk.ParseCoinsNormalized(params.MarketMakerFee)
	suite.Require().NoError(err)

	initialPrice := sdk.ZeroInt()
	staticAmount := "1000000"
	for i := 1; i <= 10; i++ {
		initialPrice = initialPrice.AddRaw(int64(i))
		orderMsg := types.MsgCreateOrder{
			Amount:    staticAmount,
			Price:     initialPrice.String(),
			MarketId:  getMarketId(),
			OrderType: orderType,
			Creator:   addr1.String(),
		}

		var paidCoins sdk.Coins
		if orderType == types.OrderTypeBuy {
			paidCoins = sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000).Mul(initialPrice)))
		} else {
			paidCoins = sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))
		}

		suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, params.MakerFeeDestination, makerFee).Return(nil).Times(1)
		suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr1, types.ModuleName, paidCoins).Return(nil).Times(1)
		_, err := suite.msgServer.CreateOrder(goCtx, &orderMsg)
		suite.Require().NoError(err)

		allPrices = append(allPrices, initialPrice.String())
	}

	return
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneOrderPartialFill_Sell() {
	_, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "500000",
				Price:  "1",
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)

	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	receivedCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, receivedCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, paidCoins).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 10)
	remainingOrderChecked := false
	for _, order := range all {
		//let's find our order and check if it contains the correct amount
		if order.Price != "1" {
			continue
		}

		suite.Require().EqualValues(order.Amount, fmt.Sprintf("%d", 1000000-500000))
		remainingOrderChecked = true
		break
	}

	suite.Require().True(remainingOrderChecked)
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneOrderPartialFill_Buy() {
	_, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)
	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "500000",
				Price:  "1",
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	receivedCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(500000)))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, receivedCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, paidCoins).Return(nil).Times(1)

	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 10)
	remainingOrderChecked := false
	for _, order := range all {
		//let's find our order and check if it contains the correct amount
		if order.Price != "1" {
			continue
		}

		suite.Require().EqualValues(order.Amount, fmt.Sprintf("%d", 1000000-500000))
		remainingOrderChecked = true
		break
	}

	suite.Require().True(remainingOrderChecked)
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneOrderFullFill_Sell() {
	_, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  "1",
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	receivedCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, receivedCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, paidCoins).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	for _, order := range all {
		//let's find our order and check if it contains the correct amount
		suite.Require().NotEqual(order.Price, "1")
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneOrderFullFill_Buy() {
	_, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  "1",
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	receivedCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, receivedCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, paidCoins).Return(nil).Times(1)

	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	for _, order := range all {
		//let's find our order and check if it contains the correct amount
		suite.Require().NotEqual(order.Price, "1000000")
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoOrdersOnePartialFill_Sell() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  allPrices[0],
			},
			{
				Amount: "500000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(2500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	//they receive in multiple bank calls: one for each message
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	remainingOrderChecked := false
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		//let's find our order and check if it contains the correct amount
		if order.Price != allPrices[1] {
			continue
		}

		suite.Require().EqualValues(order.Amount, fmt.Sprintf("%d", 1000000-500000))
		remainingOrderChecked = true
		break
	}

	suite.Require().True(remainingOrderChecked)
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoOrdersOnePartialFill_Buy() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  allPrices[0],
			},
			{
				Amount: "500000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	agg := suite.k.GetAllAggregatedOrder(suite.ctx)
	_ = agg
	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	//they receive in multiple bank calls: one for each message
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1500000)))).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	remainingOrderChecked := false
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		//let's find our order and check if it contains the correct amount
		if order.Price != allPrices[1] {
			continue
		}

		suite.Require().EqualValues(order.Amount, fmt.Sprintf("%d", 1000000-500000))
		remainingOrderChecked = true
		break
	}

	suite.Require().True(remainingOrderChecked)
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoFullyFilledOrders_Buy() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  allPrices[0],
			},
			{
				Amount: "1000000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(2000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	agg := suite.k.GetAllAggregatedOrder(suite.ctx)
	_ = agg
	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	//they receive in multiple bank calls: one for each message
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(3000000)))).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 8)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		suite.Require().NotEqual(order.Price, allPrices[1])
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoFullyFilledOrders_Sell() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1000000",
				Price:  allPrices[0],
			},
			{
				Amount: "1000000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(4000000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	//they receive in multiple bank calls: one for each message
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(3000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 8)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		suite.Require().NotEqual(order.Price, allPrices[1])
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneFullyFilledOrderWithExtraAmount_Buy() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	engine.ProcessQueueMessages(suite.ctx)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1500000",
				Price:  allPrices[0],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	receivedCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, receivedCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)

	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_OneFullyFilledOrderWithExtraAmount_Sell() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)

	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "1500000",
				Price:  allPrices[0],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 1)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)

	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)

	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 9)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoFullyFilledOrdersWithExtraAmounts_Buy() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders: []*types.FillOrderItem{
			{
				Amount: "2000000",
				Price:  allPrices[0],
			},
			{
				Amount: "1500000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(3500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.MakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeSell)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillBuy)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(3000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)

	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 8)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		suite.Require().NotEqual(order.Price, allPrices[1])
		suite.Require().Equal(order.Amount, "1000000")
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_TwoFullyFilledOrdersWithExtraAmounts_Sell() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	//Let's fill 1 order with amount lower than the available order
	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders: []*types.FillOrderItem{
			{
				Amount: "2000000",
				Price:  allPrices[0],
			},
			{
				Amount: "1500000",
				Price:  allPrices[1],
			},
		},
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(6500000)))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	//check that the new message is saved accordingly to queue messages storage
	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, 2)
	suite.Require().Equal(qmList[0].MarketId, getMarketId())
	suite.Require().Equal(qmList[0].Amount, fillOrder.Orders[0].Amount)
	suite.Require().Equal(qmList[0].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[0].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[0].Price, fillOrder.Orders[0].Price)
	suite.Require().Equal(qmList[0].Owner, fillOrder.Creator)
	suite.Require().Equal(qmList[1].MarketId, getMarketId())
	suite.Require().Equal(qmList[1].Amount, fillOrder.Orders[1].Amount)
	suite.Require().Equal(qmList[1].OrderType, types.OrderTypeBuy)
	suite.Require().Equal(qmList[1].MessageType, types.MessageTypeFillSell)
	suite.Require().Equal(qmList[1].Price, fillOrder.Orders[1].Price)
	suite.Require().Equal(qmList[1].Owner, fillOrder.Creator)

	//they receive in multiple bank calls: one for each message
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Base, sdk.NewInt(1000000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(1500000)))).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, sdk.NewCoins(sdk.NewCoin(market.Quote, sdk.NewInt(3000000)))).Return(nil).Times(1)
	engine.ProcessQueueMessages(suite.ctx)

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().NotEmpty(all)
	//the same number of orders should be present (we filled only half of first order)
	suite.Require().Len(all, 8)
	for _, order := range all {
		suite.Require().NotEqual(order.Price, allPrices[0])
		suite.Require().NotEqual(order.Price, allPrices[1])
		suite.Require().Equal(order.Amount, "1000000")
	}
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_FillAllWithExtraAmounts_Sell() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeSell)
	suite.Require().NotEmpty(allPrices)
	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)

	engine.ProcessQueueMessages(suite.ctx)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Orders:    []*types.FillOrderItem{},
	}

	quoteAmount := sdk.ZeroInt()
	baseAmount := sdk.ZeroInt()
	expectedQuoteAmount := sdk.ZeroInt()
	expectedBaseAmount := sdk.ZeroInt()
	for _, pr := range allPrices {
		randAmount := 1000000 + suite.randomNumber(999999)
		fillOrder.Orders = append(fillOrder.Orders, &types.FillOrderItem{
			Price:  pr,
			Amount: fmt.Sprintf("%d", randAmount),
		})

		price, ok := sdk.NewIntFromString(pr)
		suite.Require().True(ok)
		baseAmount = baseAmount.AddRaw(int64(randAmount))
		quoteAmount = quoteAmount.Add(price.MulRaw(int64(randAmount)))

		expectedQuoteAmount = expectedQuoteAmount.Add(price.MulRaw(1000000))
		expectedBaseAmount = expectedBaseAmount.AddRaw(1000000)
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Quote, quoteAmount))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, len(allPrices))

	//they receive in multiple bank calls: one for each FillOrderItem
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, gomock.Any()).Return(nil).Times(len(allPrices))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, gomock.Any()).
		DoAndReturn(func(ctx sdk.Context, moduleName string, recipient sdk.AccAddress, coins sdk.Coins) error {
			for _, c := range coins {
				if c.Denom == market.Base {
					expectedBaseAmount = expectedBaseAmount.Sub(c.Amount)
				} else if c.Denom == market.Quote {
					quoteAmount = quoteAmount.Sub(c.Amount)
				}
			}

			return nil
		}).Times(len(allPrices) * 2)

	engine.ProcessQueueMessages(suite.ctx)

	suite.Require().True(expectedBaseAmount.IsZero())
	suite.Require().True(quoteAmount.Equal(expectedQuoteAmount))

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().Empty(all)
}

func (suite *IntegrationTestSuite) TestMsgOrderFill_FillAllWithExtraAmounts_Buy() {
	allPrices, addr1, addr2 := suite.msgOrderFillSetup(types.OrderTypeBuy)
	suite.Require().NotEmpty(allPrices)

	goCtx := sdk.WrapSDKContext(suite.ctx)

	engine, err := keeper.NewProcessingEngine(suite.k, suite.bankMock, suite.k.Logger(suite.ctx))
	suite.Require().Nil(err)
	params := suite.k.GetParams(suite.ctx)
	takerFee, err := sdk.ParseCoinsNormalized(params.MarketTakerFee)
	suite.Require().NoError(err)

	engine.ProcessQueueMessages(suite.ctx)

	fillOrder := types.MsgFillOrders{
		Creator:   addr2.String(),
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Orders:    []*types.FillOrderItem{},
	}

	quoteAmount := sdk.ZeroInt()
	baseAmount := sdk.ZeroInt()
	expectedQuoteAmount := sdk.ZeroInt()
	expectedBaseAmount := sdk.ZeroInt()
	for _, pr := range allPrices {
		randAmount := 1000000 + suite.randomNumber(999999)
		fillOrder.Orders = append(fillOrder.Orders, &types.FillOrderItem{
			Price:  pr,
			Amount: fmt.Sprintf("%d", randAmount),
		})

		price, ok := sdk.NewIntFromString(pr)
		suite.Require().True(ok)
		baseAmount = baseAmount.AddRaw(int64(randAmount))
		quoteAmount = quoteAmount.Add(price.MulRaw(int64(randAmount)))

		expectedQuoteAmount = expectedQuoteAmount.Add(price.MulRaw(1000000))
		expectedBaseAmount = expectedBaseAmount.AddRaw(1000000)
	}

	paidCoins := sdk.NewCoins(sdk.NewCoin(market.Base, baseAmount))
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, types.ModuleName, paidCoins).Return(nil).Times(1)
	suite.bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), addr2, params.TakerFeeDestination, takerFee).Return(nil).Times(1)
	_, err = suite.msgServer.FillOrders(goCtx, &fillOrder)
	suite.Require().NoError(err)

	qmList := suite.k.GetAllQueueMessage(suite.ctx)
	suite.Require().Len(qmList, len(allPrices))

	//they receive in multiple bank calls: one for each FillOrderItem
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr1, gomock.Any()).Return(nil).Times(len(allPrices))
	suite.bankMock.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, addr2, gomock.Any()).
		DoAndReturn(func(ctx sdk.Context, moduleName string, recipient sdk.AccAddress, coins sdk.Coins) error {
			for _, c := range coins {
				if c.Denom == market.Base {
					baseAmount = baseAmount.Sub(c.Amount)
				} else if c.Denom == market.Quote {
					expectedQuoteAmount = expectedQuoteAmount.Sub(c.Amount)
				}
			}

			return nil
		}).Times(len(allPrices) * 2)

	engine.ProcessQueueMessages(suite.ctx)

	suite.Require().True(expectedQuoteAmount.IsZero())
	suite.Require().True(baseAmount.Equal(expectedBaseAmount))

	all := suite.k.GetAllOrder(suite.ctx)
	suite.Require().Empty(all)
}
