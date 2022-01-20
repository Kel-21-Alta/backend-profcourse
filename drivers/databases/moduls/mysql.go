package moduls

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/moduls"
)

type mysqlModulsRepository struct {
	Conn *gorm.DB
}

func (m mysqlModulsRepository) CreateScoreModul(ctx context.Context, domain *moduls.ScoreUserModul) (moduls.ScoreUserModul, error) {
	var req = FromDomainToScoreUserModul(domain)
	var err error

	if m.Conn.Model(&req).Where("modul_id = ?", req.ModulId).Where("user_course_id = ?", req.UserCourseId).Updates(&req).RowsAffected == 0 {
		m.Conn.Create(&req)
	}
	if err != nil {
		return moduls.ScoreUserModul{}, err
	}

	return req.ToDomain(), err
}

func (m mysqlModulsRepository) GetAllModulCourse(ctx context.Context, domain *moduls.Domain) ([]moduls.Domain, error) {
	var recs []Moduls

	err := m.Conn.Where("course_id = ?", domain.CourseId).Find(&recs).Order("order asc").Error
	if err != nil {
		return []moduls.Domain{}, err
	}

	return TolistDomain(recs), nil
}

func (m mysqlModulsRepository) DeleteModul(ctx context.Context, id string) (moduls.Message, error) {
	var modul Moduls
	err := m.Conn.Delete(&modul, "id = ?", id).Error

	if err != nil {
		return "", err
	}

	return moduls.Message("Modul dengan id " + id + " telah dihapus"), nil
}

func (m mysqlModulsRepository) UpdateModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var rec Moduls
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	rec.Title = domain.Title
	rec.Order = domain.Order

	err = m.Conn.Save(&rec).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m mysqlModulsRepository) GetOneModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var rec Moduls
	err := m.Conn.Preload("Materies").First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m mysqlModulsRepository) GetOneModulWithCourse(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	type Result struct {
		TeacherId string
	}
	var result Result
	err := m.Conn.Table("moduls").Select("courses.teacher_id as TeacherId").Joins("inner join courses on moduls.course_id = courses.id").Where("moduls.id = ?", domain.ID).Find(&result).Error

	if err != nil {
		return moduls.Domain{}, err
	}
	return moduls.Domain{UserMakeModul: result.TeacherId}, nil
}

func (m mysqlModulsRepository) CreateModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var err error
	rec := FromDomain(domain)

	err = m.Conn.Create(&rec).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) moduls.Repository {
	return &mysqlModulsRepository{Conn: conn}
}
