package main

import (
	"github.com/mechiko/sqlboiler/v4/drivers"
	"github.com/mechiko/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
