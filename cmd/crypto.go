package cmd

import (
	"os"

	"github.com/lossurdo/golang-starter-project/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Example of a string cryptography command",
	Run: func(cmd *cobra.Command, args []string) {
		key := os.Getenv("KEY")
		log.Infof("KEY: %v", key)

		enc, _ := cmd.Flags().GetString("enc")
		dec, _ := cmd.Flags().GetString("dec")

		if enc != "" {
			log.Infoln("String encrypted:", utils.Encrypt(enc, key))
		} else if dec != "" {
			log.Infoln("String decrypted:", utils.Decrypt(dec, key))
		} else {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	cryptoCmd.Flags().String("enc", "", "String to encrypt")
	cryptoCmd.Flags().String("dec", "", "String to decrypt")
	rootCmd.AddCommand(cryptoCmd)
}
