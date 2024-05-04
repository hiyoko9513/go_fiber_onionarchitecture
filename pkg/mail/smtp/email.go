package smtp

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"
)

type Email struct {
	From             string
	FromName         string
	To               []string
	Cc               []string
	Bcc              []string
	Subject          string
	TemplateFileName string
	Data             map[string]string
}

// SetData Add all data in map.
func (e *Email) SetData(data map[string]string) {
	for k, v := range data {
		e.Data[k] = v
	}
}

// SetAssetsPngImages sets the image images in the directory to the image structure.
// ※ only one directory hierarchy, only png supported
func (e *Email) SetAssetsPngImages(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if e.Data == nil {
		e.Data = make(map[string]string)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".png" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return err
		}

		enc := base64.StdEncoding.EncodeToString(data)

		name := strings.TrimSuffix(entry.Name(), ext)
		e.Data["Assets/Images/"+name] = enc
	}

	return nil
}
