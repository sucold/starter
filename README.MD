# GoFrame Template For SingleRepo
```shell
go run main.go -dsn "host=database.local user=postgres password=postgres dbname=status port=5432 sslmode=disable TimeZone=Asia/Shanghai" -db "postgres" -outPath "../internal/dao"
go run ./dev gen -dsn "./gorm.db" -db "sqlite" -outPath "./app/dao"
go run ./dev gen -db "sqlite" -outPath "./app/dao"
go run ./dev sync
go run ./dev gcc-
git submodule add https://github.com/hinego/conset/  app/conset
rm app/conset
git submodule update --remote

112
```




