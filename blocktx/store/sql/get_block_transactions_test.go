package sql

import (
	"context"
	"testing"

	"github.com/bitcoin-sv/arc/blocktx/blocktx_api"
	"github.com/bitcoin-sv/arc/blocktx/store"
	. "github.com/bitcoin-sv/arc/database_testing"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetBlockTransactionsSuite struct {
	DatabaseTestSuite
}

func (s *GetBlockByHeightTestSuite) TestGetBlockTransactions() {
	block := GetTestBlock()
	tx := GetTestTransaction()
	s.InsertBlock(block)

	s.InsertTransaction(tx)

	s.InsertBlockTransactionMap(&store.BlockTransactionMap{
		BlockID:       block.ID,
		TransactionID: tx.ID,
		Pos:           2,
	})

	st, err := NewPostgresStore(DefaultParams)

	require.NoError(s.T(), err)

	txs, err := st.GetBlockTransactions(context.Background(), &blocktx_api.Block{Hash: []byte(block.Hash)})

	require.NoError(s.T(), err)
	assert.Equal(s.T(), tx.Hash, string(txs.Transactions[0].Hash))
}

func TestGetBlockTransactions(t *testing.T) {
	s := new(GetBlockTransactionsSuite)
	suite.Run(t, s)
}
