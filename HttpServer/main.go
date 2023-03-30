package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"io"
	"net/http"
)

func LinkHandler(respw http.ResponseWriter, r *http.Request) {
	log.Info(r.Method)
	reqHeaderBytes, err := json.Marshal(r.Header)
	if err != nil {
		log.Error("Marshal header error")
	}
	log.Info("Header:", string(reqHeaderBytes))
	regBodyBytes, err := json.Marshal(r.Body)
	if err != nil {
		log.Error("Marshal body error")
	}
	log.Info("Body:", string(regBodyBytes))
	return
}

func RequestOTP(respw http.ResponseWriter, r *http.Request) {
	log.Info(r.Method)
	reqHeaderBytes, err := json.Marshal(r.Header)
	if err != nil {
		log.Error("Marshal header error")
	}
	log.Info("Header:", string(reqHeaderBytes))
	regBodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Marshal body error")
	}
	log.Info("Body:", string(regBodyBytes))
	return
}

func RequestOTPForLink(respw http.ResponseWriter, r *http.Request) {
	log.Info(r.Method)
	reqHeaderBytes, err := json.Marshal(r.Header)
	if err != nil {
		log.Error("Marshal header error")
	}
	log.Info("Header:", string(reqHeaderBytes))
	regBodyBytes, err := json.Marshal(r.Body)
	if err != nil {
		log.Error("Marshal body error")
	}
	log.Info("Body:", string(regBodyBytes))
	return
}

func main() {
	// Init log file
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log.LoadConfiguration("./log4go.json")
	log.Info("----------Init logger----------")
	defer log.Close()
	http.HandleFunc("/Link", LinkHandler)
	http.HandleFunc("/RequestOTP", RequestOTP)
	http.HandleFunc("/RequestOTPForLink", RequestOTPForLink)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Error(err.Error())
	}
}
