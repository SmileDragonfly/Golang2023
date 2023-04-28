package storage

import (
	"context"
	"gorm.io/gorm"
	"gormtest/GormTest/model"
	"log"
)

type db_storage struct {
	db *gorm.DB
}

func NewDBStorage(db *gorm.DB) *db_storage {
	return &db_storage{db}
}

func (p db_storage) CreateWallet(ctx context.Context, data *model.Wallet) error {
	if err := p.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) CreateBank(ctx context.Context, data *model.Bank) error {
	if err := p.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) CreateLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	if err := p.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) DeleteWallet(ctx context.Context, condition map[string]interface{}) error {
	if err := p.db.Table(model.Wallet{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) DeleteBank(ctx context.Context, condition map[string]interface{}) error {
	if err := p.db.Table(model.Bank{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) DeleteLinkInfo(ctx context.Context, condition map[string]interface{}) error {
	if err := p.db.Table(model.LinkInfo{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) ReadWallet(ctx context.Context, condition map[string]interface{}) ([]model.Wallet, error) {
	var rows []model.Wallet
	if err := p.db.Find(rows, condition).Error; err != nil {
		return nil, err
	}
	log.Printf("Find %d rows", len(rows))
	return rows, nil
}

func (p db_storage) ReadBank(ctx context.Context, condition map[string]interface{}) ([]model.Bank, error) {
	var rows []model.Bank
	if err := p.db.Find(rows, condition).Error; err != nil {
		return nil, err
	}
	log.Printf("Find %d rows", len(rows))
	return rows, nil
}

func (p db_storage) ReadLinkInfo(ctx context.Context, condition map[string]interface{}) ([]model.LinkInfo, error) {
	var rows []model.LinkInfo
	if err := p.db.Find(rows, condition).Error; err != nil {
		return nil, err
	}
	log.Printf("Find %d rows", len(rows))
	return rows, nil
}

func (p db_storage) UpdateWallet(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Wallet) error {
	if err := p.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) UpdateBank(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Bank) error {
	if err := p.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (p db_storage) UpdateLinkInfo(ctx context.Context, condition map[string]interface{}, dataUpdate *model.LinkInfo) error {
	if err := p.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
