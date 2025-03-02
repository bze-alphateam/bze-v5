package keeper_test

import (
	keeper2 "github.com/bze-alphateam/bze/testutil/keeper"
	"github.com/bze-alphateam/bze/x/epochs/keeper"
	"github.com/bze-alphateam/bze/x/epochs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	keeper    *keeper.Keeper
	msgServer types.MsgServer
}

func (suite *IntegrationTestSuite) SetupTest() {
	k, ctx := keeper2.EpochsKeeper(suite.T())
	suite.ctx = ctx
	suite.keeper = k
	suite.msgServer = keeper.NewMsgServerImpl(*k)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
