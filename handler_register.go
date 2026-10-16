package main

import (
	"fmt"
	"context"
	"time"


	"github.com/google/uuid"
	"github.com/maxkaiser11/gator/internal/database" // your module path
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(), 
		UpdatedAt: time.Now(), 
		Name: name, 
	})
	if err != nil {
		return fmt.Errorf("couldnt create user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Printf("User created: %+v\n", user)
	return nil
}

