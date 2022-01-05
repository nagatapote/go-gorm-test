# go-sqlboiler-test

## 環境
言語：go 1.17

DB：postgresQL(docker)

フレームワーク：gin

## 環境構築
`$ make setup`

`$ make start`

## sqlboiler
`$ go get -u -t github.com/volatiletech/sqlboiler `

`$ sqlboiler --version`

`$ go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql`

`$ sqlboiler psql`

## mod
`$ go mod tidy`

## docker db 確認方法
`$ docker exec -it sample_database psql -U root`
