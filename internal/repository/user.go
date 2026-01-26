package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/yuolrui/gin-base/internal/errno"
	"github.com/yuolrui/gin-base/internal/model"
)

var (
	users      = make(map[string]*model.User)
	usersMutex sync.Mutex
)

func FindUserByID(id string) (*model.User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, ok := users[id]
	if !ok {
		return nil, errno.ErrUserNotExist
	}
	return user, nil
}

func SaveUser(user *model.User) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	if user.ID == "" {
		user.ID = generateID()
	}
	users[user.ID] = user
	return nil
}

func generateID() string {
	// 简单生成ID示例，实际建议用UUID或数据库自增
	return "user_" + uuid.NewString()
}
