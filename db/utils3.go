package db

import (
	"github.com/yagyansh5/simplebank/models"
	"gorm.io/gorm"
)

// CreateTransfer creates a new transfer record in the database.
func CreateTransfer(db *gorm.DB, transfer models.Transfer) (models.Transfer, error) {
	if err := db.Create(&transfer).Error; err != nil {
		return models.Transfer{}, err
	}
	return transfer, nil
}

// GetTransfer fetches a transfer by its ID.
func GetTransfer(db *gorm.DB, id int64) (models.Transfer, error) {
	var transfer models.Transfer
	if err := db.First(&transfer, id).Error; err != nil {
		return models.Transfer{}, err
	}
	return transfer, nil
}

// UpdateTransfer updates the details of an existing transfer.
func UpdateTransfer(db *gorm.DB, id int64, transfer models.Transfer) (models.Transfer, error) {
	var existingTransfer models.Transfer
	if err := db.First(&existingTransfer, id).Error; err != nil {
		return models.Transfer{}, err
	}

	// Update transfer details
	existingTransfer.FromAccountID = transfer.FromAccountID
	existingTransfer.ToAccountID = transfer.ToAccountID
	existingTransfer.Amount = transfer.Amount

	if err := db.Save(&existingTransfer).Error; err != nil {
		return models.Transfer{}, err
	}
	return existingTransfer, nil
}

// DeleteTransfer deletes a transfer record by ID.
func DeleteTransfer(db *gorm.DB, id int64) error {
	var transfer models.Transfer
	if err := db.First(&transfer, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&transfer).Error; err != nil {
		return err
	}
	return nil
}
