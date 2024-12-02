package repository

import (
	"context"
	"errors"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewFeatureFlagRepository(logger *types.Logger, db *pgxpool.Pool) *featureFlagRepo {
	return &featureFlagRepo{
		logger: logger,
		db:     db,
	}
}

type featureFlagRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *featureFlagRepo) Create(ctx context.Context,
	orgID int64,
	applicationID int64,
	name string,
	isEnabled bool,
	createdBy int64,
) (*types.FeatureFlag, error) {
	var (
		featureFlag types.FeatureFlag
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagCreate,
		orgID,
		applicationID,
		name,
		isEnabled,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &featureFlag, nil
}

func (r *featureFlagRepo) GetMany(ctx context.Context,
	orgID int64,
	applicationID int64,
) ([]types.FeatureFlag, error) {
	var (
		featureFlags []types.FeatureFlag
		rows         pgx.Rows
		err          error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagGetMany,
		orgID,
		applicationID,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if featureFlags, err = pgx.CollectRows(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return featureFlags, nil
}

func (r *featureFlagRepo) GetOne(ctx context.Context,
	orgID int64,
	applicationID int64,
	id *int64,
	uuid *string,
	name *string,
) (*types.FeatureFlag, error) {
	var (
		featureFlag types.FeatureFlag
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagGetOne,
		orgID,
		applicationID,
		id,
		uuid,
		name,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &featureFlag, err
}

func (r *featureFlagRepo) Update(ctx context.Context,
	orgID int64,
	applicationID int64,
	id int64,
	name string,
	isEnabled bool,
	modifiedBy int64,
) (*types.FeatureFlag, error) {
	var (
		featureFlag types.FeatureFlag
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagUpdate,
		orgID,
		applicationID,
		id,
		name,
		isEnabled,
		modifiedBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &featureFlag, nil
}

func (r *featureFlagRepo) Delete(ctx context.Context,
	orgID int64,
	applicationID int64,
	id int64,
) error {
	row := r.db.QueryRow(ctx,
		queries.FeatureFlagDelete,
		orgID,
		applicationID,
		id,
	)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(ctx, r.logger, err)
	}

	if numDeleted < 1 {
		return types.ErrNotFound
	}

	if numDeleted > 1 {
		return errors.New(
			fmt.Sprintf("expected to delete 1 row, deleted %v", numDeleted),
		)
	}

	return nil
}
