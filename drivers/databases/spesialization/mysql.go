package spesialization

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/spesializations"
)

type spesializationRepository struct {
	Conn *gorm.DB
}

func (r *spesializationRepository) GetCountSpesializations(ctx context.Context) (spesializations.Summary, error) {
	var result int
	err := r.Conn.Raw("SELECT COUNT(*) as result FROM spesializations").Scan(&result).Error
	if err != nil {
		return spesializations.Summary{}, nil
	}
	return spesializations.Summary{CountSpesialization: result}, nil
}

func (r *spesializationRepository) GetOneSpesialization(ctx context.Context, domain *spesializations.Domain) (spesializations.Domain, error) {
	var rec Spesialization
	rec.ID = domain.ID
	err := r.Conn.Preload("Courses").Find(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return spesializations.Domain{}, err
	}
	return rec.ToDomainWithCourses(), nil
}

// Paginate Fungsi ini untuk mengimplementasikan pagination pada list course
func Paginate(domain spesializations.Domain) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := domain.Offset
		limit := domain.Limit
		if limit == 0 {
			limit = 10
		}
		return db.Offset(offset).Limit(limit)
	}
}

// GetAllSpesializations Untuk mendapatkan banyak spesialization
func (r *spesializationRepository) GetAllSpesializations(ctx context.Context, domain *spesializations.Domain) ([]spesializations.Domain, error) {
	var spesializationResult []*Spesialization
	var err error
	err = r.Conn.Scopes(Paginate(*domain)).Order(domain.Sort+" "+domain.SortBy).Where("title Like ?", "%"+domain.KeywordSearch+"%").Find(&spesializationResult).Error

	if err != nil {
		return []spesializations.Domain{}, err
	}

	return ToListDomain(spesializationResult), nil
}

func (s *spesializationRepository) CreateSpasialization(ctx context.Context, domain *spesializations.Domain) (spesializations.Domain, error) {

	req := FromDomain(domain)

	err := s.Conn.Omit("Courses.*").Create(&req).Error

	if err != nil {
		return spesializations.Domain{}, err
	}

	return req.ToDomainWithCourses(), nil
}

func NewMysqlRepository(conn *gorm.DB) spesializations.Repository {
	return &spesializationRepository{Conn: conn}
}
