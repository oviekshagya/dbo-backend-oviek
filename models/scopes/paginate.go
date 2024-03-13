package scopes

import (
	grm "gorm.io/gorm"
)

func Paginate(pageSize int, page int) func(db *grm.DB) *grm.DB {
	return func(db *grm.DB) *grm.DB {
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
