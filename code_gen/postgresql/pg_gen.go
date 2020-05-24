package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/lpxxn/yggdrasill/db_meta_data"
	"github.com/lpxxn/yggdrasill/utils"
)

var (
	tableNamesSql          = `select table_name from information_schema.tables where table_schema = 'public';`
	specifiedTableNamesSql = `select table_name from information_schema.tables where table_schema = 'public' and table_name =any($1);`

	tableColumnsSql = `select column_name, is_nullable, data_type, false
from information_schema.columns where table_schema = 'public' and table_name = $1 order by ordinal_position;`
)

type PGGen struct {
	db     *sql.DB
	dbName string
}

func (p *PGGen) ConnectionDB(dsn string) error {
	dbName, err := utils.GetDbNameFromDSN(dsn)
	if err != nil {
		return err
	}
	p.dbName = dbName
	fmt.Println("Postgres Connecting dsn : " + dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *PGGen) AllTableData() (db_meta_data.TableMetaDataList, error) {
	rows, err := p.db.Query(tableNamesSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.TableMetaDataList{}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableColumnsInfo, err := p.GetTableColumns(tableName)
		if err != nil {
			return nil, err
		}
		rev = append(rev, &db_meta_data.TableMetaData{Name: tableName, Columns: tableColumnsInfo})
	}

	return rev, rows.Err()
}
func (p *PGGen) SpecifiedTables(tableNameList []string) (db_meta_data.TableMetaDataList, error) {
	if len(tableNameList) == 0 {
		return nil, errors.New("tableNameList is empty")
	}
	rows, err := p.db.Query(specifiedTableNamesSql, pq.Array(tableNameList))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.TableMetaDataList{}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableColumnsInfo, err := p.GetTableColumns(tableName)
		if err != nil {
			return nil, err
		}
		rev = append(rev, &db_meta_data.TableMetaData{Name: tableName, Columns: tableColumnsInfo})
	}

	return rev, rows.Err()
}

func (p *PGGen) GetTableColumns(tableName string) (db_meta_data.ColumnMetaDataList, error) {
	rows, err := p.db.Query(tableColumnsSql, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.ColumnMetaDataList{}
	for rows.Next() {
		var name, isNullable, dataType string
		var isUnsigned bool
		if err := rows.Scan(&name, &isNullable, &dataType, &isUnsigned); err != nil {
			return nil, err
		}
		rev = append(rev, db_meta_data.NewColumnMetaData(name,
			strings.ToLower(isNullable) == "yes", dataType, isUnsigned, tableName))
	}
	return rev, rows.Err()
}
