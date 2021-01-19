package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"io/ioutil"
	"net/http"
)

func FetchStockItem(web *http.Client, host string, category string, k hsk.Key) (core.Product, error) {
	url := fmt.Sprintf("%s/%s/%s", host, category, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Product{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Product{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Product{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchCategoryItems(web *http.Client, host, category, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/%s/%s", host, category, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Product{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
