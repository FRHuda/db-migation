# DB Migration

## How To Use

### Create migration file
create migration file on folder `pkg/migrate/files` .
each migration has an up and down migration. 
```
1_create_table_user.up.sql
1_create_table_user.down.sql
```

### Generate rice box
generating rice box for all of migrations using command:
```
rice embed-go -i=./pkg/migrate
```

### Run
running the app using command:
```
go run main.go
```