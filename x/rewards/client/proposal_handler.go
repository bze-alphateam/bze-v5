package client

import (
	"github.com/bze-alphateam/bze/x/rewards/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var ActivateTradingRewardProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitActivateTradingRewardProposal)
