package config

import "time"

var (
	CCMAXSIZE      int64  = 1000
	CCITEMSTOPRUNE uint32 = 100
	CCDEFAULTTTL          = 30 * time.Second

	DBHOST     = "items_db"
	DBPORT     = 27017
	COLLECTION = "items"

	RABBITHOST = "rabbit"
	RABBITPORT = 5672
	EXCHANGE   = "users"
)
