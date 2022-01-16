# go-gorm-test

## 環境
言語：go 1.17

DB：postgresQL(docker)


## 環境構築
`$ make setup`

`$ make start`

## mod
`$ go mod tidy`

## docker db 確認方法
`$ docker exec -it sample_database psql -U root`
