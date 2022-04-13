package repositories

import (
	"biller/models"
	"database/sql"
	"fmt"
)

type BillRepository struct {
	DB *sql.DB
}

func InitBillRepository(db *sql.DB) *BillRepository {
	return &BillRepository{
		DB: db,
	}
}

func (r *BillRepository) Save(bill *models.Bill) RepositoryResult {
	result, err := r.DB.Exec("INSERT INTO bills (name, amount, description) VALUE (?, ?, ?)", bill.Name, bill.Amount, bill.Description)
	if err != nil {
		return RepositoryResult{Error: err}
	}
	_, err = result.LastInsertId()

	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: result}

}

func (r *BillRepository) GetBills() RepositoryResult {
	var bills []models.Bill
	rows, err := r.DB.Query("SELECT * from bills")
	resultChan := make(chan models.Bill, 5)
	defer rows.Close()

	if err != nil {
		return RepositoryResult{Error: err}
	}

	go func(currentRow *sql.Rows) {
		defer close(resultChan)
		for currentRow.Next() {
			var bill models.Bill
			if err := currentRow.Scan(&bill.ID, &bill.Name, &bill.Amount, &bill.Description, &bill.CreatedAt); err != nil {
				fmt.Println(err)
				return
			}
			resultChan <- bill
		}
	}(rows)

	for bill := range resultChan {
		bills = append(bills, bill)
	}

	return RepositoryResult{Result: bills}
}

func (r *BillRepository) GetBill(id string) RepositoryResult {
	var bill models.Bill
	row := r.DB.QueryRow("SELECT * from bills where id=?", id)

	err := row.Scan(&bill.ID, &bill.Name, &bill.Amount, &bill.Description, &bill.CreatedAt)

	if err == sql.ErrNoRows {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: bill}
}
