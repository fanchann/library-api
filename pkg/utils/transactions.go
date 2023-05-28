package utils

import (
	"database/sql"
)

func TransactionsCommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		LogErrorWithPanic(errRollback)
	}
	errCommit := tx.Commit()
	LogErrorWithPanic(errCommit)
}
