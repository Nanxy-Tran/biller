package services

import (
	"biller/models"
	"biller/repositories"
)

func CreateBill(bill *models.Bill, repo repositories.BillRepository) repositories.RepositoryResult {
	result := repo.Save(bill)
	return result
}

func GetBills(repo repositories.BillRepository) repositories.RepositoryResult {
	result := repo.GetBills()
	return result
}
