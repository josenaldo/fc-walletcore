package handler

import (
	"log"
	"sync"

	"github.com/josenaldo/fc-walletcore/pkg/events"
	"github.com/josenaldo/fc-walletcore/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	log.Print("Iniciano TransactionCreatedKafkaHandler")

	defer wg.Done()

	log.Print("Publicando mensagem no Kafka")
	h.Kafka.Publish(message, nil, "transactions")

	log.Print("Mensagem publicada no Kafka")
	log.Print("TransactionCreatedKafkaHandler:", message.GetPayload())
}
