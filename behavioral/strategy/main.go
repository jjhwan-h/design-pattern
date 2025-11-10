package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	hc := &http.Client{Timeout: 5 * time.Second}

	cli := Client{
		Doer:     hc,
		Strategy: NewExponentialBackoff(),
	}

	ctx, cacnel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cacnel()

	req, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	resp, err := cli.DoWithRetry(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Status)

	cli.SetRetryStrategy(NewJitterBackoff())

	ctx, cacnel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cacnel()

	req, _ = http.NewRequest(http.MethodGet, "localhost:8080", nil)
	resp, err = cli.DoWithRetry(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
}
