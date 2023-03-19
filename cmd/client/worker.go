package main

import (
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"github.com/pablogolobaro/chequery/cmd/client/client/check"
	"log"
	"sync"
	"time"
)

type Worker struct {
	client *client.CheckGeneratingAPI
	stopCh chan struct{}
}

func NewWorker(client *client.CheckGeneratingAPI) *Worker {
	stopCh := make(chan struct{})
	return &Worker{client: client, stopCh: stopCh}
}

func (w Worker) Start() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			w.client.Check.GetGenerated(nil)
		case <-w.stopCh:
			return
		}
	}
}

func (w Worker) Stop() {
	close(w.stopCh)
}

func (w Worker) handleChecks() error {
	generated, err := w.client.Check.GetGenerated(nil)
	if err != nil {
		return err
	}

	successIds := make([]int64, 0, len(generated.Payload.IDs))

	var wg sync.WaitGroup
	for _, id := range generated.Payload.IDs {
		wg.Add(1)
		go func(id int64) {
			_, err := w.client.Check.GetPDF(&check.GetPDFParams{CheckID: id}, nil)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Sending to printer check by id №%v", id)
			//log.Printf("Length of file: %v", getPDF.Payload.Header.Size)

			successIds = append(successIds)
		}(id)
	}
	wg.Wait()

	_, err = w.client.Check.UpdateChecksStatus(&check.UpdateChecksStatusParams{IDs: successIds})
	if err != nil {
		return err
	}

	return nil
}
