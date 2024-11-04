// Package main is the server for running Dashboard App.
package main

import (
	"context"
	"log"
	"net/http"
	usermanage "security-proof/pkg/manage/user"
	"time"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/api/v1/apiv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"security-proof/internal/dashboard/controller"
	"security-proof/internal/dashboard/repository"
	"security-proof/internal/dashboard/service"
	"security-proof/internal/middleware"
	"security-proof/pkg/auth"
	dbmanage "security-proof/pkg/manage/db"
	elasticmanage "security-proof/pkg/manage/elastic"
)

func main() {
	tokenConfig := dbmanage.TokenConfig{}
	elaConfig := elasticmanage.Config{}
	readConfig := dbmanage.ReadConfig{}
	userConfig := usermanage.Config{}
	baseAddr := "127.0.0.3:8082"

	tokenDB, err := dbmanage.NewRedis(tokenConfig.Dsn())
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
	queryRepo := repository.NewDashboardQuery(readDB)

	token := auth.NewToken(tokenRepo)
	elastic := elasticmanage.NewElastic(elaConfig.FromEnv())
	user := usermanage.NewUser(userConfig.FromEnv())

	queryService := service.NewDashboardService(token, queryRepo, elastic, user)

	dashboardController := controller.NewDashboardController(queryService)

	mux := http.NewServeMux()
	path, handler := apiv1connect.NewDashboardServiceHandler(dashboardController)

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
