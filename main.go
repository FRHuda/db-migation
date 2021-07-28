package main

import (
	"flag"
	"log"

	"db-migration/config"
	"db-migration/pkg/migrate"
)

func main() {
	flag.Parse()

	var c config.Config
	if err := config.Process(&c); err != nil {
		log.Println(err)
	}

	scheme := config.GetScheme("post", "local", "yaml", "./config.yaml")

	sourceDriver, err := migrate.CreateRiceBoxSourceDriver()
	if err != nil {
		log.Println(err)
	}

	if sourceDriver == nil {
		log.Println("cannot running database migration, cause the source driver is not supplied")
	} else {
		if scheme.Migration != 0 {
			log.Println("running postgres migrations")
			v, cond := migrate.DoMigrate(c.DbMigration.DSNMigrationLocal, c.DbMigration.SourceMigrationName, sourceDriver, &scheme.Migration)
			log.Printf("Migration done, do migration = %v, version = %v\n", cond, v)
		}
	}
}
