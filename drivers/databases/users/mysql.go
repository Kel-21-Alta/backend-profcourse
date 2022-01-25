package users

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/users"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func (m mysqlUserRepository) GetCourseUser(ctx context.Context, domain *users.Domain) ([]users.Course, error) {
	type Result struct {
		CourseName string
		Progress int
		Skor int
		CourseId string
		UserId string
		ID string
	}

	var rec []Result

	err := m.Conn.Table("users_courses").Select("users_courses.id as id, progress, skor, course_id, user_id, courses.title as CourseName").Joins("INNER JOIN courses ON users_courses.course_id = courses.id").Where("user_id = ?", domain.ID).Scan(&rec).Error

	if err !=nil {
		return []users.Course{}, err
	}
	
	var list = []users.Course{}

	for _, course := range rec {
		list = append(list, users.Course{
			ID:          course.ID,
			UserId:      course.UserId,
			CourseId:    course.CourseId,
			Progres:     course.Progress,
			Score:       course.Skor,
			CourseTitle: course.CourseName,
		})

	}
	
	return list,nil
}

// Paginate Fungsi ini untuk mengimplementasikan pagination pada list course
func Paginate(domain users.Domain) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := domain.Query.Offset
		limit := domain.Query.Limit
		if limit == 0 {
			limit = 10
		}
		return db.Offset(offset).Limit(limit)
	}
}

func (m mysqlUserRepository) GetAllUser(ctx context.Context, domain *users.Domain) ([]users.Domain, error) {
	type Result struct {
		Name string
		ImgProfile string
		Id string
		TakenCourse int
		Point int
	}
	var rec []Result

	err := m.Conn.Scopes(Paginate(*domain)).Table("users").Select("name, img_profile, users.id as Id, COUNT(users_courses.user_id) as TakenCourse, SUM(users_courses.skor) as Point, users.created_at as created_at").Joins("LEFT JOIN users_courses ON users_courses.user_id = users.id").Group("Id").Where("name Like ?", "%"+domain.Query.Search+"%").Order(domain.Query.SortBy + " " + domain.Query.Sort).Scan(&rec).Error

	if err != nil {
		return []users.Domain{}, err
	}
	var listDomain []users.Domain

	for _, re := range rec {
		listDomain = append(listDomain, users.Domain{
			ID:           re.Id,
			Name:         re.Name,
			ImgProfile:   re.ImgProfile,
			TakenCourse:  re.TakenCourse,
			Point:        re.Point,
		})
	}

	return listDomain, nil
}

func (m mysqlUserRepository) UpdateDataCurrentUser(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	var rec User
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return users.Domain{}, err
	}

	rec.ImgProfile = domain.ImgProfile
	rec.Birth = domain.Birth
	rec.NoHp = domain.NoHp
	rec.Name = domain.Name
	rec.Bio = domain.Bio
	rec.BirthPlace = domain.BirthPlace

	err = m.Conn.Save(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (m mysqlUserRepository) GetCountUser(ctx context.Context) (*users.Summary, error) {
	var result int
	err := m.Conn.Raw("SELECT COUNT(*) as result FROM users").Scan(&result).Error
	if err != nil {
		return &users.Summary{}, nil
	}
	return &users.Summary{CountUser: result}, nil
}

func (m mysqlUserRepository) DeleteUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := User{}
	result := m.Conn.Where("id = ?", domain.IdUser).Delete(&rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	domain.Message = "User dengan id: " + domain.IdUser + " telah dihapus"
	return domain, nil
}

// UpdateUser Update data user dari admin
func (m mysqlUserRepository) UpdateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	first := User{}
	err := m.Conn.First(&first, "id = ?", domain.ID).Error
	first.Name = domain.Name
	first.NoHp = domain.NoHp
	first.Bio = domain.Bio
	first.Birth = domain.Birth
	first.BirthPlace = domain.BirthPlace
	result := m.Conn.Save(&first)
	if err != nil {
		return users.Domain{}, err
	}
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return first.ToDomain(), nil
}

// GetUserById Untuk mendapatkan data user berdasarkan id
func (m mysqlUserRepository) GetUserById(ctx context.Context, id string) (users.Domain, error) {
	rec := User{}
	err := m.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

/*
	Digunakan untuk mendapatkan user dengan email tertentu
**/
func (m mysqlUserRepository) GetUserByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := User{}
	err := m.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

// Digunakan untuk membuat user baru
func (m mysqlUserRepository) CreateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := FromDomain(domain)
	result := m.Conn.Create(&rec)
	if result.Error != nil {
		return rec.ToDomain(), result.Error
	}
	return rec.ToDomain(), nil
}

// untuk mengupdate password
func (m mysqlUserRepository) UpdatePassword(ctx context.Context, domain users.Domain, hash string) (users.Domain, error) {
	rec := FromDomain(domain)
	err := m.Conn.Model(&rec).Update("password", hash).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}
