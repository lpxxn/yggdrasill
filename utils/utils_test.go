package utils

import "testing"

func TestStrToJsonCamel(t *testing.T) {
	str1 := "hello_WorLd_hey"
	if CamelizeStr(str1, false) != "helloWorldHey" {
		t.Error("string to json camel error")
	}
	if CamelizeStr(str1, true) != "HelloWorldHey" {
		t.Error("string to json camel error")
	}
	str1 = "iD"
	if CamelizeStr(str1, false) != "ID" {
		t.Error("string to json camel error")
	}
	str1 = "my_TeSt_id"
	if CamelizeStr(str1, false) != "myTestID" {
		t.Error("string to json camel error")
	}
	if CamelizeStr(str1, true) != "MyTestID" {
		t.Error("string to json camel error")
	}
	str1 = "a"
	if CamelizeStr(str1, false) != "a" {
		t.Error("string to json camel error")
	}
	if CamelizeStr(str1, true) != "A" {
		t.Error("string to json camel error")
	}
}

func TestGetDbNameFromDSN(t *testing.T) {
	str1 := "postgres://:@127.0.0.1:5432/test?sslmode=disable"
	dbName, err := GetDbNameFromDSN(str1)
	if err != nil {
		t.Fatal(err)
	}
	if dbName != "test" {
		t.Fatal("error db name")
	}
	str1 = "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8&loc=Asia%2FShanghai"
	dbName, err = GetDbNameFromDSN(str1)
	if err != nil {
		t.Fatal(err)
	}
	if dbName != "test" {
		t.Fatal("error db name")
	}

	str1 = "host=127.0.0.1 dbname=test sslmode=disable Timezone=Asia/Shanghai"
	dbName, err = GetDbNameFromDSN(str1)
	if err != nil {
		t.Fatal(err)
	}
	if dbName != "test" {
		t.Fatal("error db name")
	}
}
