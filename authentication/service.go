package authentication

type Service interface {
	Login(param LoginParameter) (User, error)
	Register(param RegisterParameter) (User, error)
	ForgetPassword(param ForgetPasswordParameter) (User, error)
	ResetPassword(param ResetPasswordParameter) (User, error)

	FindById(id int) (User, error)
	FindAll() ([]User, error)

	UpdateUser(param UpdateUserParamater) (User, error)
	DeleteUser(id int) (User, error)

	CheckingEmail(email string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Login(param LoginParameter) (User, error) {
	return s.repository.Login(param.Email, param.Password)
}

func (s *service) Register(param RegisterParameter) (User, error) {
	registerParam := User{
		Fullname: param.Fullname,
		Email:    param.Email,
		Password: param.Password,
	}
	return s.repository.Register(registerParam)
}

func (s *service) ForgetPassword(param ForgetPasswordParameter) (User, error) {
	return s.repository.ForgetPassword(param.Email)
}

func (s *service) ResetPassword(param ResetPasswordParameter) (User, error) {
	return s.repository.ResetPassword(param.Password)
}

func (s *service) FindById(id int) (User, error) {
	return s.repository.FindById(id)
}
func (s *service) FindAll() ([]User, error) {
	return s.repository.FindAll()
}

func (s *service) UpdateUser(param UpdateUserParamater) (User, error) {
	idUser, _ := param.ID.Int64()

	updateParam := User{
		ID:       int(idUser),
		Fullname: param.Fullname,
		Email:    param.Email,
		Password: param.Password,
	}
	return s.repository.UpdateUser(updateParam)
}

func (s *service) DeleteUser(id int) (User, error) {
	return s.repository.DeleteUser(id)
}

func (s *service) CheckingEmail(email string) (User, error) {
	return s.repository.CheckingEmail(email)
}
