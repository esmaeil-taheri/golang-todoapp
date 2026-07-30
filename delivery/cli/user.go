package cli

import (
	"fmt"

	"todoapp/param"
)

func (c *CLI) register() {
	req := param.RegisterRequest{
		Name:     c.prompt("Name: "),
		Email:    c.prompt("Email: "),
		Password: c.prompt("Password: "),
	}

	resp, err := c.userSvc.Register(req)
	if err != nil {
		fmt.Println("register failed:", err)
		return
	}

	fmt.Printf("user registered successfully: #%d %s\n", resp.User.ID, resp.User.Email)
}

func (c *CLI) login() {
	req := param.LoginRequest{
		Email:    c.prompt("Email: "),
		Password: c.prompt("Password: "),
	}

	resp, err := c.userSvc.Login(req)
	if err != nil {
		fmt.Println("login failed:", err)
		return
	}

	user := resp.User
	c.authUser = &user
	fmt.Printf("logged in as %s\n", resp.User.Email)
}
