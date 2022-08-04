package utils

import "gorm.io/gorm"

const pageSize = 15

//TODO Better pagination https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f

func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
