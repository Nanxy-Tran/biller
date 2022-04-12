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
	//err := r.DB.(bill).Error
	//if err != nil {
	//	log.Println(err)
	//	return RepositoryResult{Error: err}
	//}
	return RepositoryResult{Result: nil}
}

func (r *BillRepository) GetBills() RepositoryResult {
	var bills []models.Bill
	rows, err := r.DB.Query("SELECT * from bills")
	resultChan := make(chan models.Bill, 5)

	if err != nil {
		return RepositoryResult{Error: err}
	}

	go func(rows2 *sql.Rows) {
		defer close(resultChan)
		for rows2.Next() {
			var bill models.Bill
			if err := rows2.Scan(&bill.ID, &bill.Name, &bill.Amount, &bill.Description, &bill.CreatedAt); err != nil {
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
	//r.DB.First(&bill, id)
	return RepositoryResult{Result: bill}
}
