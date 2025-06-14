package main

type Config struct {
	AppName           string   `mapstructure:"app-name"`
	HttpServeAddress  string   `mapstructure:"http-serve-address"`
	AxonServerAddress string   `mapstructure:"axon-server-address"`
	Dsn               string   `mapstructure:"dsn"`
	Profiles          []string `mapstructure:"profiles"`
}
