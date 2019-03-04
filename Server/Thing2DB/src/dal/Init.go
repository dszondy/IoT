package dal

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
)

var initializer string = ``

type ConnectrionStruct struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Sslmode string `json:"sslmode"`
	Dbname string `json:"dbname"`
}
var connInfo ConnectrionStruct = ConnectrionStruct{}

func getConnectionSTR() string {return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",  connInfo.Host, connInfo.Port, connInfo.User,connInfo.Password, connInfo.Sslmode)}
func getDBSTR()string{return getConnectionSTR() + fmt.Sprintf(" dbname=%s", connInfo.Dbname)}

func catch()  {
	recover()
}
func initDB(){
	file, err := os.Open("config/config.json") // just pass the file name
	if err != nil {
		log.Printf("Error: ", err.Error())
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&connInfo); err != nil {
		log.Printf("Invalid configfile")
	}

	log.Printf("Initializing the database")
	db, err := sql.Open("postgres",getConnectionSTR() )
	if err != nil {
		log.Printf("Error while executing init script: %s",err.Error())
	}
	_, err = db.Exec("CREATE DATABASE "+connInfo.Dbname)
	if err != nil {
		log.Printf("Error while executing init script: %s",err.Error())
	}
	db.Close()

	initFile, err  := ioutil.ReadFile("config/init.sql")
	if err != nil {
		log.Printf("Error: ", err.Error())
	}
	initializer = string(initFile)

	db, err = sql.Open("postgres", getDBSTR())
	if err != nil {
		log.Printf("Error while executing init script: %s",err.Error())
	}
	_, err = db.Exec(initializer)
	if err != nil {
		log.Printf("Error while executing init script: %s",err.Error())
	}
	defer db.Close()
}

