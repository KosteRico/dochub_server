package user

import (
	"checkaem_server/cmd/handlers/util"
	pwd "checkaem_server/cmd/utill/password"
)

type User struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Token    map[string]string `json:"token,omitempty"`
}

func NewEmpty() *User {
	return &User{}
}

func New(username, password string) (User, error) {

	password, err := pwd.Hash(password)

	if err != nil {
		return User{}, err
	}

	tokenString, err := util.GenerateTokenPair(username)

	if err != nil {
		return User{}, err
	}

	return User{
		Username: username,
		Password: password,
		Token:    tokenString,
	}, nil
}

func (u *User) Create() (map[string]interface{}, error) {
	hashedPassword, err := pwd.Hash(u.Password)

	if err != nil {
		return nil, err
	}

	tokenString, err := util.GenerateTokenPair(u.Username)

	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword
	u.Token = tokenString

	response := util.Message(true, "Account has been created")
	response["account"] = u

	return response, nil
}

func (u *User) ComparePassword(password string) bool {
	return pwd.Compare(u.Password, password)
}
