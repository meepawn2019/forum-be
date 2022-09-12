package schema

type post struct {
	id       int64  `db:"id"`
	title    string `db:"title"`
	content  string `db:"content"`
	author   int64  `db:"author"`
	created  string `db:"created"`
	modified string `db:"modified"`
}
