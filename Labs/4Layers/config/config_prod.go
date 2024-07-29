// +build prod

package config

const (
	MYSQL_HOST             = "mysql://127.0.0.1:3306"
	MYSQL_DATABASE         = "bookmark"
	MYSQL_CONNECTION_POOL  = 50
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
