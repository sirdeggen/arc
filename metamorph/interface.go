package metamorph

import (
	"time"

	"github.com/bitcoin-sv/arc/metamorph/metamorph_api"
	"github.com/libsv/go-p2p/chaincfg/chainhash"
)

type ProcessorI interface {
	LoadUnmined()
	Set(req *ProcessorRequest) error
	ProcessTransaction(req *ProcessorRequest)
	SendStatusForTransaction(hash *chainhash.Hash, status metamorph_api.Status, id string, err error) (bool, error)
	SendStatusMinedForTransaction(hash *chainhash.Hash, blockHash *chainhash.Hash, blockHeight uint64) (bool, error)
	GetStats(debugItems bool) *ProcessorStats
	GetPeers() ([]string, []string)
	Shutdown()
}

type PeerTxMessage struct {
	Start  time.Time
	Hash   *chainhash.Hash
	Status metamorph_api.Status
	Peer   string
	Err    error
}
