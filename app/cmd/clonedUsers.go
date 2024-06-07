/*
Copyright © 2024 Amir Alaeifar lyteabovenyte@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clonedUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clonedUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
