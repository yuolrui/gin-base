package i18n

import (
	"fmt"
	"sync"
)

type Manager struct {
	data map[string]map[int]string // lang -> code -> msg
}

var (
	mgr  *Manager
	once sync.Once
)

func Init() *Manager {
	once.Do(func() {
		mgr = &Manager{
			data: make(map[string]map[int]string),
		}
	})
	return mgr
}

func (m *Manager) Register(lang string, messages map[int]string) {
	m.data[lang] = messages
}

func (m *Manager) Get(lang string, code int) string {
	if mp, ok := m.data[lang]; ok {
		if msg, ok := mp[code]; ok {
			return msg
		}
	}
	// fallback
	if mp, ok := m.data["zh-CN"]; ok {
		return mp[code]
	}
	return fmt.Sprintf("unknown error (%d)", code)
}
