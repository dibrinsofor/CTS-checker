package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"price_checker/model"
)

func FetchPrice(fiat string, token string) (string, error) {
	// URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + fiat + "&convert=" + token
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=NGN&convert=NGNT"

	response, err := http.Get(URL)
	if err != nil {
		panic("problem ah gwan - not able to connect to URL")
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	var curr_response []model.CurrencyResponse

	// wtf dey gwan for here. i think i barb but i no sure. [ask griff]
	// also try this without that if statement
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&curr_response); err != nil {
		// panic("bro, yawa dat - unable to convert json to GO readable lang ")
		panic(err)
	}

	// panic: json: cannot unmarshal array into Go value of type model.CurrencyResponse

	return curr_response[0].CLIresponse(), nil

}
