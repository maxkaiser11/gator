package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching url: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

