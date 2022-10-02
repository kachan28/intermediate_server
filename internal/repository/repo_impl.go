package repository

import (
	"context"
	"fmt"
	pb "intermediate_server/internal/models/pb"
)

const (
	backupsTable = "backups"

	saveQuery = `
		insert into %s values(NULL, ?, ?, ?)
	`
	getLastInsertIDQuery = `select id from backups order by id desc limit 1;`
	deleteQuery          = ``
)

func (r *Repository) Create(ctx context.Context, cancel context.CancelFunc, backup *pb.BackupCreate) error {
	defer cancel()

	r.openDb()
	defer r.closeDb()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(fmt.Sprintf(saveQuery, backupsTable))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, backup.Id, backup.File.Title, backup.File.Content)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetLastInsertedRowID(ctx context.Context) (int64, error) {
	r.openDb()
	defer r.closeDb()

	var lastInsertID int64
	err := r.db.QueryRowContext(ctx, getLastInsertIDQuery).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (r *Repository) Delete(ctx context.Context) error {
	r.openDb()
	defer r.closeDb()

}
