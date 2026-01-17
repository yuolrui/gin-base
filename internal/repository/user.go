package repository

import (
	"errors"
	"sync"
)

// 模拟数据库模型
type User struct {
	ID    string
	Name  string
	Email string
}

var (
	users      = make(map[string]*User)
	usersMutex sync.Mutex
)

func FindUserByID(id string) (*User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, ok := users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func SaveUser(user *User) error {
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
	return "user_" + string(len(users)+1)
}
