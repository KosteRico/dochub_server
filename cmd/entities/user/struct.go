package user

import pwd "checkaem_server/cmd/utill/password"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(username, password string) (User, error) {

	password, err := pwd.Hash(password)

	if err != nil {
		return User{}, err
	}

	return User{
		Username: username,
		Password: password,
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	return pwd.Compare(u.Password, password)
}
