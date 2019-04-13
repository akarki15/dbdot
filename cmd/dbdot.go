package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/akarki15/dbdot/pkg/db"
	"github.com/akarki15/dbdot/pkg/draw"
	"github.com/akarki15/dbdot/pkg/flags"
	"github.com/emicklei/dot"
)

func main() {
	ctx := context.Background()
	f := flags.Parse()
	ds := db.Connect(f)
	txn, err := ds.DB.Begin()
	if err != nil {
		HandleFatal(err)

	}
	graph, err := drawTables(ctx, txn, f)
	if err != nil {
		HandleFatal(err)
	}
	fmt.Println(graph.String())
}

func HandleFatal(err error) {
	log.Fatal(err)
}

func drawTables(ctx context.Context, txn *sql.Tx, f flags.Flags) (*dot.Graph, error) {
	tablenames, err := db.GetTableNames(ctx, txn, f.Schema, f.WhiteList)
	if err != nil {
		return nil, err
	}
	fks, err := db.GetFKs(ctx, txn, tablenames)
	if err != nil {
		return nil, err
	}
	tableAndCols, err := db.GetTableAndCols(ctx, txn, tablenames)
	if err != nil {
		return nil, err
	}
	graph, err := draw.ERD(tableAndCols, fks)
	if err != nil {
		return nil, err
	}
	return &graph, nil
}
