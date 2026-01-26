package i18n

import (
	"fmt"
	"sync"
)

const DefaultLang = "zh-CN"

type Manager struct {
	data map[string]map[int]string
}

var (
	mgr  *Manager
	once sync.Once
)

func Init() {
	once.Do(func() {
		mgr = &Manager{
			data: make(map[string]map[int]string),
		}
	})
}

func Register(lang string, messages map[int]string) {
	if mgr == nil {
		panic("i18n not initialized")
	}
	if _, ok := mgr.data[lang]; ok {
		panic("duplicated i18n register: " + lang)
	}
	mgr.data[lang] = messages
}

func Get(lang string, code int) string {
	if mp, ok := mgr.data[lang]; ok {
		if msg, ok := mp[code]; ok {
			return msg
		}
	}
	if mp, ok := mgr.data[DefaultLang]; ok {
		if msg, ok := mp[code]; ok {
			return msg
		}
	}
	return fmt.Sprintf("unknown error (%d)", code)
}
