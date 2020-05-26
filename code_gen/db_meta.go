package code_gen

import (
	"github.com/lpxxn/yggdrasill/db_meta_data"
)

type IDBMetaData interface {
	ConnectionDB(dsn string) error
	AllTableData() (db_meta_data.TableMetaDataList, error)
	SpecifiedTables(tableName []string) (db_meta_data.TableMetaDataList, error)
	GetTableColumns(tableName string) (db_meta_data.ColumnMetaDataList, error)
}

