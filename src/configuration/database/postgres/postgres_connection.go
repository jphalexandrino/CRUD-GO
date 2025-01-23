package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
)

func Init() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// String de conexão
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPassword == "" {
		logger.Info("One or more database configuration variables are missing", zap.String("journey", "Db_Connection"))
		return nil, nil
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to the Database", err, zap.String("journey", "Db_Connection"))
		panic(err)
	}

	// Verifica a conexão
	if err := db.Ping(); err != nil {
		logger.Error("Error when checking connection with bank", err, zap.String("journey", "Db_Connection"))
		panic(err)
	}

	logger.Info("Connection to the database successfully established! - Postgres")
	return db, nil
}
