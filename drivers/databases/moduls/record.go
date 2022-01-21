package moduls

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/moduls"
	"profcourse/drivers/databases/materies"
	"time"
)

type Moduls struct {
	ID        string `gorm:"primaryKey;unique"`
	Title     string `gorm:"not null"`
	Order     int    `gorm:"not null"`
	CourseId  string `gorm:"not null;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Materies  []materies.Materi `gorm:"foreignKey:ModulID;references:ID"`
}

type SkorUserModul struct {
	ID           string `gorm:"primaryKey;unique"`
	Nilai        int    `gorm:"not null;default:0"`
	ModulId      string `gorm:"not null;size:191;index:idx_unique2,unique"`
	UserCourseId string `gorm:"not null;size:191;index:idx_unique2,unique"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *SkorUserModul) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *SkorUserModul) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func (c *Moduls) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *Moduls) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func (c Moduls) ToDomain() moduls.Domain {
	var listMateri []moduls.Materi
	var jumlahMateri int
	for _, materi := range c.Materies {
		jumlahMateri++
		var typeStr string
		if int(materi.Type) == 1 {
			typeStr = "text"
		} else {
			typeStr = "video"
		}
		listMateri = append(listMateri, moduls.Materi{
			UrlMateri:   materi.UrlMateri,
			Type:        int(materi.Type),
			TypeString:  typeStr,
			Title:       materi.Title,
			Order:       materi.Order,
			IsComplate:  false,
			CurrentTime: "",
		})
	}
	return moduls.Domain{
		ID:           c.ID,
		Title:        c.Title,
		Order:        c.Order,
		CourseId:     c.CourseId,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
		JumlahMateri: jumlahMateri,
		Materi:       listMateri,
	}
}

func FromDomain(domain *moduls.Domain) *Moduls {
	return &Moduls{
		Title:    domain.Title,
		Order:    domain.Order,
		CourseId: domain.CourseId,
	}
}

func TolistDomain(recs []Moduls) []moduls.Domain {
	var listDomain []moduls.Domain

	for _, modul := range recs {
		listDomain = append(listDomain, moduls.Domain{
			ID:        modul.ID,
			Title:     modul.Title,
			Order:     modul.Order,
			CourseId:  modul.CourseId,
			CreatedAt: modul.CreatedAt,
			UpdatedAt: modul.UpdatedAt,
		})
	}

	return listDomain
}

func (c *SkorUserModul) ToDomain() moduls.ScoreUserModul {
	return moduls.ScoreUserModul{
		ID:           c.ID,
		Nilai:        c.Nilai,
		ModulID:      c.ModulId,
		UserCourseId: c.UserCourseId,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func FromDomainToScoreUserModul(domain *moduls.ScoreUserModul) SkorUserModul {
	return SkorUserModul{
		ID:           domain.UserCourseId,
		Nilai:        domain.Nilai,
		ModulId:      domain.ModulID,
		UserCourseId: domain.UserCourseId,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
