package main

import (
	"encoding/json"
	"log"
)

type Request struct {
	WalletID     string `json:"walletID"`
	ActionType   string `json:"actionType"`
	SourceName   string `json:"sourceName"`
	SourceNumber string `json:"sourceNumber"`
	ResourceId   string `json:"resourceId"`
	Mobile       string `json:"mobile"`
}

type SubRequest struct {
	WalletID   string `json:"walletID"`
	ResourceId string `json:"resourceId"`
	Mobile     string `json:"mobile"`
}

func main() {
	src := Request{"1234", "Action1", "Name1", "Number1", "Resource1", "0398881726"}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatal("Marshal data failed")
	}
	var subRequest SubRequest
	err = json.Unmarshal(data, &subRequest)
	if err != nil {
		log.Fatal("Unmarshal sub structure failed")
	}
	subData, err := json.Marshal(subRequest)
	if err != nil {
		log.Fatal("Marshal sub data failed")
	}
	log.Println("Subdata:", string(subData))
}
