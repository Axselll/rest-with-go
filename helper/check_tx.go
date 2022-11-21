package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		CheckError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		CheckError(errCommit)
	}
}
