package mysql

import "testing"

func TestAllTableData(t *testing.T) {
	gen := &Gen{}
	if err := gen.ConnectionDB("root:123456@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8&loc=Asia%2FShanghai"); err != nil {
		t.Error(err)
	}
	if gen.db == nil {
		t.Errorf("db is null")
	}
	AllTableData(t, gen)
	t.Log("--------")
	SpecifiedTables(t, gen)
}

func AllTableData(t *testing.T, gen *Gen) {
	rev, err := gen.AllTableData()
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range rev {
		t.Logf("table: %s", table.Name)
		for _, column := range table.Columns {
			t.Logf("column: %s, %s, %s, isUnsigned: %t", column.Name, column.DBType, column.GoType, column.IsUnsigned)
		}
	}
}

func SpecifiedTables(t *testing.T, gen *Gen) {
	rev, err := gen.SpecifiedTables([]string{"user", "employee"})
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range rev {
		t.Logf("table: %s", table.Name)
		for _, column := range table.Columns {
			t.Logf("column: %s, %s, %s, isUnsigned: %t", column.Name, column.DBType, column.GoType, column.IsUnsigned)
		}
	}
}
