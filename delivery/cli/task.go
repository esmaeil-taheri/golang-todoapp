package cli

import (
	"fmt"
	"strconv"

	"todoapp/param"
)

func (c *CLI) createTask() {
	title := c.prompt("Task title: ")

	categoryRaw := c.prompt("Category ID: ")
	categoryID, err := strconv.ParseUint(categoryRaw, 10, 64)
	if err != nil {
		fmt.Println("category id is not valid:", err)
		return
	}

	dueDate := c.prompt("Due date: ")

	resp, err := c.taskSvc.Create(param.CreateTaskRequest{
		Title:      title,
		DueDate:    dueDate,
		CategoryID: uint(categoryID),
		UserID:     c.authUser.ID,
	})
	if err != nil {
		fmt.Println("create task failed:", err)
		return
	}

	fmt.Printf("task created: #%d %s\n", resp.Task.ID, resp.Task.Title)
}

func (c *CLI) listTask() {
	resp, err := c.taskSvc.List(param.ListTaskRequest{UserID: c.authUser.ID})
	if err != nil {
		fmt.Println("list task failed:", err)
		return
	}

	if len(resp.Tasks) == 0 {
		fmt.Println("no tasks yet.")
		return
	}

	fmt.Println("your tasks:")
	for _, t := range resp.Tasks {
		status := "todo"
		if t.IsDone {
			status = "done"
		}
		fmt.Printf("  #%d [%s] %s (due: %s)\n", t.ID, status, t.Title, t.DueDate)
	}
}
