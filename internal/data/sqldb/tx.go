package sqldb

import (
	"opspillar/internal/data/repo"

	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TxManagerGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewTxManagerGorm(data *DataGorm, logger log.Logger) repo.TxManager {
	return &TxManagerGorm{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "transaction")),
	}
}

func (tm *TxManagerGorm) RunInTX(fn func(tx repo.TX) error) (err error) {
	tx := tm.data.DB.Begin()
	if tx.Error != nil {
		tm.log.Errorf("failed to begin transaction: %v", tx.Error)
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			tm.log.Errorf("panic occurred, rolling back transaction: %v", r)
			err = fmt.Errorf("panic in transaction: %v", r)
		}
	}()

	gtx := &TxGorm{tx: tx, log: tm.log}
	if err := fn(gtx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tm.log.Errorf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}

type TxGorm struct {
	tx  *gorm.DB
	log *log.Helper
}

func (gtx *TxGorm) GetDB() interface{} {
	return gtx.tx
}

func (gtx *TxGorm) Error(err error) bool {
	if err != nil {
		gtx.log.Errorf("transaction error: %v", err)
		return true
	}
	return false
}
