package config

var (
	RABBITUSER     = "user"
	RABBITPASSWORD = "password"
	RABBITHOST     = "rabbit"
	EXCHANGE       = "users"
	RABBITPORT     = 5672

	ENDPOINTS = []string{"http://host.docker.internal:8090/items/user", "http://host.docker.internal:80", "http://host.docker.internal:9001/messages/user"}
)
