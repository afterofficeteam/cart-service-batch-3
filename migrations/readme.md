### Run Migration
```
go run migration.go ./sql "host=localhost port=5432 user=postgres dbname=ao_cart_service_batch_3 password=postgres sslmode=disable" up
```

### Down Migration
```
go run migration.go ./sql "host=localhost port=5432 user=postgres dbname=ao_cart_service_batch_3 password=postgres sslmode=disable" down
```

### Create new SQL
```
go run migration.go ./sql "host=localhost port=5432 user=postgres dbname=ao_cart_service_batch_3 sslmode=disable" create add_orders_table sql
```