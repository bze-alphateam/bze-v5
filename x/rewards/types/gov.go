package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeActivateTradingReward = "ActivateTradingRewardProposal"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeActivateTradingReward)
}

var (
	_ govtypes.Content = &ActivateTradingRewardProposal{}
)

func NewActivateTradingRewardProposal(rewardId, title, description string) govtypes.Content {
	return &ActivateTradingRewardProposal{
		RewardId:    rewardId,
		Title:       title,
		Description: description,
	}
}

func (m *ActivateTradingRewardProposal) ProposalRoute() string { return RouterKey }

func (m *ActivateTradingRewardProposal) ProposalType() string {
	return ProposalTypeActivateTradingReward
}

func (m *ActivateTradingRewardProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	return nil
}
