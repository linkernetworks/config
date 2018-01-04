package config

type AppConfig struct {
	Logger               LoggerConfig   `json:"logger"`
	Session              *SessionConfig `json:"session"`
	EnableAuthentication bool           `json:"enableAuthentication"`

	// TODO: move this the dataset config
	MaxThumbnailWidth  uint   `json:"maxThumbnailWidth"`
	MaxThumbnailHeight uint   `json:"maxThumbnailHeight"`
	DbVersion          string `json:"dbVersion"`
	Version            string `json:"version"`
}