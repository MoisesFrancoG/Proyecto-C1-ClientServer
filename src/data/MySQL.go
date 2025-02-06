package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL() *MySQL {
	dsn := "root:12345678@tcp(localhost:3306)/testgo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[MySQL] Error al conectar con la base de datos: %v", err)
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("[MySQL] Error al hacer ping a la base de datos: %v", err)
	}

	log.Println("[MySQL] Conexión establecida correctamente")
	return &MySQL{DB: db}
}

func (m *MySQL) Close() {
	if err := m.DB.Close(); err != nil {
		log.Printf("[MySQL] Error al cerrar la conexión: %v", err)
	} else {
		log.Println("[MySQL] Conexión cerrada correctamente")
	}
}
