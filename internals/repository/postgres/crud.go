package postgres

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"
	"gorm.io/gorm"
)

func (r *PostgresRepository) Create(ent interface{}) error {
	return r.db.Connection.Create(ent).Error
}

func (r *PostgresRepository) Find(filters map[string]interface{}, ent interface{}) error {
	return r.db.Connection.Where(filters).Find(ent).Error
}

func (r *PostgresRepository) First(filters map[string]interface{}, ent interface{}) error {
	return r.db.Connection.Where(filters).First(ent).Error
}

func (r *PostgresRepository) Last(filters map[string]interface{}, ent interface{}) error {
	return r.db.Connection.Where(filters).Last(ent).Error
}

func (r *PostgresRepository) List(tableName string, offset, limit int64, filters, order, ent interface{}) (*utils.Pagination, error) {
	var total int64

	if err := r.db.Connection.Where(filters).Order(order).Offset(int(offset)).Limit(int(limit)).Find(ent).Error; err != nil {
		return nil, err
	}

	if err := r.db.Connection.Table(tableName).Where(filters).Order(order).Count(&total).Error; err != nil {
		return nil, err
	}

	return utils.FormatPagination(ent, limit, total), nil
}

func (r *PostgresRepository) Raw(ent interface{}, sql string, value ...interface{}) error {
	return r.db.Connection.Raw(sql, value...).Take(ent).Error
}

func (r *PostgresRepository) Update(filters map[string]interface{}, ent interface{}) error {
	result := r.db.Connection.Where(filters).Updates(ent)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *PostgresRepository) Begin() *gorm.DB {
	return r.db.Connection.Begin()
}

func (r *PostgresRepository) Commit(db *gorm.DB) *gorm.DB {
	return db.Commit()
}

func (r *PostgresRepository) Rollback(db *gorm.DB) *gorm.DB {
	return db.Rollback()
}

func (r *PostgresRepository) RollbackTo(db *gorm.DB, input string) *gorm.DB {
	return db.RollbackTo(input)
}

func (r *PostgresRepository) SavePoint(db *gorm.DB, input string) *gorm.DB {
	return db.SavePoint(input)
}

func (r *PostgresRepository) CreateTransaction(db *gorm.DB, ent interface{}) error {
	return db.Create(ent).Error
}

func (r *PostgresRepository) UpdateTransaction(db *gorm.DB, filters map[string]interface{}, ent interface{}) error {
	return db.Where(filters).Updates(ent).Error
}

func (r *PostgresRepository) IsErrorRecordNotFound(err error) bool {
	return r.db.IsErrorRecordNotFound(err)
}

func (r *PostgresRepository) Upsert(tableName string, filters map[string]interface{}, ent interface{}) error {
	target := make(map[string]interface{})

	if err := r.db.Connection.Table(tableName).Where(filters).Take(&target).Error; err != nil {
		if !r.db.IsErrorRecordNotFound(err) {
			return err
		}
	}

	if len(target) != 0 {
		return r.db.Connection.Where(filters).Updates(ent).Error
	}

	return r.db.Connection.Create(ent).Error
}
