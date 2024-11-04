// Package repository is returning data processed from the database for the service layer.
package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/db/security_proof/proof/table"
	"security-proof/pkg/constants"
	dbmanage "security-proof/pkg/manage/db"
)

// ProofCommander interface is defining data related to write commanding proof data.
type ProofCommander interface {
	dbmanage.Beginner
	dbmanage.Commiter
	dbmanage.Rollbacker
	ProofCreator
	ProofUpdater
	ProofDeleter
	ProofUploader
	ProofConfirmer
}

// ProofCreator interface is defining data related to commanding created item.
type ProofCreator interface {
	CreateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (idx int32, err error)
}

// ProofUpdater interface is defining data related to commanding updated item.
type ProofUpdater interface {
	UpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (idx int32, err error)
}

// ProofUploader interface is defining data related to commanding uploaded item.
type ProofUploader interface {
	UploadProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (idx int32, err error)
}

// ProofDeleter interface is defining data related to commanding deleted item.
type ProofDeleter interface {
	DeleteProof(ctx context.Context, idx int32, tx *sql.Tx) (err error)
}

// ProofConfirmer interface is defining data related to commanding confirmed item.
type ProofConfirmer interface {
	ConfirmProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error
	ConfirmUpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error
}

type proofCommand struct {
	db *sql.DB
}

// NewProofCommand function is returning a ProofCommander interface accepting an DB.
func NewProofCommand(db *sql.DB) ProofCommander {
	return &proofCommand{db: db}
}

func (c *proofCommand) Begin(ctx context.Context) (*sql.Tx, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	return tx, errors.Join(constants.ErrBegin, err)
}

func (c *proofCommand) Commit(_ context.Context, tx *sql.Tx) error {
	err := tx.Commit()
	return errors.Join(constants.ErrCommit, err)
}

func (c *proofCommand) Rollback(_ context.Context, tx *sql.Tx) error {
	err := tx.Rollback()
	return errors.Join(constants.ErrRollback, err)
}

func (c *proofCommand) CreateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	insertStmt := table.Proof.
		INSERT(
			table.Proof.Num,
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.UploadedUserIdx,
			table.Proof.CreatedUserIdx,
			table.Proof.CreatedAt,
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
		).
		MODEL(proof).
		RETURNING(table.Proof.Idx)

	var executable qrm.Queryable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	dest := &model.Proof{}
	err := insertStmt.QueryContext(ctx, executable, dest)
	if err != nil {
		return 0, errors.Join(constants.ErrExecute, err)
	}

	return dest.Idx, nil
}

func (c *proofCommand) UpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	updateStmt := table.Proof.
		UPDATE(
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.UploadedUserIdx,
			table.Proof.Confirm,
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
		).
		MODEL(proof).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(proof.Idx)))

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

	return proof.Idx, nil
}

func (c *proofCommand) DeleteProof(ctx context.Context, idx int32, tx *sql.Tx) error {
	deleteStmt := table.Proof.
		DELETE().
		WHERE(table.Proof.Idx.EQ(postgres.Int32(idx)))

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

func (c *proofCommand) UploadProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	updateStmt := table.Proof.
		UPDATE(
			table.Proof.FirstImagePath,
			table.Proof.SecondImagePath,
			table.Proof.LogPath,
			table.Proof.UploadedUserIdx,
			table.Proof.UploadedAt,
			table.Proof.Confirm,
		).
		MODEL(proof).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(proof.Idx)))

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

	return proof.Idx, nil
}

func (c *proofCommand) ConfirmProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
	updateStmt := table.Proof.
		UPDATE(
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
			table.Proof.Confirm,
			table.Proof.TokenID,
		).
		MODEL(proof).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(proof.Idx)))

	var executable qrm.Executable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	sqlResult, err := updateStmt.ExecContext(ctx, executable)
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

func (c *proofCommand) ConfirmUpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
	updateStmt := table.Proof.
		UPDATE(
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
			table.Proof.Confirm,
		).
		MODEL(proof).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(proof.Idx)))

	var executable qrm.Executable
	if tx != nil {
		executable = tx
	} else {
		executable = c.db
	}

	sqlResult, err := updateStmt.ExecContext(ctx, executable)
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
