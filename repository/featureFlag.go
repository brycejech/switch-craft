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

func NewFeatureFlagRepository(db *pgxpool.Pool) *featureFlagRepo {
	return &featureFlagRepo{
		db: db,
	}
}

type featureFlagRepo struct {
	db *pgxpool.Pool
}

func (r *featureFlagRepo) Create(ctx context.Context,
	tenantID int64,
	applicationID int64,
	name string,
	slug string,
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
		tenantID,
		applicationID,
		name,
		slug,
		isEnabled,
		createdBy,
	); err != nil {
		return nil, handleError(err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(err)
	}

	return &featureFlag, nil
}

func (r *featureFlagRepo) GetMany(ctx context.Context,
	tenantID int64,
	applicationID int64,
) ([]types.FeatureFlag, error) {
	var (
		featureFlags []types.FeatureFlag
		rows         pgx.Rows
		err          error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagGetMany,
		tenantID,
		applicationID,
	); err != nil {
		return nil, handleError(err)
	}

	if featureFlags, err = pgx.CollectRows(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(err)
	}

	return featureFlags, nil
}

func (r *featureFlagRepo) GetOne(ctx context.Context,
	tenantID int64,
	applicationID int64,
	id *int64,
	uuid *string,
	slug *string,
) (*types.FeatureFlag, error) {
	var (
		featureFlag types.FeatureFlag
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.FeatureFlagGetOne,
		tenantID,
		applicationID,
		id,
		uuid,
		slug,
	); err != nil {
		return nil, handleError(err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(err)
	}

	return &featureFlag, err
}

func (r *featureFlagRepo) Update(ctx context.Context,
	tenantID int64,
	applicationID int64,
	id int64,
	name string,
	slug string,
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
		tenantID,
		applicationID,
		id,
		name,
		slug,
		isEnabled,
		modifiedBy,
	); err != nil {
		return nil, handleError(err)
	}

	if featureFlag, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.FeatureFlag],
	); err != nil {
		return nil, handleError(err)
	}

	return &featureFlag, nil
}

func (r *featureFlagRepo) Delete(ctx context.Context,
	tenantID int64,
	applicationID int64,
	id int64,
) error {
	row := r.db.QueryRow(ctx,
		queries.FeatureFlagDelete,
		tenantID,
		applicationID,
		id,
	)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(err)
	}

	if numDeleted < 1 {
		return errors.New("no rows deleted")
	}

	if numDeleted > 1 {
		return errors.New(
			fmt.Sprintf("expected to delete 1 row, deleted %v", numDeleted),
		)
	}

	return nil
}
