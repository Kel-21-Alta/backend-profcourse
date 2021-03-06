package users

import (
	"context"
	"profcourse/app/middlewares"
	"profcourse/business/send_email"
	controller "profcourse/controllers"
	"profcourse/helpers"
	"profcourse/helpers/encrypt"
	"profcourse/helpers/randomString"
	"profcourse/helpers/validators"
	"time"
)

type userUsecase struct {
	ContextTimeout time.Duration
	UserRepository Repository
	SmtpRepository send_email.Repository
	JWTConfig      middlewares.ConfigJwt
	PDFGenerater PDF
}


func (u *userUsecase) GenerateReportUser(ctx context.Context, domain *Domain) (Domain, error) {

	user, err := u.UserRepository.GetUserById(ctx, domain.ID)

	if err != nil {
		return Domain{}, err
	}

	course, err := u.UserRepository.GetCourseUser(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	result, err := u.PDFGenerater.GeneratePDFDataReport(user, course)

	if err != nil {
		return Domain{}, err
	}

	return Domain{FileReport: result}, nil
}

func (u *userUsecase) GetAllUser(ctx context.Context, domain *Domain) ([]Domain, error) {
	if domain.Role != 1 {
		return []Domain{}, controller.FORBIDDIN_USER
	}

	if domain.Query.Sort == "" {
		domain.Query.Sort = "asc"
	}

	if domain.Query.Sort == "dsc" {
		domain.Query.Sort = "desc"
	}

	if domain.Query.SortBy == "" {
		domain.Query.SortBy = "created_at"
	}

	// menvalidasi sort by yang diizinkan
	sortByAllow := []string{"asc", "desc"}
	if !helpers.CheckItemInSlice(sortByAllow, domain.Query.Sort) {
		return []Domain{}, controller.INVALID_PARAMS
	}

	// Menvalidasi sort yang diizinkan
	sortAllow := []string{"created_at", "name", "point", "jumlah_kursus"}
	if !helpers.CheckItemInSlice(sortAllow, domain.Query.SortBy) {
		return []Domain{}, controller.INVALID_PARAMS
	}

	result, err := u.UserRepository.GetAllUser(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (u *userUsecase) UpdateDataCurrentUser(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_EMPTY
	}

	if domain.Name == "" {
		return Domain{}, controller.EMPTY_NAME
	}

	result, err := u.UserRepository.UpdateDataCurrentUser(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *userUsecase) GetCountUser(ctx context.Context) (*Summary, error) {
	summary, err := u.UserRepository.GetCountUser(ctx)
	if err != nil {
		return &Summary{}, err
	}
	return summary, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, domain Domain) (Domain, error) {
	// Cek apakah yang mengirimkan request adalah admin
	if domain.Role != 1 {
		return Domain{}, controller.FORBIDDIN_USER
	}
	// Cek apakah admin mengirim user id yang akan dihapus
	if domain.IdUser == "" {
		return Domain{}, controller.ID_EMPTY
	}

	result, err := u.UserRepository.DeleteUser(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, domain Domain) (Domain, error) {
	// Cek apakah yang mengirimkan request adalah admin
	if domain.Role != 1 {
		return Domain{}, controller.FORBIDDIN_USER
	}
	// Cek apakah admin mengirim user id yang akan dihapus
	if domain.IdUser == "" {
		return Domain{}, controller.ID_EMPTY
	}
	if domain.ID == "" {
		return Domain{}, controller.ID_EMPTY
	}
	// Cek apakah Nama dikosongkan
	if domain.Name == "" {
		return Domain{}, controller.EMPTY_NAME
	}
	domain.ID = domain.IdUser
	domain.Role = 2
	result, err := u.UserRepository.UpdateUser(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u *userUsecase) LoginAdmin(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain

	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}
	if domain.Password == "" {
		return Domain{}, controller.PASSWORD_EMPTY
	}

	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	// Mengecek email apakah benar ada usernya
	if existedUser == (Domain{}) {
		return Domain{}, controller.WRONG_EMAIL
	}

	// cek apakah user yang login adalah admin
	if existedUser.Role != 1 {
		return Domain{}, controller.FORBIDDIN_USER
	}

	if err != nil {
		return Domain{}, err
	}

	// Mengecek apakan passwordnya benar
	if !encrypt.ValidateHash(domain.Password, existedUser.HashPassword) {
		return Domain{}, controller.WRONG_PASSWORD
	}

	existedUser.Token, err = u.JWTConfig.GenrateTokenJWT(existedUser.ID, existedUser.Role, existedUser.RoleText)

	if err != nil {
		return Domain{}, err
	}

	return existedUser, nil
}

func (u *userUsecase) ChangePassword(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_EMPTY
	}
	// Cocokkan password lama dengan password yang sekarang
	user, err := u.UserRepository.GetUserById(ctx, domain.ID)

	if user == (Domain{}) {
		return Domain{}, controller.EMPTY_USER
	}

	if err != nil {
		return Domain{}, err
	}

	if !encrypt.ValidateHash(domain.Password, user.HashPassword) {
		return Domain{}, controller.WRONG_PASSWORD
	}

	// Update password lama dengan password baru
	hashPasswordNew, err := encrypt.Hash(domain.PasswordNew)
	if err != nil {
		return Domain{}, err
	}
	user, err = u.UserRepository.UpdatePassword(ctx, user, hashPasswordNew)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (u *userUsecase) GetCurrentUser(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_EMPTY
	}
	user, err := u.UserRepository.GetUserById(ctx, domain.ID)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (u *userUsecase) ForgetPassword(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain
	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}

	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	// cek apakah email tersebut terdaftar
	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	if existedUser == (Domain{}) {
		return Domain{}, controller.WRONG_EMAIL
	}

	if err != nil {
		return Domain{}, err
	}

	// Membuat password baru
	domain.Password = randomString.RandomString(8)
	domain.HashPassword, err = encrypt.Hash(domain.Password)

	if err != nil {
		return Domain{}, err
	}

	resultUser, err := u.UserRepository.UpdatePassword(ctx, existedUser, domain.HashPassword)
	if err != nil {
		return Domain{}, err
	}

	// Mengirim password dengan email
	to := resultUser.Email
	subject := "Lupa Password Akun Profcouse"
	message := "" +
		"<img src=\"https://firebasestorage.googleapis.com/v0/b/crudfirebase-91413.appspot.com/o/logo.png?alt=media&token=4fa0b90f-6b13-41f3-96a3-53e277ff4d5c\" alt=\"Logo Prof Course\" width=\"75\">" +
		"<p>Dear " + resultUser.Name + "</p><br><p>Password anda telah kami reset ulang dan password anda sekarang adalah : " + domain.Password + " " + "" +
		"<p>Anda harus menjaga informasi anda</p>" +
		"<br>" +
		"<p>Terima kasih</p>" +
		"<br>" +
		"Prof Course"

	err = u.SmtpRepository.SendEmail(ctx, to, subject, message)
	if err != nil {
		return Domain{}, err
	}

	return resultUser, nil
}

func (u *userUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain

	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}
	if domain.Password == "" {
		return Domain{}, controller.PASSWORD_EMPTY
	}

	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	// Mengecek email apakah benar ada usernya
	if existedUser == (Domain{}) {
		return Domain{}, controller.WRONG_EMAIL
	}

	if err != nil {
		return Domain{}, err
	}

	// Mengecek apakan passwordnya benar
	if !encrypt.ValidateHash(domain.Password, existedUser.HashPassword) {
		return Domain{}, controller.WRONG_PASSWORD
	}

	existedUser.Token, err = u.JWTConfig.GenrateTokenJWT(existedUser.ID, existedUser.Role, existedUser.RoleText)

	if err != nil {
		return Domain{}, err
	}

	return existedUser, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var existedUser Domain

	if domain.Name == "" {
		return Domain{}, controller.EMPTY_NAME
	}

	if domain.Email == "" {
		return Domain{}, controller.EMPTY_EMAIL
	}

	// Mengecek apakah email yang diberika valid
	if !validators.CheckEmail(domain.Email) {
		return Domain{}, controller.INVALID_EMAIL
	}

	// Mengecek apakan Email telah digunakan atau belum
	existedUser, err = u.UserRepository.GetUserByEmail(ctx, domain.Email)

	if existedUser != (Domain{}) {
		return Domain{}, controller.EMAIL_UNIQUE
	}

	// Melakukan hashing pada password
	domain.Password = randomString.RandomString(8)
	domain.HashPassword, err = encrypt.Hash(domain.Password)

	if err != nil {
		return Domain{}, err
	}

	domain.CreatedAt = time.Now()
	domain.UpdatedAt = time.Now()

	// Mengirim password dengan email
	to := domain.Email
	subject := "Pendaftaran akun di Profcouse"
	message := "" +
		"<img src=\"https://firebasestorage.googleapis.com/v0/b/crudfirebase-91413.appspot.com/o/logo.png?alt=media&token=4fa0b90f-6b13-41f3-96a3-53e277ff4d5c\" alt=\"Logo Prof Course\" width=\"75\">" +
		"<p>Dear " + domain.Name + "</p><br><p>Password anda pada akun profcourse adalah : " + domain.Password + " " + "" +
		"<p>Anda harus menjaga informasi anda</p>" +
		"<br>" +
		"<p>Terima kasih</p>" +
		"<br>" +
		"Prof Course"
	err = u.SmtpRepository.SendEmail(ctx, to, subject, message)
	if err != nil {
		return Domain{}, err
	}

	// Mengirim domain to layer mysql repository user
	var resultDomain Domain
	resultDomain, err = u.UserRepository.CreateUser(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}

func NewUserUsecase(r Repository, timeout time.Duration, smtpRepo send_email.Repository, pdf PDF, configJwt middlewares.ConfigJwt) Usecase {
	return &userUsecase{
		ContextTimeout: timeout,
		UserRepository: r,
		SmtpRepository: smtpRepo,
		JWTConfig:      configJwt,
		PDFGenerater: pdf,
	}
}
