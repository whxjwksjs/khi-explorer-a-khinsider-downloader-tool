package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Format       string `yaml:"format"`
	DownloadPath string `yaml:"download_path"`
	Player       string `yaml:"player"`
}

func GetConfigPath() string {
	// Portable Mode: Get the directory where khi-explorer.exe is located
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		return filepath.Join(exeDir, ".khi_explorer.yaml")
	}

	// Fallback to current working directory
	return ".khi_explorer.yaml"
}

func LoadConfig() (Config, error) {
	cfgPath := GetConfigPath()

	defaultCfg := Config{
		Format:       "flac",
		DownloadPath: "./downloads", // Default: creates 'downloads' folder next to the .exe
		Player:       "mpv",
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		// If config doesn't exist yet, auto-create it next to the .exe
		_ = defaultCfg.Save()
		return defaultCfg, nil
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return defaultCfg, err
	}

	if cfg.Format == "" {
		cfg.Format = defaultCfg.Format
	}
	if cfg.DownloadPath == "" {
		cfg.DownloadPath = defaultCfg.DownloadPath
	}

	return cfg, nil
}

func (c *Config) Save() error {
	cfgPath := GetConfigPath()
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(cfgPath, data, 0644)
}
