package db

import (
	"database/sql"
	"fmt"
	"github.com/ethien-salinas/echo-rest-api/utils"
	pg "github.com/lib/pq"
)

const (
	db_host     = "lab.catalina.tech"
	db_port     = 5432
	db_user     = "intelligential_user"
	db_password = "dOcFF29D4dxjyqF1BeGRA"
	db_name     = "go_test"
)

func ValidateUser(email string, password string) string {
	_ = pg.Efatal
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	db, err := sql.Open("postgres", psqlconn)
	utils.CheckError(err)

	defer db.Close()

	rows, err2 := db.Query(`SELECT "name", "email", "password" FROM "user" WHERE "email" = $1`, email)
	utils.CheckError(err2)
	defer rows.Close()

	for rows.Next() {
		var nameDB string
		var emailDB string
		var passwordDB string

		err = rows.Scan(&nameDB, &emailDB, &passwordDB)
		utils.CheckError(err)
		if password == passwordDB {
			return nameDB
		}
	}

	return ""
}
