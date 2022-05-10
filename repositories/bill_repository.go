package repositories

import (
	"biller/models"
	"errors"
	"gorm.io/gorm"
	"math"
)

type BillRepository struct {
	DB *gorm.DB
}

func InitBillRepository(db *gorm.DB) *BillRepository {
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

func (r *BillRepository) GetBills(options *BillsOptions) BillsResponse {
	var bills []models.Bill

	results := r.DB.Limit(options.Limit).Offset(options.CurrentPage - 1).Find(&bills)

	if errors.Is(results.Error, gorm.ErrRecordNotFound) {
		return BillsResponse{Error: results.Error}
	}

	return BillsResponse{
		TotalPage:   int(math.Ceil(float64(int(results.RowsAffected) / options.Limit))),
		CurrentPage: options.CurrentPage,
		Bills:       bills,
	}
}

func (r *BillRepository) GetBill(id string) RepositoryResult {
	var bill models.Bill
	result := r.DB.First(&bill, id)

	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}

	return RepositoryResult{Result: bill}
}

func (r *BillRepository) Save(bill *models.Bill) RepositoryResult {
	result := r.DB.Create(&bill)
	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}
	return RepositoryResult{Result: bill.ID}

}
