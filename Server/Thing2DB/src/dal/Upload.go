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
	defer db.Close()
	_, err = db.Exec("INSERT INTO raw_data (device_id, lot_id, lot_timestamp, is_clear) VALUES ($1, $2, $3, $4)", device_id, lot_id, lot_timestamp, is_clear)
	if err != nil {
		log.Printf("Error while executing upload: %s", err.Error())
	}
}

func ReadPlaces()(ids []int, states []bool){
	db, err := sql.Open("postgres", getDBSTR())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	ids = make([]int, 0);
	states = make([]bool, 0);
	rows, err:= db.Query("SELECT id, state FROM lot_states;")

	for rows.Next() {
		var id int
		var is_clear bool
		err = rows.Scan(&id, &is_clear)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
		states = append(states, is_clear)
	}
	return
}
