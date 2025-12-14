package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/item"
	"ecommerce/repo"
	"ecommerce/rest"
	itemHandler "ecommerce/rest/handlers/item"
	usrHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repo
	itemRepo := repo.NewItemRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// services
	usrSvc := user.NewService(userRepo)
	itemSvc := item.NewService(itemRepo)

	middlewares := middleware.NewMiddlewares(cnf)
	itemHandler := itemHandler.NewHandler(middlewares, itemSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(cnf, itemHandler, userHandler)
	server.Start()
}
