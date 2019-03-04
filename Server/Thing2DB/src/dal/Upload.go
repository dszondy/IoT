package dal

import "database/sql"
import (
	_ "github.com/lib/pq"
	"log"
	"time"
)

var closeDB = make(chan int)
func RunDB(){
	initDB()
	defer func() { if recover() != nil {log.Printf( "Error while initializing db")}}()
}

func UploadOne(device_id string, lot_id string, lot_timestamp time.Time, is_clear bool){
	db, err := sql.Open("postgres", getDBSTR())
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("INSERT INTO raw_data (device_id, lot_id, lot_timestamp, is_clear) VALUES ($1, $2, $3, $4)", device_id, lot_id, lot_timestamp, is_clear)
	if err != nil {
		log.Printf("Error while executing upload: %s", err.Error())
	}
	defer db.Close()
}
