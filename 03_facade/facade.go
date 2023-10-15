package main

import "fmt"

// IUser 用户接口
type IUser interface {
	Login(phone int, code int) *User
	Register(Phone int, code int) *User
}

// IUserFacade 门面模式
type IUserFacade interface {
	LoginOrRegister(phone int, code int) *User
}

// User 用户
type User struct {
	Name string
}

type UserService struct {
}

// Login 登录
func (u UserService) Login(phone int, code int) *User {
	if code != 1234 { // 校验操作...
		return nil
	}
	return &User{Name: "用户通过存在校验，login success"}
}

// Register 注册
func (u UserService) Register(phone int, code int) *User {
	// 校验操作...
	return &User{Name: "register success"}
}

// LoginOrRegister 登录或注册
func (u UserService) LoginOrRegister(phone int, code int) *User {
	user := u.Login(phone, code)
	if user != nil {
		return user
	}
	return u.Register(phone, code)
}

func main() {
	phone := 1234567890
	code := 1234

	userService := UserService{}
	user := userService.LoginOrRegister(phone, code)

	fmt.Println(user.Name)
}
