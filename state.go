package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"

	"github.com/maxkaiser11/gator/internal/config"
	"github.com/maxkaiser11/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
} 

type commands struct {
	registeredCommands map[string]func(*state, command) error
}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user %q: %w", name, err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User has been set")
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		fmt.Errorf("Failed to delete usesrs")
	}
	fmt.Println("Successfully deleted all users from the database")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to fetch users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

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

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f

}
