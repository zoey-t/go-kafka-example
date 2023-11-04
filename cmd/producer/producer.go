package main

import (
	"fmt"
	"go-kafka-example/config"
	"go-kafka-example/pkg/controllers"
	"go-kafka-example/pkg/models"
	"log"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
)

func setUpProdcuer() (sarama.SyncProducer, error) {
	saramConfig := sarama.NewConfig()
	saramConfig.Producer.Return.Successes = true // ensures that producer receives an acknowledgement once the message is successfully stored in the topics
	producer, err := sarama.NewSyncProducer([]string{config.KafkaServerAddr}, saramConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to set up producer: %w", err)
	}
	return producer, nil
}

func main() {
	// kafka producer
	producer, err := setUpProdcuer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}

	defer producer.Close()

	fmt.Printf("Kafka PRODUCER 📨 started at http://localhost%s\n",
		config.ProducerPort)

	// routes
	e := echo.New()
	e.GET("/addresses/:address", controllers.GetAddress(models.AddressDB))
	e.GET("/addresses", controllers.GetAddresses(models.AddressDB))
	e.POST("/addresses/:address", controllers.CreateAddress(producer, models.AddressDB))
	e.PUT("/addresses/:address", controllers.UpdateAddress(producer, models.AddressDB))
	e.DELETE("/addresses/:address", controllers.DeleteAddress(producer, models.AddressDB))

	// LabelRoutes(e)
	e.GET("/labels/:label", controllers.GetLabel(models.LabelDB))
	e.GET("/labels", controllers.GetLabels(models.LabelDB))

	// TransactionRoutes(e)
	e.GET("/transactions/:hash", controllers.GetTransaction(models.TransactionDB))
	e.GET("/transactions", controllers.GetTransactions(models.TransactionDB))
	e.POST("transactions/:hash", controllers.CreateTransaction(producer, models.TransactionDB))
	e.PUT("transactions/:hash", controllers.UpdateTransaction(producer, models.TransactionDB))
	e.DELETE("transactions/:hash", controllers.DeleteTransaction(producer, models.TransactionDB))
	e.Logger.Fatal(e.Start(config.ProducerPort))

}
