package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	English   = "en"
	Japanese  = "ja"
	ChineseTW = "zh_TW"

	LangDir = "./pkg/i18n/lang"

	EnglishFile   = "en.json"
	JapaneseFile  = "ja.json"
	ChineseTWFile = "zh_TW.json"
)

func GetLanguageMap(lang string) (map[string]string, error) {
	var jsonFilePath string
	switch lang {
	case English:
		jsonFilePath = filepath.Join(LangDir, EnglishFile)
	case Japanese:
		jsonFilePath = filepath.Join(LangDir, JapaneseFile)
	case ChineseTW:
		jsonFilePath = filepath.Join(LangDir, ChineseTWFile)
	default:
		return nil, fmt.Errorf("invalid language: %s", lang)
	}

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var langMap map[string]string
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&langMap)
	if err != nil {
		return nil, err
	}

	return langMap, nil
}
