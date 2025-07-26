package init

import "github.com/thisismeamir/kage/internal/models"

type Config struct {
	Name        string              `json:"name"`
	BasePath    string              `json:"base_path"`
	ModulePaths []models.ModulePath `json:"module_paths"`
	AtomPaths   []models.AtomPath   `json:"atom_paths"`
	Version     string              `json:"version"`
	Server      ServerConfig        `json:"server"`
	Client      ClientConfig        `json:"client"`
}

type ServerConfig struct {
	Port     int            `json:"port"`
	Host     string         `json:"host"`
	Api      ApiConfig      `json:"api"`
	Database DatabaseConfig `json:"database"`
	Logging  LoggingConfig  `json:"logging"`
}

type ApiConfig struct {
	BaseUrl string `json:"base_url"`
	Version string `json:"version"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type LoggingConfig struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

type ClientConfig struct {
	Web WebConfig `jsong:"web"`
}

type WebConfig struct {
	Port              int    `json:"port"`
	Host              string `json:"host"`
	NetworkAccessible bool   `json:"network_accessible"`
	Path              string `json:"path"`
	Build             string `json:"build"`
	Run               string `json:"run"`
}
