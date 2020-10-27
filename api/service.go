package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"net/http"
)

func FetchService(web *http.Client, host string, k hsk.Key) (core.Service, error) {
	url := fmt.Sprintf("%s/services/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Service{}, err
	}

	defer resp.Body.Close()

	result := core.Service{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllServices(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/services/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Service{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
