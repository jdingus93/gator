package main

import (
	"context"
	"errors"
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func resetCommandHandler(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Println("Failed to reset database:", err)
		return err
	}
	fmt.Println("Database reset successfully")
	return nil
}

func usersCommand(s *state, cmd command) error {
	ctx := context.Background()
	
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, u := range users {
		if u.Name == currentUser {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}