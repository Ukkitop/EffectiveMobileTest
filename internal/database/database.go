package database

type DatabaseInterface interface {
	GetData() string
	DeleteData(id string) bool
	UpdateData(username, password string) string
	SetupDatabase() error
}
