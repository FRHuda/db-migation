package migrate

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source"
)

func DoMigrate(dbDSN string, sourceDriverName string, sourceDriver source.Driver, specifyVersion *int) (int, bool) {
	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Fatalln("error connection db: ", err)
	}
	defer db.Close()

	// Setup the database driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error when creating postgres instance: ", err)
	}

	m, err := migrate.NewWithInstance(
		sourceDriverName, sourceDriver,
		"postgres", driver)

	if err != nil {
		log.Fatal("error when creating data instance: ", err)
	}

	if err != nil {
		log.Fatal("error when creating database instance: ", err)
	}

	fmt.Println("specifi version = ", specifyVersion)
	if specifyVersion != nil {
		fmt.Println("masuk kah??")
		res := migrateToVersion(m, *specifyVersion)
		v, _, err := m.Version()
		if nil != err {
			log.Fatal(err)
		}
		return int(v), res
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error when migrate up: ", err)
	}

	v, _, err := m.Version()
	if nil != err {
		log.Fatal(err)
	}

	return int(v), true
}

func migrateToVersion(m *migrate.Migrate, version int) bool {
	dbVersion, dirty, err := m.Version()
	if int(dbVersion) >= version {
		return false
	}

	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal("Error when check version", err)
	}

	if dirty {
		log.Fatal("Migration is dirty")
	}

	if err := m.Migrate(uint(version)); err != nil && err != migrate.ErrNoChange {
		fmt.Println("error di sini?? = ", version)
		log.Fatal("error when migrate up: ", err)
	}

	return true
}
