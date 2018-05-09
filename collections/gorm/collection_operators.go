package gorm

import (
	"strings"

	"github.com/infobloxopen/atlas-app-toolkit/collections"
	"github.com/jinzhu/gorm"
)

// ApplyFiltering applies filtering operator f to gorm instance db.
func ApplyFiltering(db *gorm.DB, f *collections.Filtering) (*gorm.DB, error) {
	str, args, err := FilteringToGorm(f)
	if err != nil {
		return nil, err
	}
	if str != "" {
		return db.Where(str, args...), nil
	}
	return db, nil
}

// ApplySorting applies sorting operator s to gorm instance db.
func ApplySorting(db *gorm.DB, s *collections.Sorting) *gorm.DB {
	var crs []string
	for _, cr := range s.GetCriterias() {
		if cr.IsDesc() {
			crs = append(crs, cr.GetTag()+" desc")
		} else {
			crs = append(crs, cr.GetTag())
		}
	}
	if len(crs) > 0 {
		return db.Order(strings.Join(crs, ","))
	}
	return db
}

// ApplyPagination applies pagination operator p to gorm instance db.
func ApplyPagination(db *gorm.DB, p *collections.Pagination) *gorm.DB {
	return db.Offset(p.GetOffset()).Limit(p.DefaultLimit())
}

// ApplyFieldSelection applies field selection operator fs to gorm instance db.
func ApplyFieldSelection(db *gorm.DB, fs *collections.FieldSelection) *gorm.DB {
	var fields []string
	for _, f := range fs.GetFields() {
		fields = append(fields, f.GetName())
	}
	if len(fields) > 0 {
		return db.Select(fields)
	}
	return db
}
