package main

import "database/sql"

type SQLiteDriver struct {
	DB *sql.DB
}

func (d SQLiteDriver) Insert(label string, isComplete bool) (interface{}, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare("insert into todos(label, is_complete) values(?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rowsAffected, err := stmt.Exec(label, isComplete)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := rowsAffected.LastInsertId()
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return lastInsertID, nil
}
