package db

import (
	"github.com/yagyansh5/simplebank/models"
	"gorm.io/gorm"
)

// CreateEntry creates a new entry in the database (deposit/withdrawal).
func CreateEntry(db *gorm.DB, entry models.Entry) (models.Entry, error) {
	if err := db.Create(&entry).Error; err != nil {
		return models.Entry{}, err
	}
	return entry, nil
}

// GetEntry fetches an entry by its ID.
func GetEntry(db *gorm.DB, id int64) (models.Entry, error) {
	var entry models.Entry
	if err := db.First(&entry, id).Error; err != nil {
		return models.Entry{}, err
	}
	return entry, nil
}

// UpdateEntry updates an entry's amount or other details.
func UpdateEntry(db *gorm.DB, id int64, entry models.Entry) (models.Entry, error) {
	var existingEntry models.Entry
	if err := db.First(&existingEntry, id).Error; err != nil {
		return models.Entry{}, err
	}

	// Update the amount
	existingEntry.Amount = entry.Amount

	if err := db.Save(&existingEntry).Error; err != nil {
		return models.Entry{}, err
	}
	return existingEntry, nil
}

// DeleteEntry deletes an entry by its ID.
func DeleteEntry(db *gorm.DB, id int64) error {
	var entry models.Entry
	if err := db.First(&entry, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&entry).Error; err != nil {
		return err
	}
	return nil
}
