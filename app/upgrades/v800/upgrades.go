package v800

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const UpgradeName = "v8.0.0"

func CreateUpgradeHandler(
	cfg module.Configurator,
	mm *module.Manager,
) upgradetypes.UpgradeHandler {

	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

		//run default migrations in order to init new module's genesis and to have them in vm
		return mm.RunMigrations(ctx, cfg, vm)
	}
}
