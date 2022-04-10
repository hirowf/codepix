package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka "github.com/santosant/codepix/application/kafka"
	"github.com/santosant/codepix/infrastructure/db"
	"github.com/spf13/cobra"
	"os"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache kafka",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Production message")
		deliveryChan := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()

		//kafka.Publish("Hello Kafka", "test", producer, deliveryChan)
		go kafka.DeliveryReport(deliveryChan)

		KafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		KafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
