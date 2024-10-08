package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var EnvConfig Config

// SaveConfigToYaml 将配置保存到YAML文件
func SaveConfigToYaml(config *Config, filePath string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config to file %s: %w", filePath, err)
	}
	return nil
}

// LoadConfigFromYaml 从YAML文件加载配置
func LoadConfigFromYaml(filePath string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file %s: %w", filePath, err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return config, nil
}
