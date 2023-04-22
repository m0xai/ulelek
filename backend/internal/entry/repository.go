// repository.go

package entry

import (
	"database/sql"
)

type Repository interface {
	GetByID(id int) (*Entry, error)
	GetAll() ([]*Entry, error)
	Create(*CreateEntryRequest) (*Entry, error)
	Update(*UpdateEntryRequest) (*Entry, error)
	Delete(id int) error
}

type EntryRepository struct {
	db *sql.DB
}

func NewEntryRepo(db *sql.DB) *EntryRepository {
	return &EntryRepository{
		db: db,
	}
}

func (r *EntryRepository) GetByID(id int) (*Entry, error) {
	entry := &Entry{}
	err := r.db.QueryRow("SELECT id, content, is_deleted, created_at, updated_at FROM entry WHERE id=$1", id).
		Scan(&entry.ID, &entry.Content, &entry.IsDeleted, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *EntryRepository) GetAll() ([]*Entry, error) {
	rows, err := r.db.Query("SELECT id, content, is_deleted, created_at, updated_at FROM entry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []*Entry{}
	for rows.Next() {
		entry := &Entry{}
		err := rows.Scan(&entry.ID, &entry.Content, &entry.IsDeleted, &entry.CreatedAt, &entry.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *EntryRepository) Create(entry *CreateEntryRequest) (*Entry, error) {
	_, err := r.db.Exec("INSERT INTO entry (content, is_deleted) VALUES ($1, $2)",
		entry.Content, entry.IsDeleted)
	if err != nil {
		return nil, err
	}

	// SELECT the inserted row
	var newEntry Entry
	err = r.db.QueryRow("SELECT id, content, is_deleted, created_at, updated_at FROM entry WHERE id = (SELECT max(id) FROM entry)").Scan(&newEntry.ID, &newEntry.Content, &newEntry.IsDeleted, &newEntry.CreatedAt, &newEntry.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newEntry, nil
}

func (r *EntryRepository) Update(entry_req *UpdateEntryRequest) (*Entry, error) {
	_, err := r.db.Exec("UPDATE entry SET content=$1, is_deleted=$2, WHERE id=$3",
		entry_req.Content, entry_req.IsDeleted, entry_req.ID)
	if err != nil {
		return nil, err
	}

	// SELECT the inserted row
	var newEntry Entry
	err = r.db.QueryRow("SELECT id, content, is_deleted, created_at, updated_at FROM entry WHERE id = (SELECT max(id) FROM entry)").Scan(&newEntry.ID, &newEntry.Content, &newEntry.IsDeleted, &newEntry.CreatedAt, &newEntry.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newEntry, nil
}

func (r *EntryRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM entry WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
