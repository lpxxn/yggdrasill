yggdrasill can turn the table into a go struct，support `PostgreSQL`, `MySQL`    
eg:    
![generat model](/gen.gif)
## install 

```
GO111MODULE=on go get -u github.com/lpxxn/yggdrasill/cmd/yggdrasill
```
### help
```
yggdrasill -help 
```
```
Usage of yggdrasill:
  -dir string
        Destination dir for files generated. (default "./tmp")
  -dsn string
        dsn (default "postgresql")
  -package_name string
        package name default model. (default "model")
  -table_names value
        if it is empty, will generate all tables in database
  -target string
        mysql postgresql[pg] (default "postgresql")
  -template_path string
        custom template file path

```


## command

### MySql
`-target` is `mysql`
```
yggdrasill -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" 
```

use `-table_names` specify table names
```
yggdrasill -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

### PostgreSql
`-target` is `postgresql` or `pg`
```
yggdrasill -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable"
```
use `-table_names` specify table names
```
yggdrasill -target=pg -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

custom template
```
yggdrasill  -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -package_name=db_model -template_path=../../test/test_template.tml 
```


