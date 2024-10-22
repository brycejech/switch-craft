package repository

func NewFeatureFlagRepository() *featureFlagRepo {
	return &featureFlagRepo{}
}

type featureFlagRepo struct{}

type featureFlag struct {
	ID            int64
	TenantID      int64
	ApplicationID int64
	Name          string
	Enabled       bool
}

func (r *featureFlagRepo) GetFeatureFlags() ([]featureFlag, error) {
	return []featureFlag{}, nil
}

func (r *featureFlagRepo) GetFeatureFlag(ID int64) (*featureFlag, error) {
	return &featureFlag{}, nil
}

func (r *featureFlagRepo) CreateFeatureFlag(name string) (*featureFlag, error) {
	return &featureFlag{}, nil
}

func (r *featureFlagRepo) UpdateFeatureFlag(f *featureFlag) error {
	return nil
}

func (r *featureFlagRepo) DeleteFeatureFlag(ID int64) error {
	return nil
}
