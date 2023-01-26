//go:generate zebrapack
package store

import (
	"context"
	"errors"
	"time"

	"github.com/TAAL-GmbH/arc/metamorph/metamorph_api"
	"github.com/libsv/go-p2p/chaincfg/chainhash"
)

type StoreData struct {
	StoredAt      time.Time            `zid:"0"`
	AnnouncedAt   time.Time            `zid:"1"`
	MinedAt       time.Time            `zid:"2"`
	Hash          *chainhash.Hash      `badgerhold:"key" zid:"3"`
	Status        metamorph_api.Status `zid:"4"`
	BlockHeight   uint64               `zid:"5"`
	BlockHash     *chainhash.Hash      `zid:"6"`
	ApiKeyId      int64                `zid:"7"`
	StandardFeeId int64                `zid:"8"`
	DataFeeId     int64                `zid:"9"`
	SourceIp      string               `zid:"10"`
	CallbackUrl   string               `zid:"11"`
	CallbackToken string               `zid:"12"`
	MerkleProof   bool                 `zid:"13"`
	RawTx         []byte               `zid:"14"`
	RejectReason  string               `zid:"15"`
}

var ErrNotFound = errors.New("txid could not be found")

type MetamorphStore interface {
	Get(ctx context.Context, key []byte) (*StoreData, error)
	Set(ctx context.Context, key []byte, value *StoreData) error
	Del(ctx context.Context, key []byte) error

	GetUnmined(_ context.Context, callback func(s *StoreData)) error
	UpdateStatus(ctx context.Context, hash *chainhash.Hash, status metamorph_api.Status, rejectReason string) error
	UpdateMined(ctx context.Context, hash *chainhash.Hash, blockHash *chainhash.Hash, blockHeight uint64) error
	Close(ctx context.Context) error
	GetBlockProcessed(ctx context.Context, blockHash *chainhash.Hash) (*time.Time, error)
	SetBlockProcessed(ctx context.Context, blockHash *chainhash.Hash) error
}
