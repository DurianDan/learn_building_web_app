package congfig 

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// blank import, need some of the initiation code in the blank package 
	// but wont use any method presented in the package
)

var (
	db * gorm.DB
) // use brackets for creating multiple variables
// db is a pointer


func Connect(){
	// this function help the app to connect to the database 
	d, err := gorm.Open( // The DB will return back the error to the var err, and value to var d
		"mysql", // choosing the find of SQL database, (can also be PostGre, SQLite)
		"root:MySQL_000/simplerest?charset=utf8&parseTime=True&loc=Local", 
		// {user-name}:{password}/{database}?{MySQL specification}
	)
	if err!= nil {
		panic(err) // check if the err has been returned (not nil)
	}
	db = d 
}

func GetDB() *gorm.DB{
	// return the db from the database
	return db // to talk to the database 
}

