package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kafka called")
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
