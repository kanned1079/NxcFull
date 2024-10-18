package local

import (
	"encoding/xml"
	"fmt"
	"os"
)

// MyLocalConfig 本地配置
var MyLocalConfig LocalConfig

func (e *LocalConfig) SaveConfigToXml(filePath string) error {
	data, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config to file %s: %w", filePath, err)
	}
	return nil
}

func (e *LocalConfig) LoadConfigFromXml(filePath string) error {
	//var config EtcdConfig
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, &MyLocalConfig)
	if err != nil {
		return err
	}
	return nil
}
