package repository

import (
	"context"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	tableMerchant = "toko"
	// layoutDateTime = "2006-01-02 15:04:05"
)

func AddMerchant(ctx context.Context, toko model.Toko, user_id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (user_id,name,alamat,created_at, updated_at) values(%v,'%v','%v', NOW(), NOW())", tableMerchant,
		user_id,
		toko.Name,
		toko.Alamat,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMerchant(ctx context.Context, toko model.Toko, tokoid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET name = '%v', alamat = '%v' WHERE id = %v", tableMerchant,
		toko.Name,
		toko.Alamat,
		tokoid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
