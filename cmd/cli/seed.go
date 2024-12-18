package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"switchcraft/core"
	"switchcraft/types"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func registerSeedModule(core *core.Core) {
	seedCmd := &cobra.Command{
		Use:   "seed",
		Short: "SwitchCraft CLI database seeder",
	}
	seedAllCmd(core, seedCmd)

	rootCmd.AddCommand(seedCmd)
}

func seedAllCmd(core *core.Core, parentCmd *cobra.Command) {
	var dataFile string
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Seed all tables with test data",
		Run: func(cmd *cobra.Command, _ []string) {
			seed := mustParseSeedFile(dataFile)

			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				seedOrganizations(wg, core, seed.Organizations)
			}()
			wg.Wait()
		},
	}
	allCmd.Flags().StringVar(&dataFile, "dataFile", "", "Path to json seed file")
	allCmd.MarkFlagRequired("dataFile")

	parentCmd.AddCommand(allCmd)
}

func seedOrganizations(wg *sync.WaitGroup, core *core.Core, seedOrgs []seedOrganization) {
	wg.Add(len(seedOrgs))
	for _, seedOrg := range seedOrgs {
		defer wg.Done()

		account, err := core.Signup(context.Background(), core.NewOrgAccountSignupArgs(
			seedOrg.Owner.FirstName,
			seedOrg.Owner.LastName,
			seedOrg.Owner.Email,
			seedOrg.Owner.Username,
			os.Getenv("SWITCHCRAFT_SEED_PASS"),
		))
		if err != nil {
			fmt.Printf(
				"error creating owner '%s' for org '%s' - %s\n",
				seedOrg.Owner.Username,
				seedOrg.Name,
				err,
			)
			continue
		}
		fmt.Printf("Account created - '%s'\n", account.Username)

		initialCtx := types.NewOperationCtx(context.Background(), "", time.Now(), *account)

		org, err := core.OrgCreate(
			initialCtx,
			core.NewOrgCreateArgs(
				seedOrg.Name,
				seedOrg.Slug,
				account.ID,
			),
		)
		if err != nil {
			fmt.Printf(
				"error creating org '%s' - %s",
				seedOrg.Name,
				err,
			)
			continue
		}
		fmt.Printf("Organization created - '%s'\n", org.Name)

		owner, err := core.OrgAccountSetOrgID(initialCtx, org.ID, account.ID)
		if err != nil {
			fmt.Printf(
				"error setting owner '%s' orgID (orgName: %s, orgID: %v) - %s",
				account.Username,
				org.Name,
				org.ID,
				err,
			)
			continue
		}
		fmt.Printf("Set org '%s' owner's orgID - '%s' orgID '%v'\n", org.Name, owner.Username, *owner.OrgID)

		opCtx := types.NewOperationCtx(context.Background(), "", time.Now(), *owner)

		wg.Add(1)
		go func() {
			defer wg.Done()
			seedApplications(wg, core, opCtx, org.Slug, seedOrg.Applications)
		}()

		// Ensure orgAccounts are created before attempting to create groups
		// and populate with members
		wg.Add(1)
		go func() {
			defer wg.Done()

			acctWg := &sync.WaitGroup{}
			acctWg.Add(1)
			go func() {
				defer acctWg.Done()
				seedOrgAccounts(acctWg, core, opCtx, org.Slug, seedOrg.Accounts)
			}()
			acctWg.Wait()

			seedOrgGroups(wg, core, opCtx, org.Slug, seedOrg.Groups)
		}()
	}
}

func seedOrgAccounts(
	wg *sync.WaitGroup,
	core *core.Core,
	ctx context.Context,
	orgSlug string,
	seedAccounts []seedAccount,
) {
	wg.Add(len(seedAccounts))
	for _, seedAccount := range seedAccounts {
		defer wg.Done()
		account, err := core.OrgAccountCreate(ctx,
			core.NewOrgAccountCreateArgs(
				orgSlug,
				seedAccount.FirstName,
				seedAccount.LastName,
				seedAccount.Email,
				seedAccount.Username,
				nil,
			),
		)
		if err != nil {
			fmt.Printf(
				"error creating account '%s' for org '%s' - %s\n",
				seedAccount.Username,
				orgSlug,
				err,
			)
			continue
		}

		fmt.Printf("OrgAccount created - '%s'\n", account.Username)
	}
}

func seedOrgGroups(
	wg *sync.WaitGroup,
	core *core.Core,
	ctx context.Context,
	orgSlug string,
	seedOrgGroups []seedGroup,
) {
	wg.Add(len(seedOrgGroups))
	for _, seedGroup := range seedOrgGroups {
		defer wg.Done()
		group, err := core.OrgGroupCreate(ctx,
			core.NewOrgGroupCreateArgs(
				orgSlug,
				seedGroup.Name,
				seedGroup.Description,
			),
		)
		if err != nil {
			fmt.Printf(
				"error creating orgGroup '%s' - %s",
				seedGroup.Name,
				err,
			)
			continue
		}
		fmt.Printf("OrgGroup created - '%s'\n", group.Name)

		addAccWg := &sync.WaitGroup{}
		addAccWg.Add(len(seedGroup.Members))
		for _, username := range seedGroup.Members {
			go func() {
				defer addAccWg.Done()
				account, err := core.OrgAccountGetOne(ctx,
					core.NewOrgAccountGetOneArgs(orgSlug, nil, nil, &username),
				)
				if err != nil {
					fmt.Printf(
						"error locating account '%s' - %s\n",
						username,
						err,
					)
					return
				}

				if _, err := core.OrgGroupAccountAdd(ctx,
					core.NewOrgGroupAccountAddArgs(
						orgSlug,
						group.ID,
						account.ID,
					),
				); err != nil {
					fmt.Printf(
						"error adding account '%s' to group '%s' - %s\n",
						username,
						group.Name,
						err,
					)
					return
				}
				fmt.Printf(
					"Added account '%s' to group '%s'\n",
					account.Username,
					group.Name,
				)
			}()
		}
		addAccWg.Wait()
	}
}

func seedApplications(
	wg *sync.WaitGroup,
	core *core.Core,
	ctx context.Context,
	orgSlug string,
	seedApps []seedApplication,
) {
	wg.Add(len(seedApps))
	for _, seedApp := range seedApps {
		defer wg.Done()
		app, err := core.AppCreate(ctx,
			core.NewAppCreateArgs(
				orgSlug,
				seedApp.Name,
				seedApp.Slug,
			),
		)
		if err != nil {
			fmt.Printf(
				"error creating app '%s' for org '%s' - %s\n",
				seedApp.Name,
				orgSlug,
				err,
			)
			continue
		}
		fmt.Printf("Application created - '%s'\n", app.Name)

		wg.Add(1)
		go func() {
			defer wg.Done()
			seedFeatureFlags(wg, core, ctx, orgSlug, seedApp.Slug, seedApp.FeatureFlags)
		}()
	}

}

func seedFeatureFlags(
	wg *sync.WaitGroup,
	core *core.Core,
	ctx context.Context,
	orgSlug string,
	appSlug string,
	seedFlags []seedFeatureFlag,
) {
	wg.Add(len(seedFlags))
	for _, seedFlag := range seedFlags {
		defer wg.Done()
		flag, err := core.FeatFlagCreate(ctx,
			core.NewFeatFlagCreateArgs(
				orgSlug,
				appSlug,
				seedFlag.Name,
				seedFlag.Label,
				seedFlag.Description,
				seedFlag.IsEnabled,
			),
		)
		if err != nil {
			fmt.Printf(
				"error creating feature flag '%s' for app '%s' - %s\n",
				seedFlag.Name,
				appSlug,
				err,
			)
			continue
		}
		fmt.Printf("Feature flag created - '%s'\n", flag.Name)
	}
}

type seed struct {
	Organizations []seedOrganization `json:"organizations"`
}

type seedOrganization struct {
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	Owner        seedAccount       `json:"owner"`
	Accounts     []seedAccount     `json:"accounts"`
	Groups       []seedGroup       `json:"groups"`
	Applications []seedApplication `json:"applications"`
}

type seedAccount struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type seedGroup struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []string `json:"members"`
}

type seedApplication struct {
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	FeatureFlags []seedFeatureFlag `json:"featureFlags"`
}

type seedFeatureFlag struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"isEnabled"`
}

func mustParseSeedFile(filepath string) seed {
	seedFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var seed seed
	if err = json.Unmarshal([]byte(seedFile), &seed); err != nil {
		log.Fatal(err)
	}

	return seed
}
