package database

var connection string

// func init() akan otomatis dijalankan ketika package ini diimport, jadi kita bisa gunakan untuk inisialisasi package
func init() {
	connection = "MySQL Database"
}

func GetDatabase() string {
	return connection
}
