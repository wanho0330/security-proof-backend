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
)

// ProofQuerier interface is defining data related to simply querying proof data.
type ProofQuerier interface {
	ProofReader
	ProofsLister
	ProofImageReader
	ProofLogReader
}

// ProofReader interface is defining data related to querying read data.
type ProofReader interface {
	ReadProof(ctx context.Context, idx int32) (proof *model.Proof, err error)
}

// ProofsLister interface is defining data related to querying listed data.
type ProofsLister interface {
	AllProofs(ctx context.Context) ([]*model.Proof, error)
	SearchProofs(ctx context.Context, category string) (proofs []*model.Proof, err error)
}

// ProofImageReader interface is defining data related to querying read image data.
type ProofImageReader interface {
	ReadFirstProofImage(ctx context.Context, idx int32) (proof *model.Proof, err error)
	ReadSecondProofImage(ctx context.Context, idx int32) (proof *model.Proof, err error)
}

// ProofLogReader interface is defining data related to querying read log data.
type ProofLogReader interface {
	ReadProofLog(ctx context.Context, idx int32) (log *model.Proof, err error)
}

type proofQuery struct {
	db *sql.DB
}

// NewProofQuery function is returning ProofQuerier interface accepting an DB.
func NewProofQuery(db *sql.DB) ProofQuerier {
	return &proofQuery{db: db}
}

func (q *proofQuery) ReadProof(ctx context.Context, idx int32) (*model.Proof, error) {
	readStmt := table.Proof.
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
			table.Proof.TokenID,
		).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(idx))).
		LIMIT(1)

	dest := &model.Proof{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *proofQuery) AllProofs(ctx context.Context) ([]*model.Proof, error) {
	listStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.Num,
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.UploadedAt,
			table.Proof.Confirm,
		)

	dest := make([]*model.Proof, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *proofQuery) SearchProofs(ctx context.Context, category string) ([]*model.Proof, error) {
	listStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.Num,
			table.Proof.Category,
			table.Proof.Description,
			table.Proof.UploadedAt,
			table.Proof.Confirm,
		).WHERE(table.Proof.Category.LIKE(postgres.String(category)))

	dest := make([]*model.Proof, 0)
	err := listStmt.QueryContext(ctx, q.db, &dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *proofQuery) ReadFirstProofImage(ctx context.Context, idx int32) (*model.Proof, error) {
	readStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.FirstImagePath,
		).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(idx))).
		LIMIT(1)

	dest := &model.Proof{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *proofQuery) ReadSecondProofImage(ctx context.Context, idx int32) (*model.Proof, error) {
	readStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.SecondImagePath,
		).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(idx))).
		LIMIT(1)

	dest := &model.Proof{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}

func (q *proofQuery) ReadProofLog(ctx context.Context, idx int32) (*model.Proof, error) {
	readStmt := table.Proof.
		SELECT(
			table.Proof.Idx,
			table.Proof.LogPath,
		).
		WHERE(table.Proof.Idx.EQ(postgres.Int32(idx))).
		LIMIT(1)

	dest := &model.Proof{}

	err := readStmt.QueryContext(ctx, q.db, dest)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.Join(constants.ErrQuery, constants.ErrItemNotFound)
	} else if err != nil {
		return nil, errors.Join(constants.ErrQuery, err)
	}

	return dest, nil
}
