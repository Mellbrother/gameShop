package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"		// フレームワーク

	"github.com/jinzhu/gorm"	// ORMのためのライブラリ
	"github.com/swaggo/http-swagger"	// apiドキュメント作成で使う
	_ "test/docs"
	resp "github.com/nicklaw5/go-respond"	// レスポンスを返すときに使う
	_ "github.com/go-sql-driver/mysql"
)

type (
	item struct {
		ID           	uint64      `json:"id"`
		Name		 	string 		`json:"name"`
		Price 			uint64 		`json:"price"`
		ExhibitorAdress string      `json:"exhibitor_addrss"`
		Description  	string      `json:"description"`
		IsSold 	 	  	bool        `json:"is_sold"`
	}

	items []item
)

// API Documents
// @title Sample API
// @version 1.0.0
func main() {
	log.Println("Starting API Server!!")

	db := connectDB()
	defer db.Close()

	router := chi.NewRouter()

	// @Tags Item
	// @Accept json
	// @Summary 商品情報を取得します
	// @Success 200 {array} item
	// @Router /items [get]
	router.Get("/items", func(w http.ResponseWriter, r *http.Request) {
		items := []item{}
		db.Find(&items)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		resp.NewResponse(w).Ok(items)
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", router)
}

// Gormを利用してDBに接続する
func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql)"
	DBNAME := "test_db"
	option := "?charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + option

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
	panic(err.Error())
	}
	return db
}
