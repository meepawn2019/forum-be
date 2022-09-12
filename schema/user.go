package schema

type user struct {
	id       int64  `db:"id"`
	name     string `db:"name"`
	username string `db:"username"`
	email    string `db:"email"`
	phone    string `db:"phone"`
	password string `db:"password"`
}
