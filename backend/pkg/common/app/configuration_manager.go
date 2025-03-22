package app

import "github.com/SoumyaJha0410/backend/pkg/common/postgresql"

type ConfigurationManager struct{
	PostgresqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager{
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "root"
	dbName := "inventory"

	postgresqlConfig := postgresql.Config{
		Host:                  dbHost,
		Port:                  dbPort,
		UserName:              dbUser,
		Password:              dbPassword,
		DbName:                dbName,
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}

	return &ConfigurationManager{postgresqlConfig}
}