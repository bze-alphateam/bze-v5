package simapp

import (
	"github.com/bze-alphateam/bze/app/params"
	bzecmd "github.com/bze-alphateam/bze/cmd/bzed/cmd"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/bze-alphateam/bze/app"
)

// New creates application instance with in-memory database and disabled logging.
func New(dir string) bzecmd.App {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	encoding := params.MakeEncodingConfig()

	a := app.New(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
		simapp.EmptyAppOptions{})
	// InitChain updates deliverState which is required when app.NewContext is called
	a.InitChain(abci.RequestInitChain{
		ConsensusParams: defaultConsensusParams,
		AppStateBytes:   []byte("{}"),
	})
	return a
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}
