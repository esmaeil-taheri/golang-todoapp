package cli

import (
	"fmt"

	"todoapp/param"
)

func (c *CLI) createCategory() {
	req := param.CreateCategoryRequest{
		Title:  c.prompt("Category title: "),
		Color:  c.prompt("Category color: "),
		UserID: c.authUser.ID,
	}

	resp, err := c.categorySvc.Create(req)
	if err != nil {
		fmt.Println("create category failed:", err)
		return
	}

	fmt.Printf("category created: #%d %s (%s)\n", resp.Category.ID, resp.Category.Title, resp.Category.Color)
}
