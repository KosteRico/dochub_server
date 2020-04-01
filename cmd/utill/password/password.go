package password

import "golang.org/x/crypto/bcrypt"

//Hash returns hash of a password as a string
func Hash(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(res), err
}

//Compare compares user password and hashed password which is stored in database
func Compare(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
