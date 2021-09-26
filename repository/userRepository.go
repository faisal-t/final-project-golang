package repository

import (
	"context"
	"database/sql"
	"errors"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	table          = "user"
	layoutDateTime = "2021-09-26 15:19:40"
)

func Register(ctx context.Context, user model.User) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// queryText := fmt.Sprintf("INSERT INTO %v (username,password,role,created_at, updated_at) values('%v','%v','user', NOW(), NOW())",
	// 	table,
	// 	user.Username,
	// 	user.Password,
	// )

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	query1 := fmt.Sprintf("INSERT INTO %v (username, password,role,created_at,updated_at) VALUES('%v', '%v','user',NOW(),NOW());",
		table,
		user.Username,
		user.Password)

	_, err = tx.ExecContext(ctx, query1)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return err
	}

	query2 := "INSERT INTO detail_user (user_id,created_at,updated_at) VALUES(LAST_INSERT_ID(),NOW(),NOW());"

	_, err = tx.ExecContext(ctx, query2)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// _, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

func Login(ctx context.Context, username string, password string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v where username = '%v' AND password = '%v'", table,
		username,
		password,
	)
	fmt.Println(queryText)
	rowQuery := db.QueryRow(queryText)

	if err != nil {
		return err
	}

	switch err := rowQuery.Scan(); err {
	case sql.ErrNoRows:
		return errors.New("Username atau password salah")

	}

	return nil
}
