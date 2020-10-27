package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"net/http"
)

func FetchCar(web *http.Client, host string, k hsk.Key) (core.Car, error) {
	url := fmt.Sprintf("%s/cars/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Car{}, err
	}

	defer resp.Body.Close()

	result := core.Car{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllCars(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/cars/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Car{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
