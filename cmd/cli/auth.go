package cli

import (
	"fmt"
	"log"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

func registerAuthModule(core *core.Core) {
	var authCmd = &cobra.Command{
		Use:   "auth",
		Short: "SwitchCraft CLI auth module",
	}
	authzCmd(core, authCmd)
	authValidateJWTCmd(core, authCmd)
	authHashPasswordCmd(core, authCmd)
	authComparePasswordCmd(core, authCmd)
	authCreateSigningKeyCmd(core, authCmd)

	rootCmd.AddCommand(authCmd)
}

func authzCmd(core *core.Core, parentCmd *cobra.Command) {
	authzCmd := &cobra.Command{
		Use:   "authorize",
		Short: "Create signed JWT for current user",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			jwt, err := core.AuthCreateJWT(authAccount)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(jwt)
		},
	}

	parentCmd.AddCommand(authzCmd)
}

func authValidateJWTCmd(core *core.Core, parentCmd *cobra.Command) {
	var token string
	validateJWTCmd := &cobra.Command{
		Use:   "validateJWT",
		Short: "Validate a signed JWT",
		Run: func(_ *cobra.Command, _ []string) {
			mustAuthn(core)

			account, err := core.AuthValidateJWT(token)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("JWT is valid for user '%s'\n", account.Username)
		},
	}
	validateJWTCmd.Flags().StringVar(&token, "token", "", "The JWT token to check")
	validateJWTCmd.MarkFlagRequired("token")

	parentCmd.AddCommand(validateJWTCmd)
}

func authHashPasswordCmd(core *core.Core, parentCmd *cobra.Command) {
	var password string
	hashPasswordCmd := &cobra.Command{
		Use:   "hashPassword",
		Short: "Hash a password",
		Run: func(_ *cobra.Command, _ []string) {
			hash, err := core.AuthPasswordHash(password)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(hash)
		},
	}
	hashPasswordCmd.Flags().StringVar(&password, "password", "", "The password to hash")
	hashPasswordCmd.MarkFlagRequired("password")

	parentCmd.AddCommand(hashPasswordCmd)
}

func authComparePasswordCmd(core *core.Core, parentCmd *cobra.Command) {
	var password string
	var hash string
	comparePasswordCmd := &cobra.Command{
		Use:   "comparePassword",
		Short: "Compare a password to a hash",
		Run: func(_ *cobra.Command, _ []string) {
			ok, err := core.AuthPasswordCheck(password, hash)
			if err != nil {
				log.Fatal(err)
			}
			if ok {
				fmt.Println("Passwords match")
			} else {
				fmt.Println("Passwords do not match")
			}
		},
	}
	comparePasswordCmd.Flags().StringVar(&password, "password", "", "The password to compare")
	comparePasswordCmd.MarkFlagRequired("password")
	comparePasswordCmd.Flags().StringVar(&hash, "hash", "", "The hashed password to compare")
	comparePasswordCmd.MarkFlagRequired("hash")

	parentCmd.AddCommand(comparePasswordCmd)
}

func authCreateSigningKeyCmd(core *core.Core, parentCmd *cobra.Command) {
	var keyLen uint32
	createSigningKeyCmd := &cobra.Command{
		Use:   "createSigningKey",
		Short: "Create a cryptographically secure signing key encoded to hexadecimal",
		Run: func(_ *cobra.Command, _ []string) {
			key, err := core.AuthCreateSigningKey(keyLen)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(key)
		},
	}
	createSigningKeyCmd.Flags().Uint32Var(&keyLen, "bitLength", 0, "Key length in bits. Must be >= 256 and divisible by 8")
	createSigningKeyCmd.MarkFlagRequired("bitLength")

	parentCmd.AddCommand(createSigningKeyCmd)
}
