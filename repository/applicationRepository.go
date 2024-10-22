package repository

func NewApplicationRepository() *applicationRepo {
	return &applicationRepo{}
}

type applicationRepo struct{}

type application struct {
	ID       int64
	TenantID int64
	Name     string
}

func (r *applicationRepo) GetApplications() ([]application, error) {
	return []application{}, nil
}

func (r *applicationRepo) GetApplication(ID int64) (*application, error) {
	return &application{}, nil
}

func (r *applicationRepo) CreateApplication(tenantID int64, name string) (*application, error) {
	return &application{}, nil
}

func (r *applicationRepo) UpdateApplication(f *application) error {
	return nil
}

func (r *applicationRepo) DeleteApplication(ID int64) error {
	return nil
}
