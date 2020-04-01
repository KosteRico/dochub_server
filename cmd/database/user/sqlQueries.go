package user

const (
	selectAllQuery = "select username from \"user\";"
	selectQuery    = "select username, password from \"user\" where username = $1"
)
