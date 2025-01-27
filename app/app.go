package app

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	"github.com/bze-alphateam/bze/app/openapi"
	v711 "github.com/bze-alphateam/bze/app/upgrades/v711"
	"github.com/bze-alphateam/bze/x/epochs"
	epochskeeper "github.com/bze-alphateam/bze/x/epochs/keeper"
	epochstypes "github.com/bze-alphateam/bze/x/epochs/types"
	"github.com/bze-alphateam/bze/x/rewards"
	rewardsclient "github.com/bze-alphateam/bze/x/rewards/client"
	rewardskeeper "github.com/bze-alphateam/bze/x/rewards/keeper"
	rewardstypes "github.com/bze-alphateam/bze/x/rewards/types"
	"github.com/bze-alphateam/bze/x/tradebin"
	tradebinkeeper "github.com/bze-alphateam/bze/x/tradebin/keeper"
	tradebintypes "github.com/bze-alphateam/bze/x/tradebin/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/bze-alphateam/bze/x/tokenfactory"
	tokenfactorykeeper "github.com/bze-alphateam/bze/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/bze-alphateam/bze/x/tokenfactory/types"

	"io"
	"net/http"
	"os"
	"path/filepath"

	"encoding/json"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcTransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	"github.com/spf13/cast"

	"github.com/bze-alphateam/bze/docs"
	cointrunkmodule "github.com/bze-alphateam/bze/x/cointrunk"
	cointrunkmoduleclient "github.com/bze-alphateam/bze/x/cointrunk/client"
	cointrunkmodulekeeper "github.com/bze-alphateam/bze/x/cointrunk/keeper"
	cointrunkmoduletypes "github.com/bze-alphateam/bze/x/cointrunk/types"
	scavengemodule "github.com/bze-alphateam/bze/x/scavenge"
	scavengemodulekeeper "github.com/bze-alphateam/bze/x/scavenge/keeper"
	scavengemoduletypes "github.com/bze-alphateam/bze/x/scavenge/types"

	burnermodule "github.com/bze-alphateam/bze/x/burner"
	burnermoduleclient "github.com/bze-alphateam/bze/x/burner/client"
	burnermodulekeeper "github.com/bze-alphateam/bze/x/burner/keeper"
	burnermoduletypes "github.com/bze-alphateam/bze/x/burner/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

const (
	AccountAddressPrefix = "bze"
	Name                 = "bze"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		cointrunkmoduleclient.AcceptedDomainProposalHandler,
		cointrunkmoduleclient.PublisherProposalHandler,
		burnermoduleclient.BurnCoinsProposalHandler,
		rewardsclient.ActivateTradingRewardProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		ibcTransfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		scavengemodule.AppModuleBasic{},
		cointrunkmodule.AppModuleBasic{},
		burnermodule.AppModuleBasic{},
		tokenfactory.AppModuleBasic{},
		tradebin.AppModuleBasic{},
		epochs.AppModuleBasic{},
		rewards.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:         nil,
		distrtypes.ModuleName:              nil,
		minttypes.ModuleName:               {authtypes.Minter},
		stakingtypes.BondedPoolName:        {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:     {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:                {authtypes.Burner},
		ibctransfertypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		scavengemoduletypes.ModuleName:     nil,
		cointrunkmoduletypes.ModuleName:    nil,
		burnermoduletypes.ModuleName:       {authtypes.Burner},
		tokenfactorytypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		tradebintypes.ModuleName:           nil,
		epochstypes.ModuleName:             nil,
		rewardstypes.ModuleName:            {authtypes.Burner},
		burnermoduletypes.RaffleModuleName: {authtypes.Burner},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// a map of modules allowed to receive external funds
	// this list will be excluded from bank/distribution modules blocked addresses
	allowedModules = map[string]struct{}{
		burnermoduletypes.ModuleName: {},
		tradebintypes.ModuleName:     {},
	}
)

var (
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          client.TxConfig

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	AuthzKeeper           *authzkeeper.Keeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper        evidencekeeper.Keeper
	TransferKeeper        ibctransferkeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	ScavengeKeeper     scavengemodulekeeper.Keeper
	CointrunkKeeper    cointrunkmodulekeeper.Keeper
	BurnerKeeper       burnermodulekeeper.Keeper
	TokenFactoryKeeper tokenfactorykeeper.Keeper
	TradebinKeeper     tradebinkeeper.Keeper
	EpochsKeeper       epochskeeper.Keeper
	RewardsKeeper      rewardskeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm           *module.Manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

func (app *App) RegisterNodeService(context client.Context) {
	//TODO implement me
	panic("implement me")
}

// New returns a reference to an initialized Gaia.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		crisistypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		ibcexported.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		authzkeeper.StoreKey,
		consensusparamtypes.StoreKey,
		scavengemoduletypes.StoreKey,
		cointrunkmoduletypes.StoreKey,
		burnermoduletypes.StoreKey,
		tokenfactorytypes.StoreKey,
		tradebintypes.StoreKey,
		epochstypes.StoreKey,
		rewardstypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		txConfig:          txConfig,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[upgradetypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String())
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	// Sealing prevents other modules from creating scoped sub-keepers
	app.CapabilityKeeper.Seal()

	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		AccountAddressPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	authzKeeper := authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		app.AccountKeeper,
	)
	app.AuthzKeeper = &authzKeeper

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetBlockedAddresses(), authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], stakingKeeper,
		app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.AccountKeeper, app.BankKeeper,
		stakingKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, cdc, keys[slashingtypes.StoreKey], stakingKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// register the staking hooks
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibcexported.StoreKey], app.GetSubspace(ibcexported.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := ibcTransfer.NewAppModule(app.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	app.ScavengeKeeper = *scavengemodulekeeper.NewKeeper(
		appCodec,
		keys[scavengemoduletypes.StoreKey],
		keys[scavengemoduletypes.MemStoreKey],
		app.BankKeeper,
	)

	app.EpochsKeeper = *epochskeeper.NewKeeper(
		appCodec,
		keys[epochstypes.StoreKey],
		keys[epochstypes.MemStoreKey],
	)

	app.CointrunkKeeper = *cointrunkmodulekeeper.NewKeeper(
		appCodec,
		keys[cointrunkmoduletypes.StoreKey],
		keys[cointrunkmoduletypes.MemStoreKey],
		app.GetSubspace(cointrunkmoduletypes.ModuleName),
		app.BankKeeper,
		app.GovKeeper,
		app.AccountKeeper,
		app.DistrKeeper,
	)

	app.BurnerKeeper = *burnermodulekeeper.NewKeeper(
		appCodec,
		keys[burnermoduletypes.StoreKey],
		keys[burnermoduletypes.MemStoreKey],
		app.GetSubspace(burnermoduletypes.ModuleName),
		app.BankKeeper,
		app.AccountKeeper,
		app.EpochsKeeper,
	)

	app.TokenFactoryKeeper = *tokenfactorykeeper.NewKeeper(
		appCodec,
		keys[tokenfactorytypes.StoreKey],
		keys[tokenfactorytypes.MemStoreKey],
		app.GetSubspace(tokenfactorytypes.ModuleName),
		app.BankKeeper,
		app.AccountKeeper,
		app.DistrKeeper,
	)

	app.TradebinKeeper = *tradebinkeeper.NewKeeper(
		appCodec,
		keys[tradebintypes.StoreKey],
		keys[tradebintypes.MemStoreKey],
		app.GetSubspace(tradebintypes.ModuleName),
		app.BankKeeper,
		app.DistrKeeper,
	)

	app.RewardsKeeper = *rewardskeeper.NewKeeper(
		appCodec,
		keys[rewardstypes.StoreKey],
		keys[rewardstypes.MemStoreKey],
		app.GetSubspace(rewardstypes.ModuleName),
		app.BankKeeper,
		app.DistrKeeper,
		app.TradebinKeeper,
		app.EpochsKeeper,
	)

	app.TradebinKeeper.SetOnOrderFillHooks(
		[]tradebintypes.OnMarketOrderFill{
			app.RewardsKeeper.GetOnOrderFillHook(),
		},
	)

	app.EpochsKeeper.SetHooks(
		[]epochstypes.EpochHook{
			app.RewardsKeeper.GetDistributeAllStakingRewardsHook(),
			app.RewardsKeeper.GetUnlockPendingUnlockParticipantsHook(),
			app.RewardsKeeper.GetRemoveExpiredPendingTradingRewardsHook(),
			app.RewardsKeeper.GetTradingRewardsDistributionHook(),
			app.BurnerKeeper.GetBurnerRaffleCleanupHook(),
			app.BurnerKeeper.GetBurnerPeriodicBurnHook(),
		},
	)

	scavengeModule := scavengemodule.NewAppModule(appCodec, app.ScavengeKeeper)
	cointrunkModule := cointrunkmodule.NewAppModule(appCodec, app.CointrunkKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper)
	burnerModule := burnermodule.NewAppModule(appCodec, app.BurnerKeeper, app.AccountKeeper, app.BankKeeper)
	tokenfactoryModule := tokenfactory.NewAppModule(appCodec, app.TokenFactoryKeeper, app.AccountKeeper, app.BankKeeper)
	tradebinModule := tradebin.NewAppModule(appCodec, app.TradebinKeeper, app.BankKeeper)
	epochsModule := epochs.NewAppModule(appCodec, app.EpochsKeeper)
	rewardsModule := rewards.NewAppModule(appCodec, app.RewardsKeeper, app.AccountKeeper, app.BankKeeper)
	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// register the proposal types
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(cointrunkmoduletypes.RouterKey, cointrunkmodule.NewCointrunkProposalHandler(app.CointrunkKeeper)).
		AddRoute(burnermoduletypes.RouterKey, burnermodule.NewBurnerProposalHandler(app.BurnerKeeper)).
		AddRoute(rewardstypes.RouterKey, rewards.NewRewardsProposalHandler(app.RewardsKeeper))
	govKeeper.SetLegacyRouter(govRouter)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)
	transferIBCModule := ibcTransfer.NewIBCModule(app.TransferKeeper)
	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		authzmodule.NewAppModule(appCodec, *app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		scavengeModule,
		cointrunkModule,
		burnerModule,
		tokenfactoryModule,
		tradebinModule,
		epochsModule,
		rewardsModule,
		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		authtypes.ModuleName,
		genutiltypes.ModuleName,
		banktypes.ModuleName,
		capabilitytypes.ModuleName,
		stakingtypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		govtypes.ModuleName,
		paramstypes.ModuleName,
		crisistypes.ModuleName,
		slashingtypes.StoreKey,
		feegrant.ModuleName,
		ibcexported.ModuleName,
		upgradetypes.ModuleName,
		consensusparamtypes.ModuleName,
		evidencetypes.ModuleName,
		scavengemoduletypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		cointrunkmoduletypes.ModuleName,
		burnermoduletypes.ModuleName,
		tokenfactorytypes.ModuleName,
		tradebintypes.ModuleName,
		epochstypes.ModuleName,
		rewardstypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		authtypes.ModuleName,
		genutiltypes.ModuleName,
		banktypes.ModuleName,
		capabilitytypes.ModuleName,
		stakingtypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		govtypes.ModuleName,
		paramstypes.ModuleName,
		crisistypes.ModuleName,
		slashingtypes.StoreKey,
		feegrant.ModuleName,
		ibcexported.ModuleName,
		upgradetypes.ModuleName,
		consensusparamtypes.ModuleName,
		evidencetypes.ModuleName,
		scavengemoduletypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		cointrunkmoduletypes.ModuleName,
		burnermoduletypes.ModuleName,
		tokenfactorytypes.ModuleName,
		tradebintypes.ModuleName,
		epochstypes.ModuleName,
		rewardstypes.ModuleName,
	)

	genesisModuleOrder := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		scavengemoduletypes.ModuleName,
		feegrant.ModuleName,
		upgradetypes.ModuleName,
		consensusparamtypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		cointrunkmoduletypes.ModuleName,
		burnermoduletypes.ModuleName,
		tokenfactorytypes.ModuleName,
		tradebintypes.ModuleName,
		epochstypes.ModuleName,
		rewardstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(genesisModuleOrder...)
	app.mm.SetOrderExportGenesis(genesisModuleOrder...)

	app.mm.RegisterInvariants(app.CrisisKeeper)
	cfg := module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.configurator = cfg
	app.mm.RegisterServices(cfg)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	anteHandler, err := NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)
	app.setupUpgradeHandlers(cfg)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

func (app *App) setupUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		v711.UpgradeName,
		v711.CreateUpgradeHandler(cfg, app.mm),
	)
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// TxConfig returns App's TxConfig.
func (app *App) TxConfig() client.TxConfig {
	return app.txConfig
}

// Configurator get app configurator
func (app *App) Configurator() module.Configurator {
	return app.configurator
}

// ModuleManager returns the app ModuleManager
func (app *App) ModuleManager() *module.Manager {
	return app.mm
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// GetBlockedAddresses returns all addresses that should be blocked from direct sending
func (app *App) GetBlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		// from the list of modules exclude the allowed ones (if any)
		_, ok := allowedModules[acc]
		if ok {
			continue
		}

		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	if apiConfig.Swagger {
		apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
		apiSvr.Router.HandleFunc("/", openapi.Handler(Name, "/static/openapi.yml"))
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibcexported.ModuleName)
	paramsKeeper.Subspace(scavengemoduletypes.ModuleName)
	paramsKeeper.Subspace(cointrunkmoduletypes.ModuleName)
	paramsKeeper.Subspace(burnermoduletypes.ModuleName)
	paramsKeeper.Subspace(tokenfactorytypes.ModuleName)
	paramsKeeper.Subspace(tradebintypes.ModuleName)
	paramsKeeper.Subspace(rewardstypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}
