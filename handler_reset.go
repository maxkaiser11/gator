package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		fmt.Errorf("Failed to delete usesrs")
	}
	fmt.Println("Successfully deleted all users from the database")
	return nil
}

