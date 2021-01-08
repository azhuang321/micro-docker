package service

import (
	"errors"
	"git.imooc.com/cap1573/user/domain/model"
	"git.imooc.com/cap1573/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (userId int64, err error)
	DeleteUser(userId int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(username string) (user *model.User, err error)
	CheckPwd(username string, pwd string) (isOk bool, err error)
}

//创建实例
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//加密用户密码
func GeneratePassword(userPwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPwd), bcrypt.DefaultCost)
}

//验证用户密码
func ValidatePassword(userPwd, hashedPwd string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(userPwd)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

//插入用户
func (u *UserDataService) AddUser(user *model.User) (userId int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

//删除用户
func (u *UserDataService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUserByID(userId)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(username string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(username)
}

func (u *UserDataService) CheckPwd(username string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
