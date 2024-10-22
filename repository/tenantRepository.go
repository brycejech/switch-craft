package repository

type tenant struct {
	ID   int64
	Name string
}

func NewTenantRepository() *tenantRepo {
	return &tenantRepo{}
}

type tenantRepo struct{}

func (r *tenantRepo) GetTenants() ([]tenant, error) {
	return []tenant{}, nil
}

func (r *tenantRepo) GetTenant(ID int64) (*tenant, error) {
	return &tenant{}, nil
}

func (r *tenantRepo) CreateTenant(name string) (*tenant, error) {
	return &tenant{}, nil
}

func (r *tenantRepo) UpdateTenant(t *tenant) error {
	return nil
}

func (r *tenantRepo) DeleteTenant(ID int64) error {
	return nil
}
