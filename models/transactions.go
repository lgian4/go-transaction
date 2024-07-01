package models

import (
	"errors"
	"time"

	"finance/db"
)

type Transaction struct {
	ID             int64
	DateTime       time.Time `binding:"required"`
	Description    string    `binding:"required"`
	Debit          float64   `binding:"required"`
	Credit         float64   `binding:"required"`
	CurrentBalance float64   `binding:"required"`
}

func (e *Transaction) Save() error {
	query := `
	insert into transactions (date_time, description, debit, credit, current_balance)
	values (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.DateTime, e.Description, e.Debit, e.Credit, e.CurrentBalance)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func GetAll() ([]Transaction, error) {
	query := `
	select id, date_time, description, debit, credit, current_balance from transactions
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Transaction
	for rows.Next() {
		var data Transaction
		err := rows.Scan(&data.ID, &data.DateTime, &data.Description, &data.Debit, &data.Credit, &data.CurrentBalance)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}

	return list, nil
}

func GetOne(id int64) (*Transaction, error) {
	query := `
	select id, date_time, description, debit, credit, current_balance 
	from transactions
	where id = ?
	`
	row := db.DB.QueryRow(query, id)
	if row == nil {
		return nil, errors.New("data not found")
	}
	var data Transaction
	err := row.Scan(&data.ID, &data.DateTime, &data.Description, &data.Debit, &data.Credit, &data.CurrentBalance)

	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (e *Transaction) Update() error {
	query := `
	update transactions  
	set	 date_time = ?, description = ?, debit = ?, credit = ?, current_balance = ?
	where id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.DateTime, e.Description, e.Debit, e.Credit, e.CurrentBalance, e.ID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (e *Transaction) Delete() error {
	query := `
	delete from transactions  
	where id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
