package model

import "time"

// User --- [Entity] 数据库对应模型 (GORM) ---
type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;type:varchar(50)"`
	Password  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
}

// CreateUserReq --- [Request] 接口入参 ---
type CreateUserReq struct {
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserReq struct {
	Email string `json:"email" binding:"email"`
}

// UserRes --- [Response] 接口出参 ---
type UserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ToResponse 将 Entity 转换为 Response DTO (单条)
func (u *User) ToResponse() *UserRes {
	if u == nil {
		return nil
	}
	return &UserRes{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

// ToResponses 将 Entity 切片转换为 Response 切片 (列表)
func ToResponses(users []*User) []*UserRes {
	res := make([]*UserRes, len(users))
	for i, u := range users {
		res[i] = u.ToResponse()
	}
	return res
}
