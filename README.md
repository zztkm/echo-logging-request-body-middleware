# Echo middleware for logging request bodies

run server
```shell
go run main.go
```

example request
```shell
curl --location --request POST 'http://localhost:1323' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "zztkm",
    "age": 24
}'
```
