package keeper_test

import (
	keepertest "github.com/bze-alphateam/bze/testutil/keeper"
	"github.com/bze-alphateam/bze/x/rewards/keeper"
	"github.com/bze-alphateam/bze/x/rewards/testutil"
	"github.com/bze-alphateam/bze/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite
	ctx       sdk.Context
	k         *keeper.Keeper
	msgServer types.MsgServer

	bankMock  *testutil.MockBankKeeper
	distrMock *testutil.MockDistrKeeper
	epochMock *testutil.MockEpochKeeper
	tradeMock *testutil.MockTradingKeeper
}

func (suite *IntegrationTestSuite) SetupTest() {
	t := suite.T()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBank := testutil.NewMockBankKeeper(mockCtrl)
	require.NotNil(t, mockBank)

	mockDistr := testutil.NewMockDistrKeeper(mockCtrl)
	require.NotNil(t, mockBank)

	mockTbin := testutil.NewMockTradingKeeper(mockCtrl)
	require.NotNil(t, mockBank)

	mockEpochs := testutil.NewMockEpochKeeper(mockCtrl)
	require.NotNil(t, mockBank)

	k, ctx := keepertest.RewardsKeeper(t, mockBank, mockDistr, mockTbin, mockEpochs)

	suite.ctx = ctx

	suite.k = k
	suite.msgServer = keeper.NewMsgServerImpl(*k)
	suite.bankMock = mockBank
	suite.distrMock = mockDistr
	suite.epochMock = mockEpochs
	suite.tradeMock = mockTbin
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
