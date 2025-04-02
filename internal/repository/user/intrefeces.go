package user

import (
	"context"
	userProto "github.com/IIIByte/user-service/model/user"
)

type Repository interface {
	// GetByID - получить по ID пользователя
	GetByID(ctx context.Context, req *userProto.GetByIDRequest) (*userProto.GetByIDResponse, error)

	// GetByLogin - получить по Login пользователя
	GetByLogin(ctx context.Context, req *userProto.GetByLoginRequest) (*userProto.GetByLoginResponse, error)

	// Create - создать пользователя
	Create(ctx context.Context, user *userProto.CreateRequest) (*userProto.CreateResponse, error)

	// Update - обновить данные у пользователя
	Update(ctx context.Context, user *userProto.UpdateRequest) (*userProto.UpdateResponse, error)

	// MarkAsDeleted - пометить в очередь на удаление
	MarkAsDeleted(ctx context.Context, req *userProto.MarkAsDeletedRequest) (*userProto.MarkAsDeletedRequest, error)

	// MarkAsBanned - пометить как заблокированного
	MarkAsBanned(ctx context.Context, req *userProto.MarkAsBannedRequest) (*userProto.MarkAsBannedResponse, error)
}
