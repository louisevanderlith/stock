package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/core"
	"io/ioutil"
	"net/http"
)

func FetchProduct(web *http.Client, host string, k hsk.Key) (core.Product, error) {
	url := fmt.Sprintf("%s/products/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Product{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Product{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Product{ImageKey: keys.CrazyKey(), CategoryKey: keys.CrazyKey()}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllProducts(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/products/%s", host, pagesize)
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

func FetchSearchProducts(web *http.Client, host, pagesize string, in core.Product) (records.Page, error) {
	bits, err := json.Marshal(in)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/products/%s/%s", host, pagesize, base64.URLEncoding.EncodeToString(bits))
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
