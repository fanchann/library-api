package utils

import (
	"database/sql"
)

func StartTransaction(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	LogErrorWithPanic(err)

	return tx
}

func TransactionsCommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		LogErrorWithPanic(errRollback)
	}
	errCommit := tx.Commit()
	LogErrorWithPanic(errCommit)
}
