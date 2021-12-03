package impl

import (
	"strings"

	"github.com/textileio/go-tableland/internal/tableland"
	"github.com/textileio/go-tableland/pkg/sqlstore"
	"github.com/textileio/go-tableland/pkg/tableregistry"
)

// TablelandMesa is the main implementation of Tableland spec
type TablelandMesa struct {
	Store    sqlstore.SQLStore
	Registry tableregistry.TableRegistry
}

func (t *TablelandMesa) CreateTable(args tableland.SQLArgs) (tableland.Response, error) {
	if strings.Contains(strings.ToLower(args.Statement), "create") {
		err := t.Store.Write(args.Statement)
		if err != nil {
			return tableland.Response{Message: err.Error()}, err
		}
		return tableland.Response{Message: "Table created"}, nil
	}

	return tableland.Response{Message: "Invalid command"}, nil
}

func (t *TablelandMesa) UpdateTable(args tableland.SQLArgs) (tableland.Response, error) {
	// this is not going to be implemented
	return tableland.Response{Message: "Table updated"}, nil
}

func (t *TablelandMesa) RunSQL(args tableland.SQLArgs) (tableland.Response, error) {
	if strings.Contains(strings.ToLower(args.Statement), "insert") || strings.Contains(strings.ToLower(args.Statement), "update") {
		return t.runInsertOrUpdate(args)
	}

	if strings.Contains(strings.ToLower(args.Statement), "select") {
		return t.runSelect(args)
	}

	return tableland.Response{Message: "Invalid command"}, nil
}

func (t *TablelandMesa) runInsertOrUpdate(args tableland.SQLArgs) (tableland.Response, error) {
	err := t.Store.Write(args.Statement)
	if err != nil {
		return tableland.Response{Message: err.Error()}, err
	}
	return tableland.Response{Message: "Command executed"}, nil
}

func (t *TablelandMesa) runSelect(args tableland.SQLArgs) (tableland.Response, error) {
	data, err := t.Store.Read(args.Statement)
	if err != nil {
		return tableland.Response{Message: err.Error()}, err
	}

	return tableland.Response{Message: "Select executed", Data: data}, nil
}
