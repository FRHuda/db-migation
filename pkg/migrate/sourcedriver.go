package migrate

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/golang-migrate/migrate/source"
)

// CreateRiceBoxSourceDriver create ricee box source driver for go migrate
func CreateRiceBoxSourceDriver() (source.Driver, error) {
	sourceDriver := &RiceBoxSource{}
	err := sourceDriver.PopulateMigrations(rice.MustFindBox("./files"))
	return sourceDriver, err
}
