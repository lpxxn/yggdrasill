package db_meta_data

import (
	"fmt"
	"strings"

	"github.com/lpxxn/yggdrasill/utils"
)

type TableMetaData struct {
	Name    string
	Columns ColumnMetaDataList
}

type TableMetaDataList []*TableMetaData

func (t TableMetaData) Imports() []string {
	imports := map[string]string{}

	for _, column := range t.Columns {
		columnType := column.GoType
		if v, ok := customerColumnDataTypeImport[columnType]; ok {
			imports[columnType] = v
			continue
		}
		switch columnType {
		case "time.Time":
			imports["time.Time"] = "time"
		}
	}
	rev := []string{}
	for _, packageImport := range imports {
		rev = append(rev, packageImport)
	}
	return rev
}

func (t TableMetaData) ColumnsNameWithPrefixAndIgnoreColumn(col string, prefix string) string {
	rev := ""
	for _, item := range t.Columns {
		if strings.ToLower(item.Name) == col {
			continue
		}
		if len(rev) > 0 {
			rev += ", "
		}
		rev += prefix + "." + utils.CamelizeStr(item.Name, true)
	}
	return rev
}

type ColumnMetaData struct {
	Name       string
	DBType     string
	GoType     string
	IsUnsigned bool
	IsNullable bool
	TableName  string
}

type ColumnMetaDataList []*ColumnMetaData

var customerColumnDataType map[string]string
var customerColumnDataTypeImport map[string]string

func NewColumnMetaData(name string, isNullable bool, dataType string, isUnsigned bool, tableName string) *ColumnMetaData {
	columnMetaData := &ColumnMetaData{
		Name:       name,
		IsNullable: isNullable,
		DBType:     dataType,
		IsUnsigned: isUnsigned,
		TableName:  tableName,
	}
	columnMetaData.GoType = columnMetaData.getGoType()
	return columnMetaData
}

func CustomerColumnDataType(dbColumnType string, customerType string, importStr string) {
	customerColumnDataType[dbColumnType] = customerType
	customerColumnDataTypeImport[customerType] = importStr
}

func (c ColumnMetaData) getGoType() string {
	if value, ok := customerColumnDataType[c.DBType]; ok {
		return value
	}
	switch c.DBType {
	case "boolean":
		return "bool"
	case "tinyint":
		return "int8"
	case "smallint", "year":
		return "int16"
	case "integer", "mediumint", "int":
		return "int32"
	case "bigint":
		return "int64"
	case "date", "timestamp without time zone", "timestamp with time zone", "time with time zone", "time without time zone",
		"timestamp", "datetime", "time":
		return "time.Time"
	case "bytea",
		"binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "[]byte"
	case "text", "character", "character varying", "tsvector", "bit", "bit varying", "money", "json", "jsonb", "xml", "point", "interval", "line", "ARRAY",
		"char", "varchar", "tinytext", "mediumtext", "longtext":
		return "string"
	case "real":
		return "float32"
	case "numeric", "decimal", "double precision", "float", "double":
		return "float64"
	default:
		return "string"
	}
}
func (c ColumnMetaData) Tag() string {
	return fmt.Sprintf("`db:\"%s\" json:\"%s\"`", c.Name, utils.CamelizeStr(c.Name, false))
}
