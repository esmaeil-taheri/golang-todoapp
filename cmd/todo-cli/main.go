package main

import (
	"log"

	"todoapp/config"
	"todoapp/delivery/cli"
	"todoapp/repository/filestore"
	"todoapp/repository/memory"
	"todoapp/services/category"
	"todoapp/services/task"
	"todoapp/services/user"
)

func main() {
	cfg := config.Load()

	// Repositories (driven adapters).
	userStore, err := filestore.NewUserStore(cfg.UserFilePath)
	if err != nil {
		log.Fatalln("can't initialize user store:", err)
	}

	taskStore := memory.NewTaskStore()
	categoryStore := memory.NewCategoryStore()

	// Services (application / business logic). categoryStore is injected twice:
	// as the category repository and as the task service's CategoryValidator.
	userSvc := user.New(userStore, cfg.BcryptCost)
	categorySvc := category.New(categoryStore)
	taskSvc := task.New(taskStore, categoryStore)

	app := cli.New(userSvc, taskSvc, categorySvc)
	app.Run()
}