package repositories

import (
	"biller/models"
	"database/sql"
	"fmt"
	"math"
)

type BillRepository struct {
	DB *sql.DB
}

func InitBillRepository(db *sql.DB) *BillRepository {
	return &BillRepository{
		DB: db,
	}
}

type BillsOptions struct {
	Limit       int `form:"limit"`
	CurrentPage int `form:"current_page"`
}

type BillsResponse struct {
	CurrentPage int           `json:"current_page,omitempty"`
	TotalPage   int           `json:"total_page,omitempty"`
	Bills       []models.Bill `json:"bills,omitempty"`
	Error       error         `json:"error,omitempty"`
}

//TODO: make prepared statement for it
// bill -> categories

func (r *BillRepository) GetBills(options *BillsOptions) BillsResponse {
	var bills []models.Bill
	var totalBills int

	rows, err := r.DB.Query("SELECT SQL_CALC_FOUND_ROWS * from bills LIMIT ? OFFSET ?",
		options.Limit, (options.CurrentPage-1)*options.Limit,
	)

	resultChan := make(chan models.Bill, 10)
	defer rows.Close()

	if err != nil {
		return BillsResponse{Error: err}
	}

	go func(currentRow *sql.Rows) {
		defer close(resultChan)
		for currentRow.Next() {
			var bill models.Bill
			if err := currentRow.Scan(
				&bill.ID,
				&bill.Name,
				&bill.Amount,
				&bill.Description,
				&bill.CreatedAt,
			); err != nil {
				fmt.Println(err)
				return
			}
			resultChan <- bill
		}
	}(rows)

	for bill := range resultChan {
		bills = append(bills, bill)
	}

	row := r.DB.QueryRow("SELECT FOUND_ROWS()")
	err = row.Scan(&totalBills)

	return BillsResponse{
		TotalPage:   int(math.Ceil(float64(totalBills / options.Limit))),
		CurrentPage: options.CurrentPage,
		Bills:       bills,
	}
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
