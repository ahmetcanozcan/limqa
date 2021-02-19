package cmd

import (
	"fmt"

	"github.com/ahmetcanozcan/limqa"
	"github.com/spf13/cobra"
)

var (
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produces a message",
	Long: `Produces a message to an exchange on AMQP server using limqa.Producer`,
	Run: func(cmd *cobra.Command, args []string) {
		b := base()
		p,err := limqa.NewProducer(b,*exchangeF)
		cobra.CheckErr(err)
		cobra.CheckErr(p.Produce([]byte(*messageF)))
		fmt.Println("Message sent")
	},
}

func init() {
	rootCmd.AddCommand(produceCmd)
	messageF = produceCmd.Flags().StringP("message", "m", "", "the message that's sent to AMQP server")
}
