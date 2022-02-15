//go:generate go run github.com/99designs/gqlgen generate
package graphql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type Resolver struct {
	tasks *aztables.Client
}

func NewResolver(ctx context.Context, connectionString string) (*Resolver, error) {
	svc, err := aztables.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}

	tasks, err := svc.CreateTable(ctx, "tasks", nil)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		tasks: tasks,
	}, nil
}
