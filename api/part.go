package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"net/http"
)

func FetchPart(web *http.Client, host string, k hsk.Key) (core.Part, error) {
	url := fmt.Sprintf("%s/parts/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Part{}, err
	}

	defer resp.Body.Close()

	result := core.Part{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllParts(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/parts/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Part{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
