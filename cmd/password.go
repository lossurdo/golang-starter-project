package cmd

import (
	"github.com/lossurdo/golang-starter-project/utils"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Example of password generator",
	Run: func(cmd *cobra.Command, args []string) {
		fullname, _ := cmd.Flags().GetString("fullname")
		simple, _ := cmd.Flags().GetBool("simple")
		size, _ := cmd.Flags().GetInt("size")

		if fullname != "" {
			log.Infof("Username: %s", utils.GenerateUsername(fullname))
		} else if simple {
			log.Infof("Simple password: %s", utils.GenerateSimplePassword())
		} else {
			log.Infof("Strong password: %s // %s",
				utils.GeneratePassword(size, false), utils.GeneratePassword(size, true))
		}
	},
}

func init() {
	passwordCmd.Flags().String("fullname", "", "Generate username based on fullname")
	passwordCmd.Flags().Bool("simple", false, "Simple password (default: strong)")
	passwordCmd.Flags().Int("size", 12, "Password size (default: 12)")
	rootCmd.AddCommand(passwordCmd)
}
