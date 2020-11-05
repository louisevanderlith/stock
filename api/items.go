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

func FetchCategory(web *http.Client, host string, k hsk.Key) (core.Category, error) {
	url := fmt.Sprintf("%s/info/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Category{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Category{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Category{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllCategories(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/info/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Category{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchStockItem(web *http.Client, host string, category string, k hsk.Key) (core.StockItem, error) {
	url := fmt.Sprintf("%s/%s/%s", host, category, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.StockItem{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.StockItem{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.StockItem{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchCategoryItems(web *http.Client, host string, category, pagesize string) (records.Page, error) {
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

	result := records.NewResultPage(core.StockItem{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
