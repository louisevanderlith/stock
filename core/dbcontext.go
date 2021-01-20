package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
)

type StockContext interface {
	//Categories
	GetCategory(k hsk.Key) (hsk.Record, error)
	CreateCategory(obj Category) (hsk.Key, error)
	UpdateCategory(k hsk.Key, obj Category) error
	ListCategories(page, size int) (records.Page, error)
	SearchCategories(page, size int, in Category) (records.Page, error)

	//Products
	GetProduct(productKey hsk.Key) (hsk.Record, error)
	CreateProduct(obj Product) (hsk.Key, error)
	UpdateProduct(prodKey hsk.Key, obj Product) error
	ListProducts(page, size int) (records.Page, error)
	SearchProducts(page, size int, in Product) (records.Page, error)

	//Search
	FindCategoriesByClient(page, size int, clientId string) (records.Page, error)
	FindCategoriesByOwner(page, size int, owner hsk.Key) (records.Page, error)
	FindProductsByItem(page, size int, itemKey hsk.Key) (records.Page, error)
	FindProductsByCategory(page, size int, categoryKey hsk.Key) (records.Page, error)
	FindProductsByCategoryName(page, size int, category string) (records.Page, error)
}

type context struct {
	Categories husk.Table
	Products   husk.Table
}

func (c context) ListProducts(page, size int) (records.Page, error) {
	return c.Products.Find(page, size, op.Everything())
}

func (c context) SearchCategories(page, size int, in Category) (records.Page, error) {
	return c.Products.Find(page, size, byCategorySearch(in))
}

func (c context) SearchProducts(page, size int, in Product) (records.Page, error) {
	return c.Products.Find(page, size, byProductSearch(in))
}

func (c context) GetCategory(k hsk.Key) (hsk.Record, error) {
	return c.Categories.FindByKey(k)
}

func (c context) CreateCategory(obj Category) (hsk.Key, error) {
	return c.Categories.Create(obj)
}

func (c context) UpdateCategory(k hsk.Key, obj Category) error {
	return c.Categories.Update(k, obj)
}

func (c context) GetProduct(productKey hsk.Key) (hsk.Record, error) {
	return c.Products.FindByKey(productKey)
}

func (c context) CreateProduct(obj Product) (hsk.Key, error) {
	return c.Products.Create(obj)
}

func (c context) UpdateProduct(productKey hsk.Key, obj Product) error {
	return c.Products.Update(productKey, obj)
}

func (c context) FindProductsByItem(page, size int, itemKey hsk.Key) (records.Page, error) {
	return c.Products.Find(page, size, byItemKey(itemKey))
}

func (c context) FindProductsByCategory(page, size int, categoryKey hsk.Key) (records.Page, error) {
	return c.Products.Find(page, size, byCategory(categoryKey))
}

func (c context) FindProductsByCategoryName(page, size int, category string) (records.Page, error) {
	cat, err := c.Categories.FindFirst(byName(category))

	if err != nil {
		return nil, err
	}

	return c.Products.Find(page, size, byCategory(cat.GetKey()))
}

func (c context) ListCategories(page, size int) (records.Page, error) {
	return c.Categories.Find(page, size, op.Everything())
}

func (c context) FindCategoriesByClient(page, size int, clientid string) (records.Page, error) {
	return c.Categories.Find(page, size, byClient(clientid))
}

func (c context) FindCategoriesByOwner(page, size int, owner hsk.Key) (records.Page, error) {
	return c.Categories.Find(page, size, byOwner(owner))
}

var ctx context

func Context() StockContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Categories: husk.NewTable(Category{}),
		Products:   husk.NewTable(Product{}),
	}

	seed()
}

func seed() {
	coll, err := categorySeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Categories.Seed(coll)

	if err != nil {
		panic(err)
	}
}

func categorySeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/categories.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Category
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func Shutdown() {
	ctx.Categories.Save()
}
