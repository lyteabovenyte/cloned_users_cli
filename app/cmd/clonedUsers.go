package cmd

import (
	// #include<stdio.h>
	// #include<utmpx.h>
	// #include <utmpx.h>
	// #include <string.h>
	"C"

	"github.com/spf13/cobra"
)
import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/stephane-martin/skewer/sys/utmpx"
)

// clonedUsersCmd represents the clonedUsers command
var clonedUsersCmd = &cobra.Command{
	Use:   "clonedUsers",
	Short: "A Simple command line prompt to read /var/run/utmpx file and show the users logged in to the system",
	Long: `clonedUsers is just a simple command line prompt that reads /var/run/utmpx file
	which determines the name of the users and host from where they are connected to the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := cmd.Flags()

		for _, ent := range utmpx.All() {
			switch {
			case mustBool(fs, "user") && mustBool(fs, "timestamp"):
				fmt.Printf("User: %v, Timestamp: %v\n", ent.User, ent.Timestamp)
			case mustBool(fs, "user") && !mustBool(fs, "timestamp"):
				fmt.Printf("User: %v\n", ent.User)
			case !mustBool(fs, "user") && mustBool(fs, "timestamp"):
				fmt.Printf("Timestamp: %v\n", ent.Timestamp)
			default:
				fmt.Printf("User: %v, Timestamp: %v, ID: %v\n", ent.User, ent.Timestamp, ent.ID)
			}
		}
	},
}

func mustString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}
	return v
}

func mustBool(fs *pflag.FlagSet, name string) bool {
	v, err := fs.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	rootCmd.AddCommand(clonedUsersCmd)

	clonedUsersCmd.Flags().BoolP("user", "u", true,
		"get the logged in users")

	clonedUsersCmd.Flags().Bool("timestamp", false,
		"get the user's timestamp")

	clonedUsersCmd.Flags().BoolP("help", "h", false,
		"display help text")
}
