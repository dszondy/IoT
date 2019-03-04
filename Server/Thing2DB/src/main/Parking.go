package main

import (
	"github.com/gorilla/mux"
	"strings"
	"net/http"
	"conn"
	"strconv"
)

type espLotData struct{
	DevId   string `json:"devid"`
	LotId   string `json:"lotid"`
	IsClear bool   `json:"isclear"`
}

type state struct {
	IsClear bool   `json:"isclear"`
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Handler(errorHandler(route.HandlerFunc, route.Name)).Name(route.Name)
	}
	return router
}

func ReadBody(writer http.ResponseWriter, request *http.Request){
	var data *espLotData = &espLotData{}
	if 	tryReadRequestBody(request, data, writer)  {
		conn.RequestChannel <- conn.CreateLotState(data.DevId, data.LotId, data.IsClear)
		respondOK(writer)
	}
}


func ReadMixed(writer http.ResponseWriter, request *http.Request) {
	var state *state = &state{}
	if tryReadRequestBody(request, state, writer) {
		vars := mux.Vars(request)
		deviceId := vars["deviceId"]
		lotId := vars["lotId"]
		conn.RequestChannel <- conn.CreateLotState(deviceId, lotId, state.IsClear)
		respondOK(writer)
	}
}

func ReadURI(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		deviceId := vars["deviceId"]
		lotId := vars["lotId"]
		isClear := vars["isClear"]
		IsClear, err := strconv.ParseBool(isClear)
		if err !=nil{
			respondWithError(writer, http.StatusBadRequest, "State is not valid")
			return;
		}
		conn.RequestChannel <- conn.CreateLotState(deviceId, lotId, IsClear)
		respondOK(writer)}

var routes = []route{
	route{
		Name:        "ReceiveJson",
		Method:      strings.ToUpper("Put"),
		Pattern:     "/v1/device/json",
		HandlerFunc: ReadBody,
	},
	route{
		Name:        "ReceiveMixed",
		Method:      strings.ToUpper("Put"),
		Pattern:     "/v1/device/{deviceId}/{lotId}",
		HandlerFunc: ReadMixed,
	},
	route{
		Name:        "ReceiveStateURI",
		Method:      strings.ToUpper("Put"),
		Pattern:     "/v1/device/{deviceId}/{lotId}/{isClear}",
		HandlerFunc: ReadURI,
	},
}