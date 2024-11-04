// Package repository is returning data processed from the database for the service layer.
package repository

import (
	"context"
	"database/sql"
	"errors"
	"security-proof/pkg/constants"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/db/security_proof/proof/table"
)

// DashboardQuerier interface is defining data related to simply querying dashboard data.
type DashboardQuerier interface {
	ProofNotConfirmer
	ProofNotUploader
}

// ProofNotConfirmer interface is defining data related to querying unconfirmed items.
type ProofNotConfirmer interface {
	NotConfirmProof(ctx context.Context) (proofs []*model.Proof, err error)
}

// ProofNotUploader interface is defining data related to querying uploaded items.
type ProofNotUploader interface {
	NotUploadProof(ctx context.Context) (proofs []*model.Proof, err error)
}

type dashboardQuery struct {
	db *sql.DB
}

// NewDashboardQuery function is returning DashboardQuerier interface accepting an DB.
func NewDashboardQuery(db *sql.DB) DashboardQuerier {
	return &dashboardQuery{db: db}
}

func (q *dashboardQuery) NotConfirmProof(ctx context.Context) ([]*model.Proof, error) {
	listStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.Num,
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.CreatedUserIdx,
			table.Proof.CreatedAt,
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
			table.Proof.UploadedUserIdx,
			table.Proof.UploadedAt,
			table.Proof.Confirm,
		).
		WHERE(table.Proof.Confirm.EQ(postgres.Int32(constants.NotConfirm))).
		LIMIT(10)

	dest := make([]*model.Proof, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *dashboardQuery) NotUploadProof(ctx context.Context) ([]*model.Proof, error) {
	listStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.Num,
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.CreatedUserIdx,
			table.Proof.CreatedAt,
			table.Proof.UpdatedUserIdx,
			table.Proof.UpdatedAt,
			table.Proof.UploadedUserIdx,
			table.Proof.UploadedAt,
			table.Proof.Confirm,
		).
		WHERE(table.Proof.UploadedAt.IS_NULL()).
		LIMIT(10)

	dest := make([]*model.Proof, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}
