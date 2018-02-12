package config

type AppConfig struct {
	Session  *SessionConfig  `json:"session"`
	Logger   LoggerConfig    `json:"logger"`
	Brand    BrandConfig     `json:"brandIdentifier"`
	Session  *SessionConfig  `json:"session"`
	Logger   LoggerConfig    `json:"logger"`
	Socketio *SocketioConfig `json:"socketio"`

	EnableAuthentication bool `json:"enableAuthentication"`

	// TODO: move this the dataset config
	MaxThumbnailWidth  uint   `json:"maxThumbnailWidth"`
	MaxThumbnailHeight uint   `json:"maxThumbnailHeight"`
	DbVersion          string `json:"dbVersion"`
	Version            string `json:"version"`
}
