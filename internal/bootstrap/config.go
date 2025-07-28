package bootstrap

import (
	. "github.com/thisismeamir/kage/pkg/form"
)

type Config struct {
	Name       string       `json:"name"`
	BasePath   string       `json:"base_path"`
	GraphPaths []FormPath   `json:"graph_paths"`
	NodePaths  []FormPath   `json:"node_paths"`
	MapPaths   []FormPath   `json:"map_paths"`
	Version    string       `json:"version"`
	Server     ServerConfig `json:"server"`
	Client     ClientConfig `json:"client"`
}

type ServerConfig struct {
	Port     int              `json:"port"`
	Host     string           `json:"host"`
	Api      ApiConfig        `json:"api"`
	Database []DatabaseConfig `json:"database"`
	Logging  LoggingConfig    `json:"logging"`
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
