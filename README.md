[中文简介](README.md)    
[English](README_US.md)

yggdrasill 把数据库的表转换成`go`语言的`struct`，支持 `PostgreSQL`, `MySQL`    
eg:    
![generat model](/gen.gif)
## 安装 
安装到`GOPATH`的 `bin`目录.
```
GO111MODULE=on go get -u github.com/lpxxn/yggdrasill/cmd/yggdrasill
```
### 帮助
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


## 命令

### MySql
`-target`为 `mysql`
默认生成数据库内的所有表
```
yggdrasill -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" 
```
使用 `-table_names` 指定想生成的表    
```
yggdrasill -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

### PostgreSql
`-target` 为 `postgresql`或者`pg`
默认生成数据库内的所有表
```
yggdrasill -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable"
```
使用 `-table_names` 指定想生成的表    
```
yggdrasill -target=pg -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

自定义 template
使用 `-template_path` 自定义模板 
```
yggdrasill  -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -package_name=db_model -template_path=../../test/test_template.tml 
```


