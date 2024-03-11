package config

import (
	"github.com/beatrausch/fincacal/internal/finca"
	"os"
	"sigs.k8s.io/yaml"
)

func ReadConfig(file string) (*finca.Accommodation, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := new(finca.Accommodation)
	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, err
	}
	return config, nil
}
