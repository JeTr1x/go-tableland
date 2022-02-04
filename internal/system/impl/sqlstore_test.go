package impl

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/textileio/go-tableland/internal/tableland"
	parserimpl "github.com/textileio/go-tableland/pkg/parsing/impl"
	"github.com/textileio/go-tableland/pkg/sqlstore/impl"
	txnimpl "github.com/textileio/go-tableland/pkg/txn/impl"
	"github.com/textileio/go-tableland/tests"
)

func TestSystemSQLStoreService(t *testing.T) {
	url, err := tests.PostgresURL()
	require.NoError(t, err)

	ctx := context.Background()
	store, err := impl.New(ctx, url)
	require.NoError(t, err)

	// populate the system_tables with a table
	txnp, err := txnimpl.NewTxnProcessor(url)
	require.NoError(t, err)
	b, err := txnp.OpenBatch(ctx)
	require.NoError(t, err)

	parser := parserimpl.New("system_")
	id, _ := tableland.NewTableID("42")
	createStmt, err := parser.ValidateCreateTable("create table foo (bar int)")
	require.NoError(t, err)

	err = b.InsertTable(ctx, id, "0xb451cee4A42A652Fe77d373BAe66D42fd6B8D8FF", "descrp-1", createStmt)
	require.NoError(t, err)
	require.NoError(t, b.Commit(ctx))
	require.NoError(t, b.Close(ctx))

	svc := NewSystemSQLStoreService(store)
	metadata, err := svc.GetTableMetadata(ctx, id)
	require.NoError(t, err)

	require.Equal(t, fmt.Sprintf("https://tableland.com/tables/%s", id), metadata.ExternalURL)
	require.Equal(t, "https://hub.textile.io/thread/bafkqtqxkgt3moqxwa6rpvtuyigaoiavyewo67r3h7gsz4hov2kys7ha/buckets/bafzbeicpzsc423nuninuvrdsmrwurhv3g2xonnduq4gbhviyo5z4izwk5m/todo-list.png", metadata.Image) //nolint
	require.Equal(t, "date", metadata.Attributes[0].DisplayType)
	require.Equal(t, "created", metadata.Attributes[0].TraitType)

	// this is hard to test because the created_at comes from the database. just testing is not the 1970 value
	require.NotEqual(t, new(time.Time).Unix(), metadata.Attributes[0].Value)

	tables, err := svc.GetTablesByController(ctx, "0xb451cee4A42A652Fe77d373BAe66D42fd6B8D8FF")
	require.NoError(t, err)
	require.Equal(t, 1, len(tables))
	require.Equal(t, id, tables[0].ID)
	require.Equal(t, "0xb451cee4A42A652Fe77d373BAe66D42fd6B8D8FF", tables[0].Controller)
	require.Equal(t, "foo", tables[0].Name)
	require.Equal(t, "descrp-1", tables[0].Description)
	// sha256(bar:int4) = 60b0e90a94273211e4836dc11d8eebd96e8020ce3408dd112ba9c42e762fe3cc
	require.Equal(t, "60b0e90a94273211e4836dc11d8eebd96e8020ce3408dd112ba9c42e762fe3cc", tables[0].Structure)
	require.Equal(t, metadata.Attributes[0].Value, tables[0].CreatedAt.Unix())
}
