package user

import "github.com/lstrihic/webapp/adapter/db/entity"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// MapUserEntityToDto map user entity to dto.
func MapUserEntityToDto(user *entity.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
	}
}
