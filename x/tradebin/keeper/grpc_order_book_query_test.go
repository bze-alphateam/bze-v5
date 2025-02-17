package keeper_test

import (
	"github.com/bze-alphateam/bze/x/tradebin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestAssetMarkets_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.AssetMarkets(goCtx, nil)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestAssetMarkets_InvalidRequest_InvalidAsset() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.AssetMarkets(goCtx, &types.QueryAssetMarketsRequest{Asset: ""})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestAssetMarkets_OneMarketAsBaseDenom_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.k.SetMarket(suite.ctx, market)

	res, err := suite.k.AssetMarkets(goCtx, &types.QueryAssetMarketsRequest{Asset: denomStake})
	suite.Require().Nil(err)

	suite.Require().Empty(res.Quote)
	suite.Require().NotEmpty(res.Base)
	suite.Require().Equal(len(res.Base), 1)
	suite.Require().Equal(res.Base[0], market)
}

func (suite *IntegrationTestSuite) TestAssetMarkets_OneMarketAsQuoteDenom_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.k.SetMarket(suite.ctx, market)

	res, err := suite.k.AssetMarkets(goCtx, &types.QueryAssetMarketsRequest{Asset: denomBze})
	suite.Require().Nil(err)

	suite.Require().Empty(res.Base)
	suite.Require().NotEmpty(res.Quote)
	suite.Require().Equal(len(res.Quote), 1)
	suite.Require().Equal(res.Quote[0], market)
}

func (suite *IntegrationTestSuite) TestAssetMarkets_MoreMarket_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	fakeDenom1 := types.Market{
		Base:    "fake1",
		Quote:   denomBze,
		Creator: "fake_addr",
	}
	fakeDenom2 := types.Market{
		Base:    denomBze,
		Quote:   "fake2",
		Creator: "fake_addr",
	}
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetMarket(suite.ctx, fakeDenom1)
	suite.k.SetMarket(suite.ctx, fakeDenom2)

	res, err := suite.k.AssetMarkets(goCtx, &types.QueryAssetMarketsRequest{Asset: denomBze})
	suite.Require().Nil(err)

	suite.Require().Equal(len(res.Base), 1)
	suite.Require().Equal(res.Base[0], fakeDenom2)

	suite.Require().Equal(len(res.Quote), 2)
	suite.Require().Equal(res.Quote[0], fakeDenom1)
	suite.Require().Equal(res.Quote[1], market)
}

func (suite *IntegrationTestSuite) TestMarketAggregatedOrders_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.MarketAggregatedOrders(goCtx, nil)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarketAggregatedOrders_InvalidMarket() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.MarketAggregatedOrders(goCtx, &types.QueryMarketAggregatedOrdersRequest{Market: ""})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarketAggregatedOrders_InvalidOrderType() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.k.SetMarket(suite.ctx, market)

	_, err := suite.k.MarketAggregatedOrders(goCtx, &types.QueryMarketAggregatedOrdersRequest{Market: getMarketId(), OrderType: "yeahsureinvalidordertype"})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarketAggregatedOrders_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.k.SetMarket(suite.ctx, market)
	aggBuy := types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeBuy,
		Amount:    "100",
		Price:     "2",
	}

	aggSell := types.AggregatedOrder{
		MarketId:  getMarketId(),
		OrderType: types.OrderTypeSell,
		Amount:    "92233720368547758061",
		Price:     "232183213232131221313",
	}

	suite.k.SetAggregatedOrder(suite.ctx, aggBuy)
	suite.k.SetAggregatedOrder(suite.ctx, aggSell)

	res, err := suite.k.MarketAggregatedOrders(goCtx, &types.QueryMarketAggregatedOrdersRequest{Market: getMarketId(), OrderType: types.OrderTypeBuy})
	suite.Require().Nil(err)
	suite.Require().Len(res.List, 1)
	suite.Require().Equal(res.List[0], aggBuy)

	res, err = suite.k.MarketAggregatedOrders(goCtx, &types.QueryMarketAggregatedOrdersRequest{Market: getMarketId(), OrderType: types.OrderTypeSell})
	suite.Require().Nil(err)
	suite.Require().Len(res.List, 1)
	suite.Require().Equal(res.List[0], aggSell)
}

func (suite *IntegrationTestSuite) TestMarketAll_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.MarketAll(goCtx, nil)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarketAll_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	fakeDenom1 := types.Market{
		Base:    "fake1",
		Quote:   denomBze,
		Creator: "fake_addr",
	}
	fakeDenom2 := types.Market{
		Base:    denomBze,
		Quote:   "fake2",
		Creator: "fake_addr",
	}
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetMarket(suite.ctx, fakeDenom1)
	suite.k.SetMarket(suite.ctx, fakeDenom2)

	res, err := suite.k.MarketAll(goCtx, &types.QueryAllMarketRequest{})
	suite.Require().Nil(err)
	suite.Require().NotEmpty(res.Market)
	suite.Require().Equal(len(res.Market), 3)
}

func (suite *IntegrationTestSuite) TestMarket_InvalidRequest() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.Market(goCtx, nil)
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarket_InvalidArguments() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.k.Market(goCtx, &types.QueryGetMarketRequest{Base: denomStake, Quote: ""})
	suite.Require().NotNil(err)

	_, err = suite.k.Market(goCtx, &types.QueryGetMarketRequest{Base: "", Quote: ""})
	suite.Require().NotNil(err)
}

func (suite *IntegrationTestSuite) TestMarket_Success() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	fakeDenom1 := types.Market{
		Base:    "fake1",
		Quote:   denomBze,
		Creator: "fake_addr",
	}
	suite.k.SetMarket(suite.ctx, market)
	suite.k.SetMarket(suite.ctx, fakeDenom1)

	res, err := suite.k.Market(goCtx, &types.QueryGetMarketRequest{Base: denomStake, Quote: denomBze})
	suite.Require().Nil(err)
	suite.Require().Equal(res.Market, market)

	res, err = suite.k.Market(goCtx, &types.QueryGetMarketRequest{Base: "fake1", Quote: denomBze})
	suite.Require().Nil(err)
	suite.Require().Equal(res.Market, fakeDenom1)
}
