package repository

import (
	"context"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	tableProduk = "produk"
	// layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAllProduk(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v where statusdelete = false", tableProduk)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var produk model.Product
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&produk.ID,
			&produk.Toko_Id,
			&produk.Kategori_Id,
			&produk.Name,
			&produk.Description,
			&produk.Harga,
			&produk.ImageUrl,
			&produk.StatusDelete,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// produk.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// produk.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		products = append(products, produk)
	}
	return products, nil
}

func AddProduk(ctx context.Context, produk model.Product, tokoid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (toko_id,kategori_id,name,description,harga,image_url,statusdelete,created_at, updated_at) values(%v,%v,'%v','%v',%v,'%v',false, NOW(), NOW())",
		tableProduk,
		tokoid,
		produk.Kategori_Id,
		produk.Name,
		produk.Description,
		produk.Harga,
		produk.ImageUrl,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateProduk(ctx context.Context, produk model.Product, produkid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET kategori_id = %v, name = '%v', description='%v', harga=%v, image_url='%v' WHERE id = %v",
		tableProduk,
		produk.Kategori_Id,
		produk.Name,
		produk.Description,
		produk.Harga,
		produk.ImageUrl,
		produkid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteProduk(ctx context.Context, produkid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET statusdelete = true WHERE id = %v",
		tableProduk,
		produkid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
