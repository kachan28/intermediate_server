package repository

import (
	"fmt"
	pb "intermediate_server/internal/models/pb"
)

const (
	backupsTable = "backups"
)

var (
	saveQuery = fmt.Sprintf(`
		insert into %s values(NULL, ?, ?, ?);
	`, backupsTable)
	getLastInsertIDQuery = fmt.Sprintf(`
		select id from %s order by id desc limit 1;
	`, backupsTable)
	deleteQuery = fmt.Sprintf(`
		delete from %s where id=?;
	`, backupsTable)
)

func (r *Repository) Create(backup *pb.BackupCreate) error {
	r.openDb()
	defer r.closeDb()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(saveQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(backup.Id, backup.File.Title, backup.File.Content)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetLastInsertedRowID() (int64, error) {
	r.openDb()
	defer r.closeDb()

	var lastInsertID int64
	err := r.db.QueryRow(getLastInsertIDQuery).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (r *Repository) Delete(backupID int64) error {
	r.openDb()
	defer r.closeDb()
	_, err := r.db.Query(deleteQuery)
	return err
}
