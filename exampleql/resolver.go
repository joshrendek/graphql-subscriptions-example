package exampleql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"math/rand"
	"time"
)

type Resolver struct{}

func New() Config {
	return Config{
		Resolvers: &Resolver{},
	}
}

func (r *subscriptionResolver) ExampleAdded(ctx context.Context) (<-chan *Example, error) {
	msgs := make(chan *Example, 1)

	go func() {
		for {
			msgs <- &Example{Message: randString(50)}
			time.Sleep(1 * time.Second)
		}
	}()
	return msgs, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) *string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	s := string(b)
	return &s
}
