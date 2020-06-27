package appconfig

type AppConfig struct {
	Neo4jSid  string
	Neo4jHost string
	Neo4jPort string

	EtcdServerEndpoints []string
}

var Config *AppConfig
