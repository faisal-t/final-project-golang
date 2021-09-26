package repository

import (
	"context"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	tableKategori = "kategori"
	// layoutDateTime = "2006-01-02 15:04:05"
)

func AddKategori(ctx context.Context, kategori model.Kategori) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name,created_at, updated_at) values('%v', NOW(), NOW())", tableKategori,
		kategori.Name,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateKategori(ctx context.Context, toko model.Kategori, tokoid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET name = '%v' WHERE id = %v", tableKategori,
		toko.Name,
		tokoid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
