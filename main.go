package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

// Os nomes das Strucs ex Cep, Logradouro, tem que começar com letra maiuscula
// Em Go quando temos nomes com letras maiusculas, o mesmo pode ser exportado
var waitGroup sync.WaitGroup

type ResponseViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type ResponseCepApi struct {
	Cep         string `json:"code"`
	Logradouro  string `json:"address"`
	Complemento string `json:"district"`
	Cidade      string `json:"city"`
	Estado      string `json:"state"`
}

var (
	responseApiViaCep ResponseViaCep
	responseCepApi    ResponseCepApi
)

func SearchCepByViaCep(cep string, c chan<- ResponseViaCep) {
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

	// fmt.Printf("%+v \n", responseApiViaCep)

	waitGroup.Done()
	c <- responseApiViaCep
	close(c)
}

func SearchCepByCepApi(cep string, s chan<- ResponseCepApi) {
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

	// fmt.Printf("%+v \n", responseCepApi)
	waitGroup.Done()
	s <- responseCepApi
	close(s)
}

func main() {
	fmt.Println("Programming Initialization")
	fmt.Println("Numero de processadores", runtime.NumCPU())
	fmt.Println("Numero de gorotines antes de executar", runtime.NumGoroutine())
	channel1 := make(chan ResponseCepApi)
	channel2 := make(chan ResponseViaCep)
	waitGroup.Add(2)

	go SearchCepByCepApi("71215600", channel1)
	go SearchCepByViaCep("01001000", channel2)

	select {
	case returnValue := <-channel1:
		fmt.Printf("Valor retornado primeiro foi a api CepApi: %v \n\n", returnValue)
	case returnValue := <-channel2:
		fmt.Printf("Valor retornado primeiro foi a api ViaCep: %v \n\n", returnValue)
	}

	// fmt.Printf("Informacao tirada do canal Cep Api foi: %v\n\n", <-channel1)
	// fmt.Printf("A informacao retirada do calan Via Cep foi: %v \n\n", <-channel2)
	// fmt.Println("Numero de gorotines depois de executar", runtime.NumGoroutine())
	waitGroup.Wait()
}

// ResponseViaCep
// ResponseCepApi
