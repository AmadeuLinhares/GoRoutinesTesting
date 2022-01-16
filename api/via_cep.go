package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

var responseApiViaCep ResponseViaCep

func SearchCepByViaCep(cep string) ResponseViaCep {
	urlToCall := "https://viacep.com.br/ws/" + cep + "/json/"
	response, err := http.Get(urlToCall)
	if err != nil {
		log.Fatal(err)
	}

	bytes, bytesError := ioutil.ReadAll(response.Body)

	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if bytesError != nil {
		log.Fatal(bytesError)
	}

	Unmarshalerror := json.Unmarshal(bytes, &responseApiViaCep)

	if Unmarshalerror != nil {
		fmt.Println("error")
		log.Fatal(Unmarshalerror)
	}

	fmt.Printf("%+v", responseApiViaCep)
	return responseApiViaCep
}
