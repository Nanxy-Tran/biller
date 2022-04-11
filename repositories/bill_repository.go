package repositories

import (
	"biller/models"
	"gorm.io/gorm"
	"log"
)

type BillRepository struct {
	DB *gorm.DB
}

func InitBillRepository(db *gorm.DB) *BillRepository {
	return &BillRepository{
		DB: db,
	}
}

func (r *BillRepository) Save(bill *models.Bill) RepositoryResult {
	err := r.DB.Create(bill).Error
	if err != nil {
		log.Println(err)
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: bill}
}

func (r *BillRepository) GetBills() RepositoryResult {
	var bills []models.Bill
	r.DB.Find(&bills)
	return RepositoryResult{Result: bills}
}
