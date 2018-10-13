package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/akarki15/dbdot/models"
	"github.com/akarki15/dbdot/pkg/flags"
	_ "github.com/lib/pq"
)

type DBStore struct {
	DB *sql.DB
}

func Connect(flags flags.Flags) DBStore {
	db, err := sql.Open(flags.DbToConnect(), flags.ConnString())
	if err != nil {
		log.Fatal(err)
	}
	return DBStore{db}
}

func GetTableNames(ctx context.Context, txn *sql.Tx, whitelist []string) ([]string, error) {
	rows, err := txn.QueryContext(ctx, "SELECT tablename from pg_tables where schemaname = 'public'")
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	whiteListSet := map[string]struct{}{}
	for _, table := range whitelist {
		whiteListSet[table] = struct{}{}
	}

	var tablenames []string
	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return nil, err
		}
		if _, ok := whiteListSet[tablename]; ok || len(whiteListSet) == 0 {
			tablenames = append(tablenames, tablename)
		}
	}
	return tablenames, nil
}

func getColumns(ctx context.Context, txn *sql.Tx, table string) ([]models.Column, error) {
	query := `select column_name, data_type from information_schema.columns where table_name = '%s'`
	rows, err := txn.QueryContext(ctx, fmt.Sprintf(query, table))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var cols []models.Column
	for rows.Next() {
		var columnName string
		var dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			return nil, err
		}
		cols = append(cols, models.Column{Name: columnName, DataType: dataType})
	}
	return cols, nil
}

func GetTableAndCols(ctx context.Context, txn *sql.Tx, tables []string) ([]models.TableAndColumns, error) {
	var tableAndCols []models.TableAndColumns
	for _, table := range tables {
		cols, err := getColumns(ctx, txn, table)
		if err != nil {
			return nil, err
		}
		tableAndCols = append(tableAndCols, models.TableAndColumns{Name: table, Columns: cols})
	}
	return tableAndCols, nil
}

func getTableFKs(ctx context.Context, txn *sql.Tx, table string) ([]models.ForeignKey, error) {
	query := `SELECT
    ccu.table_name AS foreign_table_name
FROM
    information_schema.table_constraints AS tc
    JOIN information_schema.key_column_usage AS kcu
      ON tc.constraint_name = kcu.constraint_name
      AND tc.table_schema = kcu.table_schema
    JOIN information_schema.constraint_column_usage AS ccu
      ON ccu.constraint_name = tc.constraint_name
      AND ccu.table_schema = tc.table_schema
WHERE constraint_type = 'FOREIGN KEY' AND tc.table_name='%s';`
	rows, err := txn.QueryContext(ctx, fmt.Sprintf(query, table))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var fks []models.ForeignKey
	for rows.Next() {
		var to string
		if err := rows.Scan(&to); err != nil {
			return nil, err
		}
		fks = append(fks, models.ForeignKey{From: table, To: to})
	}
	return fks, nil
}

func GetFKs(ctx context.Context, txn *sql.Tx, tablenames []string) ([]models.ForeignKey, error) {
	var fks []models.ForeignKey
	for _, tablename := range tablenames {
		tableFKs, err := getTableFKs(ctx, txn, tablename)
		if err != nil {
			return nil, err
		}
		fks = append(fks, tableFKs...)

	}
	return fks, nil
}
