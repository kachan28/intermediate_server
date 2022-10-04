package repository

import (
	"fmt"
	"intermediate_server/internal/models"
	pb "intermediate_server/internal/models/pb"
	"time"
)

const (
	backupsTable = "backups"
)

var (
	saveQuery = fmt.Sprintf(`
		insert into %s values(NULL, ?, ?, ?, ?);
	`, backupsTable)
	getLastInsertIDQuery = fmt.Sprintf(`
		select id from %s order by id desc limit 1;
	`, backupsTable)
	deleteQuery = fmt.Sprintf(`
		delete from %s where id=?;
	`, backupsTable)
	listQuery = fmt.Sprintf(`
		select system_id, id, title, created_at from %s;
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

	_, err = stmt.Exec(backup.SystemId, backup.File.Title, backup.File.Content, time.Now())
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
	res, err := r.db.Exec(deleteQuery, backupID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("row has not been deleted, check query or args in server code")
	}

	return err
}

func (r *Repository) List() (map[models.SystemID][]models.Backup, error) {
	r.openDb()
	defer r.closeDb()

	rows, err := r.db.Query(listQuery)
	if err != nil {
		return nil, err
	}

	var systemID string
	var backup models.Backup
	backups := make(map[models.SystemID][]models.Backup, 0)

	for rows.Next() {
		err = rows.Scan(&systemID, &backup.ID, &backup.Title, &backup.CreatedAt)
		if err != nil {
			return nil, err
		}

		if len(backups[models.SystemID(systemID)]) == 0 {
			backups[models.SystemID(systemID)] = make([]models.Backup, 0)
		}

		backups[models.SystemID(systemID)] = append(
			backups[models.SystemID(systemID)],
			models.Backup{
				ID:        backup.ID,
				Title:     backup.Title,
				CreatedAt: backup.CreatedAt,
			},
		)
	}

	return backups, nil
}
