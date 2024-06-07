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
	"log"
	"strconv"
	"time"
	"unsafe"

	"github.com/stephane-martin/skewer/sys/utmpx"

	"github.com/spf13/cobra"
)
import (
	"fmt"
	"os"
)

// clonedUsersCmd represents the clonedUsers command
var clonedUsersCmd = &cobra.Command{
	Use:   "clonedUsers",
	Short: "A Simple command line prompt to read /var/run/utmpx file and show the users logged in to the system",
	Long: `clonedUsers is just a simple command line prompt that reads /var/run/utmpx file
	which determines the name of the users and host from where they are connected to the system`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.ReadFile("/var/run/utmpx")
		if err != nil {
			log.Fatal(err)
		}
		var entry utmpx.Entry
		for i := 0; i < len(f); i += unsafe.Sizeof(entry) {
			var copyEntry []utmpx.Entry
			copy(copyEntry, entry)
			copy(&copyEntry, f[i:i+unsafe.Sizeof(copyEntry)])
			fmt.Printf("User: %s, Host: %s, PID: %d, Line: %s, ID: %s, Time: %v\n",
				C.GoString(&copyEntry.UtUser[0]),
				C.GoString(&copyEntry.UtHost[0]),
				copyEntry.Pid,
				C.GoString(&copyEntry.UtLine[0]),
				C.GoString(&copyEntry.UtId[0]),
				time.Unix(int64(copyEntry.Tv.TvSec), 0))
		}

		// ent := utmpx.Entry{}
	},
}

// spliting a byte to specified chunks of byte for each block of user.
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func ByteToInt(in []byte) int32 {
	x, _ := strconv.Atoi(string(in))
	return int32(x)
}

func init() {
	rootCmd.AddCommand(clonedUsersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clonedUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clonedUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
