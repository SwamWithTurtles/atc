package migrations

import "github.com/BurntSushi/migration"

func AddPublicToPipelines(tx migration.LimitedTx) error {
	_, err := tx.Exec(`ALTER TABLE pipelines ADD COLUMN public boolean NOT NULL default false`)
	return err
}
