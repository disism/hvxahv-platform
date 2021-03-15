package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	httpsig "hvxahv/pkg/activitypub"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func NewActivity(c *gin.Context) {
	url := "https://mas.to/inbox"
	method := "POST"

	idr := strconv.Itoa(rand.Int())


	obj := map[string]string {
		"id": "https://mas.to/76c94894-fa77-42be-ab8d-e35779c9cb63",
		"type": "Follow",
		"actor": "https://mas.to/users/hvturingga",
		"object": "https://47d8ab2e4028.ngrok.io/actor",
	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": fmt.Sprintf("https://%s/%s", address, idr),
		"type": "Accept",
		"actor": fmt.Sprintf("https://%s/actor", address),
		"object": obj,
	}
	byterData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}


	payload := bytes.NewBuffer(byterData)
	client := &http.Client {
	}
	fmt.Println(payload)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}


	date := time.Now().UTC().Format(http.TimeFormat)

	req.Header.Add("Host", "mas.to")
	req.Header.Add("Date", date)

	//req.Header.Add("Signature", header)

	block := httpsig.PrivateKey{
		Key: GetKey(),
	}
	httpsig.SignRequest(fmt.Sprintf("https://%s/actor", address), block, req, byterData)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	switch res.StatusCode {
	case 200:
	case 201:
	case 202:
	default:
		fmt.Errorf("https post status: %d", res.StatusCode)
	}
	log.Printf("successful post: %s %d", url, res.StatusCode)
	log.Println(req)
}

