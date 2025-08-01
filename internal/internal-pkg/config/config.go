package config

type Config struct {
	Name     string        `json:"name"`
	BasePath string        `json:"base_path"`
	Paths    []string      `json:"paths"`
	Version  string        `json:"version"`
	Logging  LoggingConfig `json:"logging"`
	Server   ServerConfig  `json:"server"`
	Client   ClientConfig  `json:"client"`
}

type ServerConfig struct {
	Port     int              `json:"port"`
	Host     string           `json:"host"`
	Api      ApiConfig        `json:"api"`
	Database []DatabaseConfig `json:"database"`
}

type ApiConfig struct {
	BaseUrl string `json:"base_url"`
	Version string `json:"version"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type LoggingConfig struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type ClientConfig struct {
	Web WebConfig `json:"web"`
}

type WebConfig struct {
	Port              int    `json:"port"`
	Host              string `json:"host"`
	NetworkAccessible bool   `json:"network_accessible"`
	Path              string `json:"path"`
	Build             string `json:"build"`
	Run               string `json:"run"`
}
