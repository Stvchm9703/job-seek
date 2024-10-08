package config

import (
	"fmt"
	"log"

	"job-seek/pkg/config"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Host string `toml:"host" mapstructure:"host"`
	Port int    `toml:"port" mapstructure:"port"`
	// grpc server config
	Server config.ServerConfig `toml:"server" mapstructure:"server"`
	// internal services
	MeiliSearchService config.DatabaseConfig `toml:"meili_search_service" mapstructure:"meili_search_service"`
	SurrealDBService   config.DatabaseConfig `toml:"surreal_db_service" mapstructure:"surreal_db_service"`
	DBService          config.DatabaseConfig `toml:"sql_db_service" mapstructure:"sql_db_service"`
}

var (
	RuntimeConfig         ServiceConfig
	CurrentConfigFilePath string = ""
	ConfigFilePath        string = "config/user_management_service"
)

func Setup() {
	if CurrentConfigFilePath != "" {
		viper.SetConfigFile(CurrentConfigFilePath)
	} else {
		viper.SetConfigName(ConfigFilePath)
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("toml")

	// basic server config
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 60020)

	// grpc server config
	viper.SetDefault("server.max_connection_idle", 60)
	viper.SetDefault("server.max_connection_age", 300)
	viper.SetDefault("server.max_connection_age_grace", 10)
	viper.SetDefault("server.time", 60)
	viper.SetDefault("server.timeout", 20)
	viper.SetDefault("server.min_time", 5)
	viper.SetDefault("server.permit_without_stream", true)

	// internal services
	viper.SetDefault("meili_search_service.host", "localhost")
	viper.SetDefault("meili_search_service.port", 7700)
	viper.SetDefault("meili_search_service.api_key", "meilisearch")

	viper.SetDefault("surreal_db_service.host", "localhost")
	viper.SetDefault("surreal_db_service.port", 8654)

	err := viper.ReadInConfig()
	if err != nil {
		// t.Fatalf("unable to decode into struct, %v", err)
		log.Fatalf("unable to decode into struct, %v", err)
	}
	err = viper.Unmarshal(&RuntimeConfig)
	if err != nil {
		// t.Fatalf("unable to decode into struct, %v", err)
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func PrintConfig() {
	// viper.WriteConfigAs("config/user_management_service.runtime.toml")
	output, err := toml.Marshal(RuntimeConfig)
	if err != nil {
		log.Fatalln("failed to marshal config:")
		log.Fatalln(err)
	} else {
		fmt.Println("# runtime config")
		fmt.Println(string(output))
		fmt.Println()
	}
}
