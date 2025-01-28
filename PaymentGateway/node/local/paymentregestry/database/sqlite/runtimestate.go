package sqlite

import (
	log "paidpiper.com/payment-gateway/log"

	"paidpiper.com/payment-gateway/node/local/paymentregestry/database/entity"
)

func (prdb *liteDb) createTableRuntimeState() error {
	return prdb.exec(`
	CREATE TABLE IF NOT EXISTS RuntimeState (
		TorRoutingMode         	INTEGER NOT NULL
	);
	INSERT INTO RuntimeState VALUES (1);
`)
}

func (prdb *liteDb) UpdateRuntimeState(item *entity.RuntimeState) error {
	tx, err := prdb.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}
	stmt, err := tx.Prepare(`UPDATE RuntimeState 
    SET TorRoutingMode=?
	;
`)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		item.TorRoutingMode,
	)
	if err != nil {
		return err
	}
	return tx.Commit()

}

func (prdb *liteDb) SelectRuntimeState() (*entity.RuntimeState, error) {
	query := `
		SELECT TorRoutingMode
		FROM RuntimeState;
	`

	res, err := prdb.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var item *entity.RuntimeState

	for res.Next() {
		var torRoutingMode int

		err := res.Scan(
			&torRoutingMode,
		)
		if err != nil {
			return nil, err
		}
		item = &entity.RuntimeState{
			TorRoutingMode: torRoutingMode,
		}
	}
	return item, nil
}
