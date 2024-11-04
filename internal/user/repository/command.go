// Package repository is returning data processed from the database for the service layer.
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
	dbmanage "security-proof/pkg/manage/db"
)

// UserCommander interface is defining data related to write commanding user data.
type UserCommander interface {
	dbmanage.Beginner
	dbmanage.Commiter
	dbmanage.Rollbacker
	UserCreator
	UserUpdater
	UserDeleter
}

// UserCreator interface is defining data related to commanding created item.
// Returns an error if the ID already exists during creation.
type UserCreator interface {
	CreateUser(ctx context.Context, user *model.User, tx *sql.Tx) (idx int32, err error)
}

// UserUpdater interface is defining data related to commanding updated item.
type UserUpdater interface {
	UpdateUser(ctx context.Context, user *model.User, tx *sql.Tx) (idx int32, err error)
}

// UserDeleter interface is defining data related to commanding deleted item.
type UserDeleter interface {
	DeleteUser(ctx context.Context, idx int32, tx *sql.Tx) (err error)
}

type userCommand struct {
	db *sql.DB
}

// NewUserCommand function is returning a UserCommander interface accepting an DB.
func NewUserCommand(db *sql.DB) UserCommander {
	return &userCommand{db: db}
}

func (c *userCommand) Begin(ctx context.Context) (*sql.Tx, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	return tx, errors.Join(constants.ErrBegin, err)
}

func (c *userCommand) Commit(_ context.Context, tx *sql.Tx) error {
	err := tx.Commit()
	return errors.Join(constants.ErrCommit, err)
}

func (c *userCommand) Rollback(_ context.Context, tx *sql.Tx) error {
	err := tx.Rollback()
	return errors.Join(constants.ErrRollback, err)
}

func (c *userCommand) CreateUser(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error) {
	insertStmt := table.User.
		INSERT(
			table.User.ID,
			table.User.Passwd,
			table.User.CreatedAt,
			table.User.UpdatedAt,
			table.User.Name,
			table.User.Email,
			table.User.Role,
		).
		MODEL(user).
		RETURNING(table.User.Idx)

	var executable qrm.Queryable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	dest := &model.User{}
	err := insertStmt.QueryContext(ctx, executable, dest)
	if err != nil {
		return 0, errors.Join(constants.ErrExecute, err)
	}

	return dest.Idx, nil
}

func (c *userCommand) UpdateUser(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error) {
	updateStmt := table.User.
		UPDATE(table.User.Name, table.User.Email, table.User.Role, table.User.UpdatedAt).
		MODEL(user).
		WHERE(table.User.Idx.EQ(postgres.Int32(user.Idx)))

	var executable qrm.Executable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	sqlResult, err := updateStmt.ExecContext(ctx, executable)
	if err != nil {
		return 0, errors.Join(constants.ErrExecute, err)
	}

	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return 0, errors.Join(constants.ErrRowResult, err)
	}
	if rowsAffected == 0 {
		return 0, constants.ErrItemNotFound
	}

	return user.Idx, nil
}

func (c *userCommand) DeleteUser(ctx context.Context, idx int32, tx *sql.Tx) error {
	deleteStmt := table.User.
		DELETE().
		WHERE(table.User.Idx.EQ(postgres.Int32(idx)))

	var executable qrm.Executable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	sqlResult, err := deleteStmt.ExecContext(ctx, executable)
	if err != nil {
		return errors.Join(constants.ErrExecute, err)
	}

	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return errors.Join(constants.ErrRowResult, err)
	}
	if rowsAffected == 0 {
		return constants.ErrItemNotFound
	}

	return nil
}
