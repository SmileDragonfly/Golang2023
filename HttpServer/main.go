package main

import (
	"HttpServer/SCB"
	"HttpServer/Utils"
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

func RequestOTPAccHandler(respw http.ResponseWriter, r *http.Request) {
	log.Info(r.Method)
	reqHeaderBytes, err := json.Marshal(r.Header)
	if err != nil {
		log.Error("Marshal header error")
		return
	}
	log.Info("Header:", string(reqHeaderBytes))
	reqBodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Read request body error")
		return
	}
	var request SCB.RequestBody
	err = json.Unmarshal(reqBodyBytes, &request)
	if err != nil {
		log.Error("Unmarshal resquest body error")
		return
	}
	log.Info("Body:", string(reqBodyBytes))

	// Create response
	var response SCB.ResponseBody
	var responseData SCB.Response_ERequestOTPAcc
	var authInfo SCB.Response_ERequestOTPAcc_AuthInfo
	response.RequestID = request.RequestID
	response.RequestDateTime = request.RequestDateTime
	response.FunctionName = request.FunctionName
	response.RespCode = "00"
	response.Description = "This is the test format structure"

	responseData.IsRequiredOTP = true
	responseData.AuthInfo = authInfo
	resData, err := json.Marshal(responseData)
	if err != nil {
		log.Error("Marshal response body data error")
		return
	}
	resDataEncrypted, err := Utils.EncryptTripleDESB64(string(resData), "TestSecretKey32Character")
	if err != nil {
		log.Error("Encrypt response body data error")
		return
	}
	response.Data = resDataEncrypted
	resBody, err := json.Marshal(response)
	if err != nil {
		log.Error("Marshal response body error")
		return
	}
	_, err = respw.Write(resBody)
	if err != nil {
		log.Error("Write response body error")
		return
	}
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
	http.HandleFunc("/RequestOTPAcc", RequestOTPAccHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Error(err.Error())
	}
}
