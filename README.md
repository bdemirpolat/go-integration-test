## Run
```
go run main.go
```

## Create user
```
curl --location --request POST 'localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"burak"
}'
```