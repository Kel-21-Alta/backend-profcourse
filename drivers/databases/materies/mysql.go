package materies

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"profcourse/business/materies"
)

type MateriesRepository struct {
	Conn *gorm.DB
}

func (m MateriesRepository) GetCountMateriFinish(ctx context.Context, domain *materies.Domain) (int, error) {
	var result int
	err := m.Conn.Table("materi_user_complates").Select("COUNT(materi_id) as result").Where("user_course_id = ?", domain.UserCourse.UserCourseId).Where("is_complate = ?", true).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (m MateriesRepository) GetCountMateriCourse(ctx context.Context, domain *materies.Domain) (int, error) {
	var result int
	// jumlah materi => SELECT COUNT(*) FROM materis INNER JOIN moduls m on materis.modul_id = m.id INNER JOIN courses c on m.course_id = c.id where course_id = "a7234d7d-ebc5-495c-ad41-782f3eb906b8"
	err := m.Conn.Table("materis").Select("COUNT(*) as result").Joins("INNER JOIN moduls m on materis.modul_id = m.id").Joins("INNER JOIN courses c on m.course_id = c.id").Where("course_id = ?", domain.User.CourseId).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (m MateriesRepository) UpdateProgressMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {

	var rec = FromDomainToMateriUserComplate(domain)
	if m.Conn.Model(&rec).Where("materi_id = ?", rec.MateriId).Where("user_course_id = ?", rec.UserCourseID).Updates(&rec).RowsAffected == 0 {
		m.Conn.Create(&rec)
	}

	var rec2 Materi

	err := m.Conn.First(&rec2, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}
	result := *domain

	result.ModulId = rec2.ModulID

	return result, nil
}

func (m MateriesRepository) GetAllMateri(ctx context.Context, domain *materies.Domain) (materies.AllMateriModul, error) {

	var rec []Materi

	err := m.Conn.Preload("MateriUserComplate").Where("modul_id = ?", domain.ModulId).Find(&rec).Error

	if err != nil {
		return materies.AllMateriModul{}, err
	}

	result := ToAllMateriModul(rec, domain.UserCourse.UserCourseId)

	return result, nil
}

func (m MateriesRepository) GetOnemateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {

	var rec Materi
	var err error

	err = m.Conn.Preload(clause.Associations).First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}
	result := rec.ToDomain()
	result.User.ID = domain.User.ID

	for _, user := range rec.MateriUserComplate {
		if user.UserCourse.UserId == domain.User.ID {
			result.User.IsComplate = user.IsComplate
			result.User.CurrentTime = user.CurrentTime
			result.User.ID = user.ID
		}
	}

	return result, nil
}

func (m MateriesRepository) UpdateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {

	var rec Materi
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}

	rec.Title = domain.Title
	rec.UrlMateri = domain.UrlMateri
	rec.ModulID = domain.ModulId
	rec.Type = TYPE(domain.Type)
	rec.Order = int8(domain.Order)

	err = m.Conn.Save(&rec).Error

	if err != nil {
		return materies.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m MateriesRepository) DeleteMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	var rec Materi
	err := m.Conn.Delete(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}

	return *domain, nil
}

func (m MateriesRepository) CreateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	var req = FromDomain(domain)

	err := m.Conn.Create(&req).Error
	if err != nil {
		return materies.Domain{}, err
	}

	return req.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) materies.Repository {
	return &MateriesRepository{Conn: conn}
}
