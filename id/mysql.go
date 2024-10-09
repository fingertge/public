package id

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlRenew struct {
	table string
	db    *gorm.DB
}

func NewMysqlRenew(db *gorm.DB, table string) *MysqlRenew {
	return &MysqlRenew{table: table, db: db}
}

func (m *MysqlRenew) Prepare() error {
	return m.db.Table(m.table).AutoMigrate(&IdGen{})
}

func gormRollBack(tx *gorm.DB) error {
	if err := tx.Rollback().Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}
	return nil
}

func (m *MysqlRenew) Renew(ctx context.Context, domain string, quantum, offset uint64) (uint64, error) {
	curr, err := m.renew(ctx, m.table, domain, quantum)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = m.db.Table(m.table).Create(&IdGen{Domain: domain, Value: offset}).Error; err != nil {
			return 0, err
		}
		return m.renew(ctx, m.table, domain, quantum)
	}
	return curr, err
}

func (m *MysqlRenew) renew(ctx context.Context, table, domain string, quantum uint64) (id uint64, err error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	var idG IdGen
	if err = tx.Table(table).Clauses(clause.Locking{Strength: "UPDATE"}).Where("domain", domain).First(&idG).Error; err != nil {
		panicOnError(gormRollBack(tx))
		return 0, err
	}
	id = idG.Value
	idG.Value += quantum
	if err = tx.Save(&idG).Error; err != nil {
		panicOnError(gormRollBack(tx))
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		panicOnError(gormRollBack(tx))
		return 0, err
	}
	return id, nil
}
