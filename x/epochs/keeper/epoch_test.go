package keeper_test

import (
	"time"

	"github.com/bze-alphateam/bze/x/epochs/types"
)

func (suite *IntegrationTestSuite) TestAddAndGetEpochInfo() {
	// Define an epoch info
	epochInfo := types.NewGenesisEpochInfo("test-epoch", time.Hour)

	// Attempt to add epoch info
	err := suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.NoError(err)

	// Fetch the epoch info
	fetchedEpochInfo := suite.keeper.GetEpochInfo(suite.ctx, "test-epoch")
	suite.NotEqual(types.EpochInfo{}, fetchedEpochInfo, "Fetch operation failed to retrieve the correct epoch info")

	// Validate fetched info
	suite.Equal("test-epoch", fetchedEpochInfo.Identifier, "Fetched epoch identifier mismatch")
	suite.Equal(epochInfo, fetchedEpochInfo, "Fetched epoch info mismatch")

	// Attempt to add the same epoch info should fail
	err = suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.Error(err, "Expected error when adding duplicate epoch info")
}

func (suite *IntegrationTestSuite) TestDeleteEpochInfo() {
	// Adding a test epoch to delete later
	epochInfo := types.NewGenesisEpochInfo("delete-test-epoch", time.Hour)
	suite.NoError(suite.keeper.AddEpochInfo(suite.ctx, epochInfo))

	// Delete the epoch
	suite.keeper.DeleteEpochInfo(suite.ctx, "delete-test-epoch")

	// Try to get the deleted epoch
	deletedEpochInfo := suite.keeper.GetEpochInfo(suite.ctx, "delete-test-epoch")
	suite.Equal(types.EpochInfo{}, deletedEpochInfo, "Epoch info should be empty after deletion")
}

func (suite *IntegrationTestSuite) TestIterateEpochInfo() {
	// Define an epoch info
	epochInfo := types.NewGenesisEpochInfo("test-epoch", time.Hour)

	// Attempt to add epoch info
	err := suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.NoError(err)
	epochInfo.Identifier = "another-epoch"
	err = suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.NoError(err)

	count := 0
	suite.keeper.IterateEpochInfo(suite.ctx, func(index int64, epochInfo types.EpochInfo) (stop bool) {
		count++
		return false
	})

	suite.Equal(2, count, "IterateEpochInfo did not iterate over all epochs correctly")
}

func (suite *IntegrationTestSuite) TestAllEpochInfos() {
	// Define an epoch info
	epochInfo := types.NewGenesisEpochInfo("test-epoch", time.Hour)

	// Attempt to add epoch info
	err := suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.NoError(err)
	epochInfo.Identifier = "another-epoch"
	err = suite.keeper.AddEpochInfo(suite.ctx, epochInfo)
	suite.NoError(err)

	// Fetch the epoch info
	fetchedEpochInfo := suite.keeper.GetEpochInfo(suite.ctx, "test-epoch")
	suite.NotEqual(types.EpochInfo{}, fetchedEpochInfo, "Fetch operation failed to retrieve the correct epoch info")

	allEpochs := suite.keeper.AllEpochInfos(suite.ctx)
	suite.Len(allEpochs, 2, "AllEpochInfos did not return all epoch infos")
}
