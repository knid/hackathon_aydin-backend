package clients

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)


type OllamaCLient struct {
	Addr string
	Model string
}

func (c OllamaCLient) CheckConnection() error {
	resp, err := http.Get(c.Addr)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(body) != "Ollama is running" {
		return errors.New("ollama instance not running")
	}

	return nil
}

func (c OllamaCLient) SendChat(ctx context.Context, msg string, respFunc func(api.GenerateResponse) error) error {
	client := api.NewClient(&url.URL{Scheme: "http", Host: "10.8.0.9:11434"}, http.DefaultClient)
	
	req := &api.GenerateRequest{
		Model: c.Model,
        Prompt: msg,
    }

	err := client.Generate(ctx, req, respFunc)
	fmt.Println("FINISHED")
	return err
}
