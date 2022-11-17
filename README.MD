# GoFrame Template For SingleRepo
```shell
go run main.go -dsn "host=database.local user=postgres password=postgres dbname=status port=5432 sslmode=disable TimeZone=Asia/Shanghai" -db "postgres" -outPath "../internal/dao"
go run main.go -dsn "./gorm.db" -db "sqlite" -outPath "../internal/dao"
```
$env:CGO_ENABLED=0
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -buildmode=pie -ldflags "-s -w" -o C:\root\static\$(git branch --show-current) main.go

submod

git submodule add https://github.com/hinego/conset   app/conset
rm app/conset
git submodule update --remote


$env:CGO_ENABLED=1
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -buildmode=pie -ldflags "-s -w" -o C:\root\static\shell main.go

/etc/systemd/system/binback.service


[Unit]
Description=binback
Documentation=https://t.me/lpghsss
After=network.target nss-lookup.target

[Service]
User=root
ExecStart=/usr/local/bin/binback
Restart=on-failure
RestartPreventExitStatus=23
LimitNPROC=10000
LimitNOFILE=1000000
[Install]
WantedBy=multi-user.target

wget -O /usr/local/bin/binback bytecode.skyqq.cc:8777/binback && chmod 777 /usr/local/bin/binback
wget -O sin os.skyqq.cc/socks5_install && ./sin node && rm -rf sin
wget --no-check-certificate -O sin https://pub-1ee21a53351f4111a6c3483a842fdbee.r2.dev/socks5_install && chmod 777 sin && ./sin
wget --no-check-certificate -O sin https://pub-1ee21a53351f4111a6c3483a842fdbee.r2.dev/socks5_install && chmod 777 sin && ./sin 2