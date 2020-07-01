package main

// func main(){

// 	for{}
// }

import (
	"github.com/keisuke/gameShop/event/bindings"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/big"
	// "math/rand"
	"os"
	"strings"
)

type Item struct {
	ID           		uint64      `json:"id"`
	Name		 		string 		`json:"name"`
	Price 				uint64 		`json:"price"`
	ExhibitorAddress 	string     `json:"exhibitor_addrss"`
	Description  		string      `json:"description"`
	IsSold 	 	  		bool        `json:"is_sold"`
}

type RegisterItemEvent struct {
	Id         		*big.Int
	Exhibitor		common.Address
	Price		 	*big.Int
}

func main() {
	db := connectDB()
	defer db.Close()

	client, err := ethclient.Dial(os.Getenv("GETH_URL"))
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	topics := map[string]common.Hash{
		"RegisterItem": crypto.Keccak256Hash([]byte("RegisterItem(uint256,address,uint256)")),
	}

	// GAME_SHOP_ADDRESSはまだ宣言していない
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(os.Getenv("GAME_SHOP_ADDRESS"))},
		Topics: [][]common.Hash{{
			topics["RegisterItem"],
		}},
	}

	// コントラクトとの接続
	event := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, event)
	if err != nil {
		log.Fatal(err)
	}

	gameShopAbi, err := abi.JSON(strings.NewReader(string(bindings.GameShopABI)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("test_test_test")

	// event検知
	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
			close(event)
		case vLog := <-event:
			if len(vLog.Topics) == 0 {
				log.Println("Topicsが存在しません")
			}

			switch vLog.Topics[0] {
			case topics["RegisterItem"]:
				var registerItemEvent RegisterItemEvent
				if err := gameShopAbi.Unpack(&registerItemEvent, "RegisterItem", vLog.Data); err != nil {
					log.Printf("Failed unpack: %v\n", err)
					continue
				}

				item := Item{
					Name:  "fuga",
					ExhibitorAddress: registerItemEvent.Exhibitor.Hex(),
					Price: registerItemEvent.Price.Uint64(),
					IsSold: false,
				}

				if err := db.Create(&item).Error; err != nil {
					log.Printf("Failed register item: %v\n", err)
				}

				break
			}
		}
	}
}



// Gormを利用してDBに接続する
func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "test_db"
	option := "?charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + option

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}