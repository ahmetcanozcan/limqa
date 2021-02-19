package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/ahmetcanozcan/limqa"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "consumes a message",
	Long: `Consumes a message from a queue in AMQP server using limqa.Consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		b := base()
		c,err := limqa.NewConsumer(b,*queueF,*exchangeF,limqa.DeclareExchange(true))
		cobra.CheckErr(err)
		if *timeoutF != "none" {
			d,err := time.ParseDuration(*timeoutF)
			cobra.CheckErr(err)
			go func(){
				time.Sleep(d)
				cobra.CheckErr(errors.New("Timeout exceed"))
			}()
		}
		m := c.Consume()
		fmt.Println(string(m))
	},
}



func init() {
	rootCmd.AddCommand(consumeCmd)
	queueF=consumeCmd.Flags().StringP("queue", "q", "", "Name of the queue which the message would be consumed")
	timeoutF=consumeCmd.Flags().StringP("timeout", "t", "10s", "Timeout for consume request. use 'none' for no timeout")
}
