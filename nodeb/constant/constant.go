package constant

type Nodea string

const (
	Databasename Nodea = "nodeb.db"
	Drivername   Nodea = "sqlite3"
	NodeAURL     Nodea = "http://localhost:8080/api/blocks"
)

func (n Nodea) ToString() string {
	return string(n)
}
