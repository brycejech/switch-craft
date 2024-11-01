package types

import "time"

type Account struct {
	TenantID   int64      `json:"tenantId" db:"tenant_id"`
	ID         int64      `json:"id" db:"id"`
	UUID       string     `json:"uuid" db:"uuid"`
	FirstName  string     `json:"firstName" db:"first_name"`
	LastName   string     `json:"lastName" db:"last_name"`
	Email      string     `json:"email" db:"email"`
	Username   string     `json:"username" db:"username"`
	Created    time.Time  `json:"created" db:"created"`
	CreatedBy  int64      `json:"createdBy" db:"created_by"`
	Modified   *time.Time `json:"modified" db:"modified"`
	ModifiedBy *int64     `json:"modifiedBy" db:"modified_by"`
}

type Tenant struct {
	ID         int64      `json:"id" db:"id"`
	UUID       string     `json:"uuid" db:"uuid"`
	Name       string     `json:"name" db:"name"`
	Slug       string     `json:"slug" db:"slug"`
	Created    time.Time  `json:"created" db:"created"`
	CreatedBy  int64      `json:"createdBy" db:"created_by"`
	Modified   *time.Time `json:"modified" db:"modified"`
	ModifiedBy *int64     `json:"modifiedBy" db:"modified_by"`
}
