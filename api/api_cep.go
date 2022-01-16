package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseCepApi struct {
	Cep         string `json:"code"`
	Logradouro  string `json:"address"`
	Complemento string `json:"district"`
	Cidade      string `json:"city"`
	Estado      string `json:"state"`
}

var responseCepApi ResponseCepApi

func SearchCepByCepApi(cep string) {
	urlToCall := "https://ws.apicep.com/cep/" + cep + ".json/"

	resp, err := http.Get(urlToCall)
	if err != nil {
		log.Fatal(err)
	}

	bytes, bytesError := ioutil.ReadAll(resp.Body)

	defer func() {
		e := resp.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if bytesError != nil {
		log.Fatal(bytesError)
	}

	Unmarshalerror := json.Unmarshal(bytes, &responseCepApi)

	if Unmarshalerror != nil {
		fmt.Println("error")
		log.Fatal(Unmarshalerror)
	}

	fmt.Printf("%+v", responseCepApi)
}
