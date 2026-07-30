// Package cli is a driving adapter: it turns terminal input into use-case calls
// and prints the results. It holds no business logic — swap it for httpserver
// later and the services stay untouched.
package cli

import (
	"bufio"
	"os"
	"fmt"

	"todoapp/param"
	"todoapp/services/category"
	"todoapp/services/task"
	"todoapp/services/user"
)

type CLI struct {
	userSvc user.Service
	taskSvc task.Service
	categorySvc category.Service

	scanner *bufio.Scanner
	authUser *param.UserInfo // nil until a user logs in
}

func New(userSvc user.Service, taskSvc task.Service, categorySvc category.Service) *CLI {
	return &CLI{
		userSvc: userSvc,
		taskSvc: taskSvc,
		categorySvc: categorySvc,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *CLI) Run() {
	fmt.Println("***** Welcome to TODO app *****")
	fmt.Println("commands: register | login | create-category | create-task | list-task | exit")

	for {
		command := c.prompt("\n Please enter a command: ")

		if command == "exit" {
			fmt.Println("bye!")
			return
		}

		// These commands need an authenticated user; log in on demand.
		if isProtected(command) && c.authUser == nil {
			fmt.Println("please login first.")
			c.login()
			if c.authUser == nil {
				continue
			}
		}

		switch command {
		case "register":
			c.register()
		case "login":
			c.login()
		case "create-category":
			c.createCategory()
		case "create-task":
			c.createTask()
		case "list-task":
			c.listTask()
		default:
			fmt.Println("command is not valid.")
		}
	}
}

func isProtected(command string) bool {
	switch command {
	case "create-category", "create-task", "list-task":
		return true
	default:
		return false
	}
}

// prompt prints a message and returns the next line of input.
func (c *CLI) prompt(msg string) string {
	fmt.Print(msg)
	c.scanner.Scan()

	return c.scanner.Text()
}
