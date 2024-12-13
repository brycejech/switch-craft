package types

import "time"

type OrgGroup struct {
	OrgID       int64      `json:"orgId" db:"org_id"`
	ID          int64      `json:"id" db:"id"`
	UUID        string     `json:"uuid" db:"uuid"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Created     time.Time  `json:"created" db:"created"`
	CreatedBy   int64      `json:"createdBy" db:"created_by"`
	Modified    *time.Time `json:"modified" db:"modified"`
	ModifiedBy  *int64     `json:"modifiedBy" db:"modified_by"`
}

type OrgGroupAccount struct {
	OrgID     int64     `json:"orgId" db:"org_id"`
	GroupID   int64     `json:"groupId" db:"group_id"`
	ID        int64     `json:"id" db:"id"`
	AccountID int64     `json:"accountId" db:"account_id"`
	Created   time.Time `json:"created" db:"created"`
	CreatedBy int64     `json:"createdBy" db:"created_by"`
}
