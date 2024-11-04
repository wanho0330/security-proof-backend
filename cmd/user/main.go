// Package main is the server for running User App.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/api/v1/apiv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"security-proof/internal/middleware"
	"security-proof/internal/user/controller"
	"security-proof/internal/user/repository"
	"security-proof/internal/user/service"
	"security-proof/pkg/auth"
	dbmanage "security-proof/pkg/manage/db"
)

func main() {
	tokenConfig := dbmanage.TokenConfig{}
	writeConfig := dbmanage.WriteConfig{}
	readConfig := dbmanage.ReadConfig{}
	baseAddr := "127.0.0.1:8080"

	tokenDB, err := dbmanage.NewRedis(tokenConfig.Dsn())
	if err != nil {
		log.Fatal(err)
		return
	}

	writeDB, err := dbmanage.NewDB(context.Background(), writeConfig.Dsn())
	if err != nil {
		log.Fatal(err)
		return
	}

	readDB, err := dbmanage.NewDB(context.Background(), readConfig.Dsn())
	if err != nil {
		log.Fatal(err)
		return
	}

	tokenRepo := auth.NewTokenRepo(tokenDB)
	commandRepo := repository.NewUserCommand(writeDB)
	queryRepo := repository.NewUserQuery(readDB)

	token := auth.NewToken(tokenRepo)
	commandService := service.NewUserCommand(token, commandRepo, queryRepo)
	queryService := service.NewUserQuery(token, queryRepo)

	userController := controller.NewUserController(commandService, queryService)

	mux := http.NewServeMux()
	path, handler := apiv1connect.NewUserServiceHandler(userController)

	mux.Handle(path, handler)
	server := &http.Server{
		Addr:              baseAddr,
		Handler:           h2c.NewHandler(middleware.WithCORS(mux), &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}
	_ = server.ListenAndServe()
}
