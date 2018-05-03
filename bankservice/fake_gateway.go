package bankservice

import (
	"context"
	"log"

	"github.com/ivanmatyash/bank-golang/api"
)

func (s *bankServer) FakeGateway(ctx context.Context, in *api.RequestClient) (*api.ResponseClient, error) {

	client := in.Req

	if client.Phone != "" {
		log.Printf("SMS for client %s", client.Name)
	}

	if client.Email != "" {
		log.Printf("EMAIL for client %s", client.Name)
	}
	return nil, nil
}
