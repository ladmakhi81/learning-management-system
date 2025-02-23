package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	queuedto "github.com/ladmakhi81/learning-management-system/internal/queue/dto"
	queueservice "github.com/ladmakhi81/learning-management-system/internal/queue/service"
	pkgrabbitmqclient "github.com/ladmakhi81/learning-management-system/pkg/rabbitmq"
	"github.com/spf13/viper"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/model/optimize"
)

func main() {
	viper.AutomaticEnv()
	config := baseconfig.NewConfig()
	if err := config.LoadConfig(); err != nil {
		fmt.Errorf("environment variable is not loaded : %v", err)

		return
	}
	rabbitmqClient, rabbitmqClientErr := pkgrabbitmqclient.NewRabbitmqClient(config.RabbitmqClientURL)
	if rabbitmqClientErr != nil {
		fmt.Errorf("rabbitmq client not connected : %v", rabbitmqClientErr)
		return
	}
	pdfQueueService, pdfQueueServiceErr := queueservice.NewPDFQueueService(rabbitmqClient)
	if pdfQueueServiceErr != nil {
		fmt.Errorf("pdf queue service not load : %v", pdfQueueServiceErr)

		return
	}

	if err := license.SetMeteredKey(config.UniPDFApiKey); err != nil {
		fmt.Errorf("license error unipdf : %v", err)

		return
	}
	messages, messagesErr := pdfQueueService.QueueService.Receive()
	if messagesErr != nil {
		fmt.Errorf("unable to receive message from pdf queue service : %v", messagesErr)

		return
	}
	fmt.Println("PDF Service Ready")
	for msg := range messages {
		body := new(queuedto.PDFCompressMessage)
		fmt.Println(body)
		if err := json.Unmarshal(msg.Body, body); err != nil {
			fmt.Println("Unable to receive data")
			continue
		}
		compressPDF(body.FileName, body.FileDestination)
		if err := msg.Ack(false); err != nil {
			fmt.Println("Unable to ack")
			continue
		}
	}
}

func compressPDF(filename string, fileDest string) error {
	fileURL := path.Join(fileDest, filename)

	file, fileErr := os.Open(fileURL)
	if fileErr != nil {
		return fmt.Errorf("Can't Open File : %v", fileErr)
	}
	defer file.Close()

	pdfReader, pdfReaderErr := model.NewPdfReader(file)
	if pdfReaderErr != nil {
		return fmt.Errorf("Can't Load PDF : %v", pdfReaderErr)
	}

	pdfWriter, pdfWriterErr := pdfReader.ToWriter(nil)
	if pdfWriterErr != nil {
		return fmt.Errorf("PDF Writer Error : %v", pdfWriterErr)
	}
	pdfWriter.SetOptimizer(optimize.New(optimize.Options{
		CombineDuplicateDirectObjects:   true,
		CombineIdenticalIndirectObjects: true,
		CombineDuplicateStreams:         true,
		CompressStreams:                 true,
		UseObjectStreams:                true,
		ImageQuality:                    60,
		ImageUpperPPI:                   100,
		CleanUnusedResources:            true,
	}))
	if writeErr := pdfWriter.WriteToFile(fileURL); writeErr != nil {
		return fmt.Errorf("Can't Write PDF : %v", writeErr)
	}
	return nil
}
