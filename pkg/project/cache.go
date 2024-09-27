package project

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func LoadCacheFileMap(cachePath string, cacheFileMap map[string]string) error {
	content, err := os.ReadFile(cachePath)
	if err != nil {
		return err
	}

	if err = toml.Unmarshal(content, &cacheFileMap); err != nil {
		return err
	}
	return nil
}

func SaveCacheFileMap(cachePath string, cacheFileMap map[string]string) error {
	b, err := toml.Marshal(cacheFileMap)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(cachePath), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(cachePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write(b); err != nil {
		return err
	}
	return nil
}
