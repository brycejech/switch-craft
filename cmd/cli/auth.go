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
	rootCmd.AddCommand(authCmd)

	/* ----------------------------- */
	/* === HASH PASSWORD COMMAND === */
	/* ----------------------------- */
	var hashPasswordCmdPass string
	var hashPasswordCmd = &cobra.Command{
		Use:   "hashPassword",
		Short: "Hash a password",
		Run: func(_ *cobra.Command, _ []string) {
			hash, err := core.AuthPasswordHash(hashPasswordCmdPass)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(hash)
		},
	}
	hashPasswordCmd.Flags().StringVar(&hashPasswordCmdPass, "password", "", "The password to hash")
	hashPasswordCmd.MarkFlagRequired("password")
	authCmd.AddCommand(hashPasswordCmd)

	/* -------------------------------- */
	/* === COMPARE PASSWORD COMMAND === */
	/* -------------------------------- */
	var comparePasswordCmdPass string
	var comparePasswordCmdHash string
	var comparePasswordCmd = &cobra.Command{
		Use:   "comparePassword",
		Short: "Compare a password to a hash",
		Run: func(_ *cobra.Command, _ []string) {
			ok, err := core.AuthPasswordCheck(comparePasswordCmdPass, comparePasswordCmdHash)
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
	comparePasswordCmd.Flags().StringVar(&comparePasswordCmdPass, "password", "", "The password to compare")
	comparePasswordCmd.MarkFlagRequired("password")
	comparePasswordCmd.Flags().StringVar(&comparePasswordCmdHash, "hash", "", "The hashed password to compare")
	comparePasswordCmd.MarkFlagRequired("hash")
	authCmd.AddCommand(comparePasswordCmd)

	/* ---------------------------------- */
	/* === CREATE SIGNING KEY COMMAND === */
	/* ---------------------------------- */
	var createSigningKeyCmdKeyLength uint32
	var createSigningKeyCmd = &cobra.Command{
		Use:   "createSigningKey",
		Short: "Create a cryptographically secure signing key encoded to hexadecimal",
		Run: func(_ *cobra.Command, _ []string) {
			key, err := core.AuthCreateSigningKey(createSigningKeyCmdKeyLength)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(key)
		},
	}
	createSigningKeyCmd.Flags().Uint32Var(&createSigningKeyCmdKeyLength, "bitLength", 0, "Key length in bits. Must be >= 256 and divisible by 8")
	createSigningKeyCmd.MarkFlagRequired("bitLength")
	authCmd.AddCommand(createSigningKeyCmd)
}
