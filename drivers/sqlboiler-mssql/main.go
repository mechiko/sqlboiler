package main

import (
	"github.com/mechiko/sqlboiler/v4/drivers"
	"github.com/mechiko/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
