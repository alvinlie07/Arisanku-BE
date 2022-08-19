package authentication

import "gorm.io/gorm"

type Repository interface {
	Login(Email string, Password string) (User, error)
	Register(user User) (User, error)
	ForgetPassword(Email string) (User, error)
	ResetPassword(Password string) (User, error)

	FindAll() ([]User, error)
	FindById(id int) (User, error)

	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id int) (User, error)

	CheckingEmail(Email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(Email string, Password string) (User, error) {
	var user User

	err := r.db.Debug().Where("email = ? AND password = ?", Email, Password).First(&user).Error

	return user, err
}

func (r *repository) Register(user User) (User, error) {
	// var user User

	err := r.db.Debug().Create(&user).Error

	return user, err
}

func (r *repository) ForgetPassword(Email string) (User, error) {
	var user User

	err := r.db.Debug().Where("email = ?", Email).Find(&user).Error

	return user, err
}

func (r *repository) ResetPassword(Password string) (User, error) {
	var user User

	err := r.db.Debug().Update("password = ?", Password).Error

	return user, err
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Debug().Find(&users).Error

	return users, err
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Debug().Find(&user, id).Error

	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Debug().Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user User) (User, error) {

	err := r.db.Debug().Save(&user).Where("id = ?", user.ID).Error

	return user, err
}

func (r *repository) DeleteUser(id int) (User, error) {
	var user User
	err := r.db.Debug().Delete(&user, id).Error

	return user, err
}

func (r *repository) CheckingEmail(Email string) (User, error) {
	var user User

	err := r.db.Debug().Where("email = ?", Email).Find(&user).Error

	return user, err
}
