package cmd

import (
	"github.com/ahmetcanozcan/limqa"
	"github.com/spf13/cobra"
)

// Flags
var (
	addressF *string
	messageF *string
	queueF *string
	exchangeF *string
	timeoutF *string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "CLI tool of limqa",
	Long: `limqa-cli is a cli for producing or consuming messages 
from a AMQP server using github.com/ahmetcanozcan/limqa.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	addressF = rootCmd.PersistentFlags().StringP("uri", "u", "amqp://guest:guest@localhost:5672", "Address of the AMQP server")
	exchangeF = rootCmd.PersistentFlags().StringP("exchange", "e", "default_exchange", "name of the target exchange")
}


func base() *limqa.Base{
	base:= limqa.New()
	cobra.CheckErr(base.Connect(*addressF))
	return base
}