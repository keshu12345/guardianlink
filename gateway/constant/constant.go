package constant

type Gateway string

const (
	Databasename Gateway = "gateway.db"
	Drivername   Gateway = "sqlite3"
	NodeBURL     Gateway = "http://localhost:8081"
	NodeAURL     Gateway = "http://localhost:8080"
)

func (n Gateway) ToString() string {
	return string(n)
}
