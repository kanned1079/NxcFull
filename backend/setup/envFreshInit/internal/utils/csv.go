package utils

import (
	"encoding/csv"
	"envFreshInit/internal/model"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ReadCSVData 从 CSV 文件读取数据
func ReadCSVData(csvFilePath string) ([]model.SiteSetting, error) {
	// 打开 CSV 文件
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open csv file: %w", err)
	}
	defer file.Close()

	// 创建 CSV 读取器
	reader := csv.NewReader(file)

	// 读取 CSV 文件的所有行
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failure on reading csv file: %w", err)
	}

	var settings []model.SiteSetting
	// 跳过标题行
	for _, line := range lines[1:] {
		// 解析每行数据
		id, err := strconv.ParseUint(line[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve id: %w", err)
		}
		category := line[1]
		key := line[2]
		value := []byte(line[3])
		createdAt, err := time.Parse("2006-01-02 15:04:05", line[4])
		if err != nil {
			return nil, fmt.Errorf("cannot resolve create_at: %w", err)
		}
		updatedAt, err := time.Parse("2006-01-02 15:04:05", line[5])
		if err != nil {
			return nil, fmt.Errorf("cannot resolve update_at: %w", err)
		}

		// 创建 SiteSetting 实例
		setting := model.SiteSetting{
			Id:        uint(id),
			Category:  category,
			Key:       key,
			Value:     value,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		settings = append(settings, setting)
	}

	return settings, nil
}
