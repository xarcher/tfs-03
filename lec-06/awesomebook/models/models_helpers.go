package models

import (
	"fmt"
)

func (dBConfig DBConfig) String() string {
	return fmt.Sprintf("[Driver: %s, Host: %s, Port: %s, Username: %s, Passwd: %s, DBName: %s]",
		dBConfig.Driver, dBConfig.Host, dBConfig.Port, dBConfig.Username, dBConfig.Passwd, dBConfig.DBName)
}
