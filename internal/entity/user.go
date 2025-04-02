package entity

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	// UUID (string) пользователя (не изменяем)
	ID string

	// Логин пользователя (не изменяем)
	Login string
	// Имя пользователя (не изменяемо)
	Name string

	// Хеш пароля (изменяемо)
	PasswordHash string

	// Флаг отвечающий за блокировку пользователя
	IsBanned bool

	// Флаг отвечающий за удаление пользователя
	// NOTE: Если пользователь помечен таким флагом, значит он в какой то момент будет удален из БД
	IsDeleted bool

	// Дата создания
	CreatedAt timestamppb.Timestamp
	// Дата последнего изменения
	UpdatedAt timestamppb.Timestamp
}

func (u *User) Update(user *User) {
	if u.Name != user.Name {
		u.Name = user.Name
	}
}
