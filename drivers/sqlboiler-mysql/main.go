package main

import (
	"github.com/mechiko/sqlboiler/v4/drivers"
	"github.com/mechiko/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
