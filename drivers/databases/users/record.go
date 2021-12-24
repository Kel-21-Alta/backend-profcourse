package users

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/users"
	"time"
)

type User struct {
	ID         string `gorm:"primaryKey;unique"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Email      string `gorm:"not null;unique"`
	Role       int8   `gorm:"not null;default:2"`
	NoHp       string
	BirthPlace string
	Bio        string
	ImgProfile string
	RoleText   string
	Birth      time.Time `gorm:"default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	if u.Role == 1 {
		u.RoleText = "Admin"
	} else {
		u.RoleText = "User"
	}
	return
}

func (u User) ToDomain() users.Domain {
	return users.Domain{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		HashPassword: u.Password,
		NoHp:         u.NoHp,
		Birth:        u.Birth,
		BirthPlace:   u.BirthPlace,
		Bio:          u.Bio,
		ImgProfile:   u.ImgProfile,
		Role:         u.Role,
		RoleText:     u.RoleText,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:         domain.ID,
		Name:       domain.Name,
		Password:   domain.HashPassword,
		Role:       domain.Role,
		RoleText:   domain.RoleText,
		Email:      domain.Email,
		NoHp:       domain.NoHp,
		BirthPlace: domain.BirthPlace,
		Bio:        domain.Bio,
		ImgProfile: domain.ImgProfile,
		Birth:      domain.Birth,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
