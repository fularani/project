package dbhelpers

import (
	"gorm.io/gorm" //module helps to execute queries over specific databases
)

//function for Insert query to add values into specific table
func InsertQuery(conn *gorm.DB, result interface{}) *gorm.DB {
	return conn.Create(result)
}

