package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/stock/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"net/http"
)

func FetchClothing(web *http.Client, host string, k hsk.Key) (core.Clothing, error) {
	url := fmt.Sprintf("%s/clothes/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Clothing{}, err
	}

	defer resp.Body.Close()

	result := core.Clothing{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllClothing(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/clothes/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Clothing{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

