package core

func NewCore(repository Repository) *Core {
	return &Core{
		repository: repository,
	}
}

type Core struct {
	repository Repository
}

func (c *Core) MigrateUp() error {
	return c.repository.MigrateUp()
}

func (c *Core) MigrateDown() error {
	return c.repository.MigrateDown()
}

/* === PORTS === */

type Repository interface {
	MigrateUp() error
	MigrateDown() error
}
