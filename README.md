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
`rice` embeds the `./pkg/migrate/files` to a Go file (`./pkg/migrate/rice-box.go`).

More about rice: https://github.com/GeertJohan/go.rice.

### Run

#### Environment
```
DSN_LOCAL=
DSN_STAGING=
DSN_PRODUCTION=
SOURCE_MIGRATION_NAME=go.rice
```


running the app using command:
```
go run main.go
```