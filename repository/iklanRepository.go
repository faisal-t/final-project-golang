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
	tableIklan = "iklan"
	// layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAllIklan(ctx context.Context) ([]model.Iklan, error) {
	var iklans []model.Iklan
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", tableIklan)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var iklan model.Iklan
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&iklan.ID,
			&iklan.Thumbnail,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// iklan.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// iklan.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		iklans = append(iklans, iklan)
	}
	return iklans, nil
}

func AddIklan(ctx context.Context, iklan model.Iklan) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (image_url,created_at, updated_at) values('%v', NOW(), NOW())", tableIklan,
		iklan.Thumbnail,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateIklan(ctx context.Context, iklan model.Iklan, iklanid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET image_url = '%v' WHERE id = %v", tableIklan,
		iklan.Thumbnail,
		iklanid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteIklan(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", tableIklan, id)

	fmt.Println(queryText)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
