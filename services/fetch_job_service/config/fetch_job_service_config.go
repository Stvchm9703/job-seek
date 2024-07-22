package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Host string `toml:"host" mapstructure:"host"`
	Port int    `toml:"port" mapstructure:"port"`

	Server           ServerConfig                      `toml:"server" mapstructure:"server"`
	SeekService      SeekServiceConfig                 `toml:"seek_service" mapstructure:"seek_service"`
	YahooSearch      YahooSearchConfig                 `toml:"yahoo_search" mapstructure:"yahoo_search"`
	DataStoreService map[string]DataStoreServiceConfig `toml:"data_store_service" mapstructure:"data_store_service"`
}

type ServerConfig struct {
	MaxConnectionIdle     int  `toml:"max_connection_idle" mapstructure:"max_connection_idle"`
	MaxConnectionAge      int  `toml:"max_connection_age" mapstructure:"max_connection_age"`
	MaxConnectionAgeGrace int  `toml:"max_connection_age_grace" mapstructure:"max_connection_age_grace"`
	Time                  int  `toml:"time" mapstructure:"time"`
	Timeout               int  `toml:"timeout" mapstructure:"timeout"`
	MinTime               int  `toml:"min_time" mapstructure:"min_time"`
	PermitWithoutStream   bool `toml:"permit_without_stream" mapstructure:"permit_without_stream"`
}

type SeekServiceConfig struct {
	Domain string `toml:"domain" mapstructure:"domain"`
}

type YahooSearchConfig struct {
	Domain string `toml:"domain" mapstructure:"domain"`
}

type DataStoreServiceConfig struct {
	Host string `toml:"host" mapstructure:"host"`
	Port int    `toml:"port" mapstructure:"port"`
}

var (
	RuntimeConfig         ServiceConfig
	CurrentConfigFilePath string = ""
	ConfigFilePath        string = "config/fetch_job_service"
)

func Setup() {
	if CurrentConfigFilePath != "" {
		viper.SetConfigFile(CurrentConfigFilePath)
	} else {
		viper.SetConfigName(ConfigFilePath)
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("toml")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 60010)

	viper.SetDefault("seek_service.domain", "https://www.seek.com.au")
	viper.SetDefault("yahoo_search.domain", "https://au.search.yahoo.com")

	viper.SetDefault("server.max_connection_idle", 60)
	viper.SetDefault("server.max_connection_age", 300)
	viper.SetDefault("server.max_connection_age_grace", 10)
	viper.SetDefault("server.time", 60)
	viper.SetDefault("server.timeout", 20)
	viper.SetDefault("server.min_time", 5)
	viper.SetDefault("server.permit_without_stream", true)

	viper.SetDefault("data_store_service[\"base\"].host", "localhost")
	viper.SetDefault("data_store_service[\"base\"].post", 60020)

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
	// viper.WriteConfigAs("config/fetch_job_service.runtime.toml")
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
