package handler

import (
	"context"
	"git.imooc.com/cap1573/user/domain/model"
	"git.imooc.com/cap1573/user/domain/service"
	user "git.imooc.com/cap1573/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//注册
func (u User) Register(ctx context.Context, request *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     request.UserName,
		FirstName:    request.FirstName,
		HashPassword: request.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u User) Login(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u User) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(userInfo)
	return nil
}

//类型转换
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
