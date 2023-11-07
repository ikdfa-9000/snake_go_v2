package cfg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const fileCfgName = "config.json"

var defaultConfig = Config{
	DeskRows:       8,
	DeskColumns:    8,
	DeskFrameSpeed: 400,
}

type Config struct {
	DeskRows       int `json:"rows"`
	DeskColumns    int `json:"columns"`
	DeskFrameSpeed int `json:"frameSpeed"`
}

func InitJsonConfig(configFolderPath string) (Config, error) {
	configPath := configFolderPath + fileCfgName
	var config Config

	input, err := os.ReadFile(configPath)
	if err != nil || len(input) == 0 {
		// Если не получилось прочитать, создаем дефолтный конфиг
		input, err = createDefaultConfig(configPath)
		if err != nil {
			return config, err
		}
	}

	if err := json.Unmarshal(input, &config); err != nil {
		return config, fmt.Errorf("ошибка десериализвации конфига: %v", err.Error())
	}

	return config, nil
}

// Deprecated: Теперь мы используем чтение конфигурации из файла JSON.
// Использовать InitJsonConfig(). Также эта функция не будет работать,
// так как отсутствует файл конфигурации: configs/config.yml
func InitConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config.DeskRows = viper.GetInt("desk.rows")
	config.DeskColumns = viper.GetInt("desk.columns")
	config.DeskFrameSpeed = viper.GetInt("desk.frameSpeed")

	return config, nil
}

// Функция сооздает пустой файл "configPath", сериализует (преобразоваывает)
// дефолтную структуры конфига в JSON, записывает в новый файл и возвращает
// сериализованные объект в виде []byte
func createDefaultConfig(configPath string) ([]byte, error) {
	var buf []byte
	file, err := os.Create(configPath)
	if err != nil {
		return buf, fmt.Errorf("ошибка создания дефолтного конфига: %s", err.Error())
	}
	defer file.Close()

	// Сериализация дефолтного конфига
	output, err := json.Marshal(defaultConfig)
	if err != nil {
		return buf, err
	}

	writer := bufio.NewWriter(file)
	if _, err := writer.Write(output); err != nil {
		return buf, fmt.Errorf("ошибка создания дефолтного конфига: %s", err.Error())
	}

	if err = writer.Flush(); err != nil {
		return buf, fmt.Errorf("ошибка создания дефолтного конфига: %s", err.Error())
	}

	return output, nil
}
