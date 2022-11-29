package scamdetector

import (
    "context"
    "github.com/demimurg/twitter/internal/usecase"
    "strings"
)

func NewDummyClient() usecase.ScamDetectorClient {
    return dummyClient{}
}

type dummyClient struct {}

func (d dummyClient) CheckEmail(_ context.Context, email string) error {
    email = strings.ToLower(email)
    if strings.Contains(email, "donald") && strings.Contains(email, "trump") {
        return usecase.ErrFakeEmail
    }
    return nil
}