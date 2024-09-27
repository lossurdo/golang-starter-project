package cmd

import (
	"fmt"

	"github.com/lossurdo/golang-starter-project/utils"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A Date usage examples",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true}))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{YYYYMMDD: true, AddTime: true}))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true, AddTimeAndSeconds: true}))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true, DashSepareted: true, AddTimeAndSeconds: true}))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{YYYYMMDD: true, DashSepareted: true, AddTimeAndSeconds: true}))

		fmt.Printf("Time: %s\n", utils.GetTime(utils.DateTimeFormat{AddTime: true}))
		fmt.Printf("Time: %s\n", utils.GetTime(utils.DateTimeFormat{AddTimeAndSeconds: true}))
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
