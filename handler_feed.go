package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxkaiser11/gator/internal/database" // your module path
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name, url := cmd.Args[0], cmd.Args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	now := time.Now().UTC()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Feed created Successfully:")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:        %s\n", feed.ID)
	fmt.Printf("* Created:   %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:   %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:      %s\n", feed.Name)
	fmt.Printf("* URL:       %s\n", feed.Url)
	fmt.Printf("* UserID:    %s\n", feed.UserID)
}
