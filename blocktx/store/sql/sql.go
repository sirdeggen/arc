package sql

import (
	"database/sql"
	"fmt"

	"github.com/TAAL-GmbH/arc/blocktx/store"
	_ "github.com/lib/pq"
	"github.com/ordishs/gocore"
	_ "modernc.org/sqlite"
)

type SQL struct {
	db *sql.DB
}

func NewSQLStore(engine string) (store.Interface, error) {
	var db *sql.DB
	var err error

	switch engine {
	case "postgres":
		dbHost, _ := gocore.Config().Get("dbHost", "localhost")
		dbPort, _ := gocore.Config().GetInt("dbPort", 5432)
		dbName, _ := gocore.Config().Get("dbName", "blocktx")
		dbUser, _ := gocore.Config().Get("dbUser", "blocktx")
		dbPassword, _ := gocore.Config().Get("dbPassword", "blocktx")

		dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d", dbUser, dbPassword, dbName, dbHost, dbPort)

		db, err = sql.Open(engine, dbInfo)
		if err != nil {
			return nil, fmt.Errorf("failed to open postgres DB: %+v", err)
		}

		if err := createPostgresSchema(db); err != nil {
			return nil, fmt.Errorf("failed to create postgres schema: %+v", err)
		}

	case "sqlite":
		db, err = sql.Open("sqlite", "block-tx.db")
		if err != nil {
			return nil, fmt.Errorf("failed to open sqlite DB: %+v", err)
		}

		if err := createSqliteSchema(db); err != nil {
			return nil, fmt.Errorf("failed to create sqlite schema: %+v", err)
		}

	default:
		return nil, fmt.Errorf("unknown database engine: %s", engine)
	}

	return &SQL{
		db: db,
	}, nil
}

func createPostgresSchema(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE blocks (
		 id           BIGSERIAL PRIMARY KEY
	  ,hash         BYTEA NOT NULL
	  ,header       BYTEA NOT NULL
	  ,height       BIGINT NOT NULL
	  ,processedyn  BOOLEAN NOT NULL DEFAULT FALSE
	  ,orphanedyn   BOOLEAN NOT NULL DEFAULT FALSE
		);
	`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks table - [%+v]\n", err)
	}

	if _, err := db.Exec(`CREATE UNIQUE INDEX ON blocks (hash);`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks table - [%+v]\n", err)
	}

	if _, err := db.Exec(`CREATE UNIQUE INDEX PARTIAL ON blocks(height) WHERE orphanedyn = FALSE;`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks table - [%+v]\n", err)
	}

	if _, err := db.Exec(`
		CREATE TABLE block_transactions (
		 blockid      BIGINT NOT NULL REFERENCES blocks(id)
	  ,txhash       BYTEA NOT NULL
	  ,PRIMARY KEY (blockid, txhash)
		);
	`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks table - [%+v]\n", err)
	}

	return nil
}

func createSqliteSchema(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS blocks (
		 id           INTEGER PRIMARY KEY AUTOINCREMENT,
		 hash         BLOB NOT NULL
	  ,header       BLOB NOT NULL
	  ,height       BIGINT NOT NULL
	  ,processedyn  BOOLEAN NOT NULL DEFAULT FALSE
	  ,orphanedyn   BOOLEAN NOT NULL DEFAULT FALSE
	 	);
	`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks table - [%+v]\n", err)
	}

	if _, err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS unique_block ON blocks (hash);`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks hash index - [%+v]\n", err)
	}

	if _, err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS orphan_check ON blocks(height) WHERE orphanedyn = FALSE;`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create blocks height index - [%+v]\n", err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS block_transactions (
		 blockid      INTEGER NOT NULL REFERENCES blocks(id)
	  ,txhash       BLOB NOT NULL
	  ,PRIMARY KEY (blockid, txhash)
	  );
	`); err != nil {
		db.Close()
		return fmt.Errorf("Could not create block_transactions table - [%+v]\n", err)
	}

	return nil
}
