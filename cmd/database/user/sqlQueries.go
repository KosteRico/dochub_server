package user

const (
	insertQuery        = "insert into \"user\" values ($1, $2)"
	deleteQuery        = "delete from \"user\" where username = $1"
	selectAllQuery     = "select username from \"user\";"
	checkIsExistsQuery = "select 1 from \"user\" where username = $1;"
	selectQuery        = "select username, password from \"user\" where username = $1"
)
