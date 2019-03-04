package conn

import (
	"container/list"
	"time"
	"dal"
	"log"
)

type LotState struct {
	dev_id string
	lot_id string
	isEmpty bool
	timestamp time.Time
}
func CreateLotState(dev_id string, lot_id string,isEmpty bool)*LotState{
	return &LotState{dev_id, lot_id, isEmpty, time.Now()}
}

var RequestChannel chan *LotState
var dataBuffer list.List

func InitConnectionLayer(){
	RequestChannel = make(chan *LotState)
	go dal.RunDB()
	go acceptData()
	}

func startCommitTimer()  {
	timer := time.NewTicker(500 * time.Millisecond)
	for true {
		_=<-timer.C
		RequestChannel <-nil
	}
}

func acceptData() {
	for true {
		data := <-RequestChannel
		if data == nil{
			break
		} else {
			log.Println("Recived data on channel.")
			dal.UploadOne(data.dev_id,data.lot_id,data.timestamp, data.isEmpty)
		}
	}
}
