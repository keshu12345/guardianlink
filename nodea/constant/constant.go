package constant

type Nodea string

const (
	Databasename Nodea = "nodea.db"
	Drivername   Nodea = "sqlite3"
	NodeBURL     Nodea = "http://localhost:8081/api/blocks"
)

func (n Nodea) ToString() string {
	return string(n)
}
