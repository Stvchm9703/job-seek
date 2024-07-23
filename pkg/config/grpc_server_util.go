package config

type ServerConfig struct {
	MaxConnectionIdle     int  `toml:"max_connection_idle" mapstructure:"max_connection_idle"`
	MaxConnectionAge      int  `toml:"max_connection_age" mapstructure:"max_connection_age"`
	MaxConnectionAgeGrace int  `toml:"max_connection_age_grace" mapstructure:"max_connection_age_grace"`
	Time                  int  `toml:"time" mapstructure:"time"`
	Timeout               int  `toml:"timeout" mapstructure:"timeout"`
	MinTime               int  `toml:"min_time" mapstructure:"min_time"`
	PermitWithoutStream   bool `toml:"permit_without_stream" mapstructure:"permit_without_stream"`
}

type DatabaseConfig struct {
	Host   string `toml:"host" mapstructure:"host"`
	Port   int    `toml:"port" mapstructure:"port"`
	ApiKey string `toml:"api_key" mapstructure:"api_key"`
}

type ApiService struct {
	Domain   string `toml:"domain" mapstructure:"domain"`
	CoolDown int    `toml:"cool_down" mapstructure:"cool_down"`
}
