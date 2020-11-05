package core

import (
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type StockContext interface {
	GetClientItems(page, size int, clientid string) (records.Page, error)
	GetOwnerItems(page, size int, owner hsk.Key) (records.Page, error)
	GetStock(category string, itemKey hsk.Key) (StockItem, error)
	FindStock(page, size int, category string) (records.Page, error)
	FindStockCategory(page, size int, categoryKey hsk.Key) (records.Page, error)
	CreateStock(category string, obj StockItem) (hsk.Key, error)
	UpdateStock(category string, obj StockItem) error
	ListCategories() (records.Page, error)
	GetCategory(k hsk.Key) (hsk.Record, error)
	CreateCategory(obj Category) (hsk.Key, error)
	UpdateCategory(k hsk.Key, obj Category) error
}

type context struct {
	Categories husk.Table
}

func (c context) GetStock(category string, itemKey hsk.Key) (StockItem, error) {
	rec, err := c.Categories.FindFirst(byName(category))

	if err != nil {
		return StockItem{}, err
	}

	cat := rec.GetValue().(Category)
	itm, _, err := cat.GetItem(itemKey)
	return itm, err
}

func (c context) FindStock(page, size int, category string) (records.Page, error) {
	rec, err := c.Categories.FindFirst(byName(category))

	if err != nil {
		return nil, err
	}

	result := records.NewRecordPage(page, size)
	val := rec.GetValue().(Category)
	skipCount := (page - 1) * size

	for _, itm := range val.Items {
		if skipCount != 0 {
			skipCount--
			continue
		}

		if !result.Add(records.MakeRecord(itm.ItemKey, itm)) {
			break
		}
	}

	return result, nil
}

func (c context) FindStockCategory(page, size int, categoryKey hsk.Key) (records.Page, error) {
	rec, err := c.Categories.FindByKey(categoryKey)

	if err != nil {
		return nil, err
	}

	result := records.NewRecordPage(page, size)
	val := rec.GetValue().(Category)
	skipCount := (page - 1) * size

	for _, itm := range val.Items {
		if skipCount != 0 {
			skipCount--
			continue
		}
		if !result.Add(records.MakeRecord(itm.ItemKey, itm)) {
			break
		}
	}

	return result, nil
}

func (c context) CreateStock(category string, obj StockItem) (hsk.Key, error) {
	rec, err := c.Categories.FindFirst(byName(category))

	if err != nil {
		return nil, err
	}

	val := rec.GetValue().(Category)
	_, _, err = val.GetItem(obj.ItemKey)

	if err == nil {
		return nil, errors.New("duplicate record")
	}

	err = obj.Valid()
	if err != nil {
		return nil, err
	}

	val.Items = append(val.Items, obj)

	err = c.Categories.Update(rec.GetKey(), val)

	return rec.GetKey(), err
}

func (c context) UpdateStock(category string, obj StockItem) error {
	rec, err := c.Categories.FindFirst(byName(category))

	if err != nil {
		return err
	}

	val := rec.GetValue().(Category)
	_, idx, err := val.GetItem(obj.ItemKey)

	if err != nil {
		return err
	}

	err = obj.Valid()
	if err != nil {
		return err
	}

	val.Items[idx] = obj

	return c.Categories.Update(rec.GetKey(), val)
}

func (c context) ListCategories() (records.Page, error) {
	return c.Categories.Find(1, 10, op.Everything())
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

func (c context) GetClientItems(page, size int, clientid string) (records.Page, error) {
	return c.Categories.Find(page, size, byClient(clientid))
}

func (c context) GetOwnerItems(page, size int, owner hsk.Key) (records.Page, error) {
	return c.Categories.Find(page, size, byOwner(owner))
}

var ctx context

func Context() StockContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Categories: husk.NewTable(Category{}),
	}
}

func Shutdown() {
	ctx.Categories.Save()
}
