package repository

import (
	"context"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	table_detail_user = "detail_user"
	// layoutDateTime = "2006-01-02 15:04:05"
)

func GetDetailUser(ctx context.Context, user_id int) ([]model.DetailUser, error) {
	var detail_users []model.DetailUser
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v where user_id = %v", table_detail_user, user_id)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var detail_user model.DetailUser
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&detail_user.ID,
			&detail_user.User_Id,
			&detail_user.FullName,
			&detail_user.Alamat,
			&detail_user.Phone,
			&detail_user.Avatar,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// detail_user.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// detail_user.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		detail_users = append(detail_users, detail_user)
	}
	return detail_users, nil
}

func UpdateDetailUser(ctx context.Context, detailuser model.DetailUser, user_id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET fullname = '%v', alamat = '%v', phone = '%v', avatar = '%v' WHERE user_id = %v;",
		table_detail_user,
		detailuser.FullName,
		detailuser.Alamat,
		detailuser.Phone,
		detailuser.Avatar,
		user_id,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
