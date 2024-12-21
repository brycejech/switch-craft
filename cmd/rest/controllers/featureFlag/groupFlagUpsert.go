package featureflag

import (
	"errors"
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type groupFlagUpdateArgs struct {
	IsEnabled bool `json:"isEnabled"`
}

func (c *featureFlagController) GroupFlagUpsert(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug    = r.PathValue("orgSlug")
		appSlug    = r.PathValue("appSlug")
		flagIDStr  = r.PathValue("flagID")
		flagID     int64
		flag       *types.FeatureFlag
		groupIDStr = r.PathValue("groupID")
		groupID    int64
		group      *types.OrgGroup
		err        error
	)
	for _, val := range []string{orgSlug, appSlug, flagIDStr, groupIDStr} {
		if val == "" {
			restutils.NotFound(w, r)
			return
		}
	}

	if flagID, err = strconv.ParseInt(flagIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	if flag, err = c.core.FeatFlagGetOne(r.Context(),
		c.core.NewFeatFlagGetOneArgs(orgSlug, appSlug, &flagID, nil, nil),
	); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	if groupID, err = strconv.ParseInt(groupIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	if group, err = c.core.OrgGroupGetOne(r.Context(),
		c.core.NewOrgGroupGetOneArgs(orgSlug, &groupID, nil),
	); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	body := &groupFlagUpdateArgs{}
	if err = restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	// Group flag upsert
	if _, err := c.core.GroupFlagGetOne(r.Context(),
		c.core.NewGroupFlagGetOneArgs(
			orgSlug,
			group.ID,
			appSlug,
			flag.ID,
		),
	); err != nil {
		// Did not exist, create
		if errors.Is(err, types.ErrNotFound) {
			var newGroupFlag *types.OrgGroupFeatureFlag
			if newGroupFlag, err = c.core.GroupFlagCreate(r.Context(),
				c.core.NewGroupFlagCreateArgs(
					orgSlug,
					group.ID,
					appSlug,
					flag.ID,
					body.IsEnabled,
				),
			); err != nil {
				restutils.HandleCoreErr(w, r, err)
			} else {
				restutils.Render(w, r, http.StatusOK, newGroupFlag)
			}
			return
		}

		restutils.HandleCoreErr(w, r, err)
		return
	} else {
		// Did exist, update
		var updatedGroupFlag *types.OrgGroupFeatureFlag
		if updatedGroupFlag, err = c.core.GroupFlagUpdate(r.Context(),
			c.core.NewGroupFlagUpdateArgs(
				orgSlug,
				group.ID,
				appSlug,
				flag.ID,
				body.IsEnabled,
			),
		); err != nil {
			restutils.HandleCoreErr(w, r, err)
			return
		}

		restutils.Render(w, r, http.StatusOK, updatedGroupFlag)
	}
}
