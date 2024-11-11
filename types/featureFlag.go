package types

import "time"

type FeatureFlag struct {
	TenantID      int64      `json:"tenantId" db:"tenant_id"`
	ApplicationID int64      `json:"applicationId" db:"application_id"`
	ID            int64      `json:"id" db:"id"`
	UUID          string     `json:"uuid" db:"uuid"`
	Name          string     `json:"name" db:"name"`
	Slug          string     `json:"slug" db:"slug"`
	IsEnabled     bool       `json:"isEnabled" db:"is_enabled"`
	Created       time.Time  `json:"created" db:"created"`
	CreatedBy     int64      `json:"createdBy" db:"created_by"`
	Modified      *time.Time `json:"modified" db:"modified"`
	ModifiedBy    *int64     `json:"modifiedBy" db:"modified_by"`
}
