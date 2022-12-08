package gorms

import (
	"math"

	"gorm.io/gorm"
)

// Param 分页参数
type Param struct {
	Page    int
	Limit   int
	OrderBy []string
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

func GetPage(CurrentPage, PageSize int) (limit, offset int) {
	if CurrentPage < 1 {
		CurrentPage = 1
	}
	if PageSize == 0 {
		limit = 10
	} else {
		limit = PageSize
	}
	if CurrentPage == 1 {
		offset = 0
	} else {
		offset = (CurrentPage - 1) * limit
	}
	return
}

// Paging 分页
func Paging(db *gorm.DB, p *Param, result interface{}) (*Paginator, error) {
	var count int64
	countRecords(db, &count)
	if db.Error != nil {
		return nil, db.Error
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}
	var paginator Paginator
	limit, offset := GetPage(p.Page, p.Limit)
	err := db.Limit(limit).Offset(offset).Find(result).Error
	if err != nil {
		return nil, err
	}
	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page
	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))
	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}
	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator, nil
}

func countRecords(db *gorm.DB, count *int64) {
	db.Count(count)
}
