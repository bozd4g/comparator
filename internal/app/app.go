package app

import (
	"fmt"
	"os"
)

func New() Application {
	return &application{}
}

func (a *application) Build() Application {
	a.AddRouter()
	a.AddControllers().InitMiddlewares().AddSwagger()
	return a
}

func (a *application) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Println("Application starting..")
	err := a.engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	return err
}
