package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"security-proof/internal/db/security_proof/user/model"
	"security-proof/internal/db/security_proof/user/table"
	"security-proof/pkg/constants"
)

// UserQuerier  interface is defining data related to simply querying user data.
type UserQuerier interface {
	UserReader
	UsersLister
	UserSignInner
}

// UserReader interface is defining data related to querying read data.
type UserReader interface {
	ReadUserByIdx(ctx context.Context, idx int32) (user *model.User, err error)
	ReadUserByID(ctx context.Context, id string) (user *model.User, err error)
}

// UsersLister interface is defining data related to querying listed data.
type UsersLister interface {
	AllUsers(ctx context.Context) ([]*model.User, error)
	SearchUsers(ctx context.Context, id string) ([]*model.User, error)
}

// UserSignInner interface is defining data related to querying sign in user.
type UserSignInner interface {
	SignInUser(ctx context.Context, id string) (user *model.User, err error)
}

type userQuery struct {
	db *sql.DB
}

// NewUserQuery function is returning UserQuerier interface accepting an DB.
func NewUserQuery(db *sql.DB) UserQuerier {
	return &userQuery{db: db}
}

func (q *userQuery) ReadUserByIdx(ctx context.Context, idx int32) (*model.User, error) {
	readStmt := table.User.
		SELECT(
			table.User.Idx,
			table.User.ID,
			table.User.CreatedAt,
			table.User.UpdatedAt,
			table.User.Name,
			table.User.Email,
			table.User.Role,
		).
		WHERE(table.User.Idx.EQ(postgres.Int32(idx))).
		LIMIT(1)

	dest := &model.User{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *userQuery) ReadUserByID(ctx context.Context, id string) (*model.User, error) {
	readStmt := table.User.
		SELECT(
			table.User.Idx,
			table.User.ID,
			table.User.CreatedAt,
			table.User.UpdatedAt,
			table.User.Name,
			table.User.Email,
			table.User.Role,
		).
		WHERE(table.User.ID.LIKE(postgres.String(id))).
		LIMIT(1)

	dest := &model.User{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *userQuery) AllUsers(ctx context.Context) ([]*model.User, error) {
	listStmt := table.User.
		SELECT(
			table.User.Idx,
			table.User.ID,
			table.User.Name,
			table.User.Email,
			table.User.Role,
		)

	dest := make([]*model.User, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *userQuery) SearchUsers(ctx context.Context, id string) ([]*model.User, error) {
	listStmt := table.User.
		SELECT(
			table.User.Idx,
			table.User.ID,
			table.User.Name,
			table.User.Email,
			table.User.Role,
		).WHERE(table.User.ID.LIKE(postgres.String(id)))

	dest := make([]*model.User, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *userQuery) SignInUser(ctx context.Context, id string) (*model.User, error) {
	readStmt := table.User.
		SELECT(
			table.User.Idx,
			table.User.ID,
			table.User.Passwd,
			table.User.Role,
		).
		WHERE(table.User.ID.LIKE(postgres.String(id))).
		LIMIT(1)

	dest := &model.User{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}
