package utils

import "gorm.io/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func SetPage(totalPage *int64, totalCount int64, count int) {
	if totalCount != 0 {
		*totalPage = totalCount / int64(count)
		if totalCount%int64(count) > 0 {
			*totalPage++
		}
	}
}
