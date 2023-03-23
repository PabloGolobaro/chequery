package main

import (
	"bytes"
	"context"
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

	ticker := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-ticker.C:
			err := w.handleChecks()
			if err != nil {
				log.Println(err)
			}

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
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, id := range generated.Payload.IDs {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			buffer := &bytes.Buffer{}

			reqParam := &check.GetPDFParams{CheckID: id}
			_, err := w.client.Check.GetPDF(reqParam.WithTimeout(time.Second*5), buffer)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Sending to printer check by id №%v", id)
			log.Printf("Length of file: %v", buffer.Len())
			mu.Lock()
			successIds = append(successIds, id)
			mu.Unlock()
		}(id)
	}
	wg.Wait()
	if len(successIds) == 0 {

		return nil
	}
	_, err = w.client.Check.UpdateChecksStatus(&check.UpdateChecksStatusParams{IDs: successIds, Context: context.Background()})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
