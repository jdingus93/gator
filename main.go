package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jdingus93/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err !=  nil {
		log.Fatal(err)
	}
	
	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	s := &state{config: cfg}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments provided")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

type state struct {
	config config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required")
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User set to: %s\n", cmd.args[0])
	return nil
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c* commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("unkown command: %s", cmd.name)
	}
	return handler(s, cmd)
}


func (c* commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}