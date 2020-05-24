package postgresql

import "testing"

func TestAllTableData(t *testing.T) {
	gen := &PGGen{}
	if err := gen.ConnectionDB("postgres://:@127.0.0.1:5432/test?sslmode=disable"); err != nil {
		t.Error(err)
	}
	if gen.db == nil {
		t.Errorf("db is null")
	}
	AllTableData(t, gen)
	t.Log("-------")
	SpecifiedTables(t, gen)
}

func AllTableData(t *testing.T, gen *PGGen) {
	rev, err := gen.AllTableData()
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range rev {
		t.Logf("table: %s", table.Name)
		for _, column := range table.Columns {
			t.Logf("column: %s, %s, %s", column.Name, column.DBType, column.GoType)
		}
	}
}

func SpecifiedTables(t *testing.T, gen *PGGen) {
	rev, err := gen.SpecifiedTables([]string{"role", "user_role"})
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
