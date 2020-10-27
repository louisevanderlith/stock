package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"net/http"
)

func FetchProperty(web *http.Client, host string, k hsk.Key) (core.Property, error) {
	url := fmt.Sprintf("%s/properties/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Property{}, err
	}

	defer resp.Body.Close()

	result := core.Property{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllProperties(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/properties/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Property{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
