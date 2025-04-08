package db

import (
	"github.com/yagyansh5/simplebank/models"
	"gorm.io/gorm"
)

// CreateAccount creates a new account in the database.
func CreateAccount(db *gorm.DB, account models.Account) (models.Account, error) {
	if err := db.Create(&account).Error; err != nil {
		return models.Account{}, err
	}
	return account, nil
}

// GetAccount fetches an account by its ID.
func GetAccount(db *gorm.DB, id int64) (models.Account, error) {
	var account models.Account
	if err := db.First(&account, id).Error; err != nil {
		return models.Account{}, err
	}
	return account, nil
}

// UpdateAccount updates the details of an existing account.
func UpdateAccount(db *gorm.DB, id int64, account models.Account) (models.Account, error) {
	var existingAccount models.Account
	if err := db.First(&existingAccount, id).Error; err != nil {
		return models.Account{}, err
	}

	// Update fields
	existingAccount.Owner = account.Owner
	existingAccount.Balance = account.Balance
	existingAccount.Currency = account.Currency
	existingAccount.CountryCode = account.CountryCode

	if err := db.Save(&existingAccount).Error; err != nil {
		return models.Account{}, err
	}
	return existingAccount, nil
}

// DeleteAccount deletes an account by ID.
func DeleteAccount(db *gorm.DB, id int64) error {
	var account models.Account
	if err := db.First(&account, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&account).Error; err != nil {
		return err
	}
	return nil
}
