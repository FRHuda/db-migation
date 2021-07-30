package main

import (
	"flag"
	"fmt"
	"log"

	"db-migration/config"
	"db-migration/config/model"
	"db-migration/pkg/migrate"
)

func main() {
	flag.Parse()

	var c config.Config
	if err := config.Process(&c); err != nil {
		log.Println(err)
	}

	sourceDriver, err := migrate.CreateRiceBoxSourceDriver()
	if err != nil {
		log.Println(err)
	}

	for _, env := range model.Environtments {
		scheme := config.GetScheme("testdb", env, "yaml", "./config.yaml")

		if sourceDriver == nil {
			log.Println("cannot running database migration, cause the source driver is not supplied")
		} else {
			if scheme != nil && scheme.Migration != 0 {
				log.Printf("running postgres migrations on %v environment\n", env)
				for _, char := range c.DbMigration.DSNMigrationLocal {
					fmt.Println(char)
				}
				// switch env {
				// case model.EnvLocal:
				// 	v, cond := migrate.DoMigrate(c.DbMigration.DSNMigrationLocal, c.DbMigration.SourceMigrationName, sourceDriver, &scheme.Migration)
				// 	log.Printf("Migration done on %v environment, do migration = %v, version = %v\n", env, cond, v)
				// case model.EnvStaging:
				// 	v, cond := migrate.DoMigrate(c.DbMigration.DSNMigrationStaging, c.DbMigration.SourceMigrationName, sourceDriver, &scheme.Migration)
				// 	log.Printf("Migration done on %v environment, do migration = %v, version = %v\n", env, cond, v)
				// case model.EnvProd:
				// 	v, cond := migrate.DoMigrate(c.DbMigration.DSNMigrationProduction, c.DbMigration.SourceMigrationName, sourceDriver, &scheme.Migration)
				// 	log.Printf("Migration done on %v environment, do migration = %v, version = %v\n", env, cond, v)
				// }

			}
		}
	}
}
