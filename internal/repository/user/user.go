package user

import (
	"context"
	"github.com/IIIByte/user-service/internal/repository/postgres"
	"github.com/IIIByte/user-service/model/user"
)

const (
	UsersTable = "users"
)

const (
	IDField = "id"
)

const (
	UsersColumns = "id, login, name, password_hash"
)

type Repo struct {
	db *postgres.DB
}

func (r *Repo) GetByID(ctx context.Context, req *user.GetByIDRequest) (*user.GetByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) GetByLogin(ctx context.Context, req *user.GetByLoginRequest) (*user.GetByLoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Create(ctx context.Context, user *user.CreateRequest) (*user.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Update(ctx context.Context, user *user.UpdateRequest) (*user.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) MarkAsDeleted(ctx context.Context, req *user.MarkAsDeletedRequest) (*user.MarkAsDeletedRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) MarkAsBanned(ctx context.Context, req *user.MarkAsBannedRequest) (*user.MarkAsBannedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *postgres.DB) Repository {
	return &Repo{
		db: db,
	}
}
