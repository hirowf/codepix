package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/santosant/codepix/application/grpc"
	"github.com/santosant/codepix/application/kafka"
	"github.com/santosant/codepix/infrastructure/db"
	"os"

	"github.com/spf13/cobra"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and Kafka Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, portNumber)

		deliveryChan := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()

		go kafka.DeliveryReport(deliveryChan)
		KafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		KafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 50051, "gRPC Port")
}
