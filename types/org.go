package types

import "time"

type Organization struct {
	ID         int64      `json:"id" db:"id"`
	UUID       string     `json:"uuid" db:"uuid"`
	Name       string     `json:"name" db:"name"`
	Slug       string     `json:"slug" db:"slug"`
	Owner      int64      `json:"owner" db:"owner"`
	Created    time.Time  `json:"created" db:"created"`
	CreatedBy  int64      `json:"createdBy" db:"created_by"`
	Modified   *time.Time `json:"modified" db:"modified"`
	ModifiedBy *int64     `json:"modifiedBy" db:"modified_by"`
}
