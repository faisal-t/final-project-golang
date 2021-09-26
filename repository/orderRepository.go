package repository

import (
	"context"
	"final-project/config"
	"final-project/model"
	"fmt"
	"log"
)

const (
	tableOrder = "`order`"
	// layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllKeranjang(ctx context.Context, userid int) ([]model.Order, error) {
	var orders []model.Order
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v where user_id = %v AND status = 'dikeranjang'", tableOrder, userid)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var order model.Order
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&order.ID,
			&order.User_Id,
			&order.Produk_Id,
			&order.Status,
			&order.AlamatPengiriman,
			&order.Harga,
			&order.Qty,
			&order.TotalBayar,
			&order.BukTiBayar,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// order.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// order.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		orders = append(orders, order)
	}
	return orders, nil
}

func AddToKeranjang(ctx context.Context, order model.Order, produk_id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	totalBayar := order.Qty * order.Harga

	queryText := fmt.Sprintf("INSERT INTO %v (user_id,product_id,status,harga,qty,total_bayar,alamat_pengiriman,bukti_bayar,created_at, updated_at) values(%v,%v,'dikeranjang',%v,%v,%v,'%v','null', NOW(), NOW())", tableOrder,
		order.User_Id,
		produk_id,
		order.Harga,
		order.Qty,
		totalBayar,
		order.AlamatPengiriman,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func BayarOrder(ctx context.Context, order model.Order, orderid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET status = 'konfirmasi admin', bukti_bayar = '%v' WHERE id = %v",
		tableOrder,
		order.BukTiBayar,
		orderid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateStatusBayar(ctx context.Context, orderid int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET status = 'sudah dibayar' WHERE id = %v",
		tableOrder,
		orderid,
	)

	fmt.Println(queryText)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func GetSudahDibayar(ctx context.Context, userid int) ([]model.Order, error) {
	var orders []model.Order
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v where user_id = %v AND status = 'sudah dibayar'", tableOrder, userid)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var order model.Order
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&order.ID,
			&order.User_Id,
			&order.Produk_Id,
			&order.Status,
			&order.AlamatPengiriman,
			&order.Harga,
			&order.Qty,
			&order.TotalBayar,
			&order.BukTiBayar,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// order.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// order.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		orders = append(orders, order)
	}
	return orders, nil
}
