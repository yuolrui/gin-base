package i18n

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// rawToml 是 TOML 文件的中间结构，用于反序列化
type rawToml struct {
	Errors map[string]map[int]string `toml:"errors"`
}

// LoadDir 加载指定目录下所有 .toml 文件作为语言包
func LoadDir(dir string) error {
	Init()

	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read i18n dir failed: %w", err)
	}

	for _, f := range files {
		if shouldSkip(f) {
			continue
		}

		lang := strings.TrimSuffix(f.Name(), ".toml")
		path := filepath.Join(dir, f.Name())

		// 将解析逻辑抽离
		flat, err := parseLanguageFile(path)
		if err != nil {
			return fmt.Errorf("parse lang [%s] failed: %w", lang, err)
		}

		Register(lang, flat)
	}

	return checkDefaultLang()
}

// 抽离逻辑 1：判断是否跳过文件
func shouldSkip(f os.DirEntry) bool {
	return f.IsDir() || filepath.Ext(f.Name()) != ".toml"
}

// 抽离逻辑 2：解析单个文件内容
func parseLanguageFile(path string) (map[int]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var raw rawToml
	if err := toml.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	return flattenErrors(raw.Errors)
}

// 抽离逻辑 3：扁平化错误码 map（解决三层嵌套循环）
func flattenErrors(groups map[string]map[int]string) (map[int]string, error) {
	flat := make(map[int]string)
	for groupName, codes := range groups {
		for code, msg := range codes {
			if _, exists := flat[code]; exists {
				return nil, fmt.Errorf("duplicate code in group=%s code=%d", groupName, code)
			}
			flat[code] = msg
		}
	}
	return flat, nil
}

// 抽离逻辑 4：默认语言校验
func checkDefaultLang() error {
	if _, ok := mgr.data[DefaultLang]; !ok {
		return fmt.Errorf("default language %s not loaded", DefaultLang)
	}
	return nil
}

// Validate 检查所有语言包是否都包含所有错误码
func Validate(allCodes []int) error {
	if mgr == nil {
		return fmt.Errorf("i18n not initialized")
	}
	return mgr.validate(allCodes)
}

func (m *Manager) validate(allCodes []int) error {
	for lang, mp := range m.data {
		for _, code := range allCodes {
			if _, ok := mp[code]; !ok {
				return fmt.Errorf("i18n missing: lang=%s code=%d", lang, code)
			}
		}
	}
	return nil
}
