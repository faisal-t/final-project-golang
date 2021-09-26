package main

import (
	"context"
	"encoding/json"
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//user
	router.POST("/user/register", RegisterUser)
	// router.POST("/user/login", LoginUser)
	//toko
	router.POST("/toko/store/:userid", BasicAuth(PostToko))
	router.PUT("/toko/update/:tokoid", BasicAuth(UpdateToko))
	//kategori
	router.POST("/kategori/store", BasicAuth(UpdateToko))
	router.PUT("/kategori/update/:id", BasicAuth(UpdateToko))
	//iklan
	router.GET("/iklan", GetIklan)
	router.POST("/iklan/store", BasicAuth(PostIklan))
	router.PUT("/iklan/update/:id", BasicAuth(UpdateIklan))
	router.DELETE("/iklan/delete/:id", BasicAuth(DeleteIklan))

	//detailUser
	router.GET("/detailuser/:user_id", BasicAuth(GetDetailUser))
	router.PUT("/detailuser/update/:user_id", BasicAuth(UpdateDetailUser))

	//produk
	router.GET("/produk", GetProducts)
	router.POST("/produk/store/:id_toko", BasicAuth(PostProduct))
	router.PUT("/produk/update/:id_product", BasicAuth(UpdateProduct))
	router.PUT("/produk/delete/:id_product", BasicAuth(Deleteproduct))

	//order
	router.GET("/order/dikeranjang/:user_id", BasicAuth(GetDikeranjang))
	router.GET("/order/transaksi/:user_id", BasicAuth(GetSudahBayar))
	router.POST("/order/keranjang/:id_product", BasicAuth(GetSudahBayar))
	router.PUT("/order/bayar/:id_order", BasicAuth(GetSudahBayar))
	router.PUT("/order/konfirmasi/:id_order", BasicAuth(GetSudahBayar))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server Running at Port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func KonfirmasiPembayaran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idOrder, _ := strconv.Atoi(ps.ByName("id_order"))

	if err := repository.UpdateStatusBayar(ctx, idOrder); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully Melakukan Konfirmasi Pembayaran Produk",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func Pembayaran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		order model.Order
	)

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idOrder, _ := strconv.Atoi(ps.ByName("id_order"))

	if err := repository.BayarOrder(ctx, order, idOrder); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully Melakukan Pembayaran Produk",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func PostKeranjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		order model.Order
	)

	idProduk, _ := strconv.Atoi(ps.ByName("id_product"))

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repository.AddToKeranjang(ctx, order, idProduk); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully add Product to cart",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func GetDikeranjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idUser, _ := strconv.Atoi(ps.ByName("user_id"))
	products, err := repository.GetAllKeranjang(ctx, idUser)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, products, http.StatusOK)
}

func GetSudahBayar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idUser, _ := strconv.Atoi(ps.ByName("user_id"))
	products, err := repository.GetSudahDibayar(ctx, idUser)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, products, http.StatusOK)
}

func GetTransaksi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idUser, _ := strconv.Atoi(ps.ByName("user_id"))
	products, err := repository.GetAllKeranjang(ctx, idUser)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, products, http.StatusOK)
}

func GetProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	products, err := repository.GetAllProduk(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, products, http.StatusOK)
}

func PostProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		product model.Product
	)

	idToko, _ := strconv.Atoi(ps.ByName("id_toko"))

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repository.AddProduk(ctx, product, idToko); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully add Product",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		product model.Product
	)

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idProduct, _ := strconv.Atoi(ps.ByName("id_product"))

	if err := repository.UpdateProduk(ctx, product, idProduct); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully update Produk",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func Deleteproduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idProduct, _ := strconv.Atoi(ps.ByName("id_product"))
	if err := repository.DeleteProduk(ctx, idProduct); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully delete Produk",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func GetDetailUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idUser, _ := strconv.Atoi(ps.ByName("user_id"))
	iklans, err := repository.GetDetailUser(ctx, idUser)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, iklans, http.StatusOK)
}

func UpdateDetailUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		detailuser model.DetailUser
	)

	if err := json.NewDecoder(r.Body).Decode(&detailuser); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idUser, _ := strconv.Atoi(ps.ByName("user_id"))

	if err := repository.UpdateDetailUser(ctx, detailuser, idUser); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully update DetailUser",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func GetIklan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	iklans, err := repository.GetAllIklan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, iklans, http.StatusOK)
}

func PostIklan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		iklan model.Iklan
	)

	if err := json.NewDecoder(r.Body).Decode(&iklan); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repository.AddIklan(ctx, iklan); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully add Iklan",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateIklan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		iklan model.Iklan
	)

	if err := json.NewDecoder(r.Body).Decode(&iklan); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idIklan, _ := strconv.Atoi(ps.ByName("id"))

	if err := repository.UpdateIklan(ctx, iklan, idIklan); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully update Iklan",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteIklan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idIklan = ps.ByName("id")
	if err := repository.DeleteIklan(ctx, idIklan); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully delete Iklan",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := repository.Register(ctx, user); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func PostToko(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		toko model.Toko
	)

	if err := json.NewDecoder(r.Body).Decode(&toko); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idToko, _ := strconv.Atoi(ps.ByName("userid"))

	if err := repository.AddMerchant(ctx, toko, idToko); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully add merchant",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateToko(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		toko model.Toko
	)

	if err := json.NewDecoder(r.Body).Decode(&toko); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idToko, _ := strconv.Atoi(ps.ByName("tokoid"))

	if err := repository.UpdateMerchant(ctx, toko, idToko); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully update merchant",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func PostKategori(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		kategori model.Kategori
	)

	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repository.AddKategori(ctx, kategori); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully add Kategori",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateKategori(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		kategori model.Kategori
	)

	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idKategori, _ := strconv.Atoi(ps.ByName("id"))

	if err := repository.UpdateKategori(ctx, kategori, idKategori); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully update Kategori",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := repository.Login(ctx, user.Username, user.Password); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Berhasil Login",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		username, password, _ := r.BasicAuth()

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		login := repository.Login(ctx, username, password)
		fmt.Println(login)

		if login != nil {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			h(w, r, ps)
		}

	}
}
