package i18n

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type rawToml struct {
	Errors map[string]map[int]string `toml:"errors"`
}

func LoadDir(dir string) error {
	m := Init()

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".toml" {
			continue
		}

		lang := f.Name()[:len(f.Name())-5]
		path := filepath.Join(dir, f.Name())

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var raw rawToml
		if err := toml.Unmarshal(b, &raw); err != nil {
			return err
		}

		flat := make(map[int]string)
		for _, group := range raw.Errors {
			for code, msg := range group {
				flat[code] = msg
			}
		}

		m.Register(lang, flat)
	}

	return nil
}

func Validate(allCodes []int) error {
	for lang, mp := range mgr.data {
		for _, code := range allCodes {
			if _, ok := mp[code]; !ok {
				return fmt.Errorf("i18n missing: lang=%s code=%d", lang, code)
			}
		}
	}
	return nil
}
