package types

import "time"

type OrgGroupFeatureFlag struct {
	OrgID      int64      `json:"orgId" db:"org_id"`
	GroupID    int64      `json:"groupId" db:"group_id"`
	AppID      int64      `json:"appId" db:"application_id"`
	FlagID     int64      `json:"flagId" db:"flag_id"`
	IsEnabled  bool       `json:"isEnabled" db:"is_enabled"`
	Created    time.Time  `json:"created" db:"created"`
	CreatedBy  int64      `json:"createdBy" db:"created_by"`
	Modified   *time.Time `json:"modified" db:"modified"`
	ModifiedBy *int64     `json:"modifiedBy" db:"modified_by"`
}
