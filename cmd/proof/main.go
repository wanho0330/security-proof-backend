// Package main is the server for running Proof App.
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
	"security-proof/internal/proof/controller"
	"security-proof/internal/proof/repository"
	"security-proof/internal/proof/service"
	"security-proof/pkg/auth"
	chainmanage "security-proof/pkg/manage/chain"
	dbmanage "security-proof/pkg/manage/db"
	usermanage "security-proof/pkg/manage/user"
)

func main() {
	tokenConfig := dbmanage.TokenConfig{}
	writeConfig := dbmanage.WriteConfig{}
	readConfig := dbmanage.ReadConfig{}
	chainConfig := chainmanage.Config{}
	userConfig := usermanage.Config{}
	baseAddr := "127.0.0.2:8081"

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
	commandRepo := repository.NewProofCommand(writeDB)
	queryRepo := repository.NewProofQuery(readDB)

	token := auth.NewToken(tokenRepo)
	chain := chainmanage.NewChain(chainConfig.FromEnv())
	user := usermanage.NewUser(userConfig.FromEnv())

	commandService := service.NewProofCommand(token, commandRepo, queryRepo, chain)
	queryService := service.NewProofQuery(token, queryRepo, user)

	proofController := controller.NewProofController(commandService, queryService)

	mux := http.NewServeMux()
	path, handler := apiv1connect.NewProofServiceHandler(proofController)

	mux.Handle(path, handler)
	// 이미지 부분은 grpc를 사용하지 않고 이미지를 전달합니다.
	mux.HandleFunc("/apiv1/readFirstImage/", proofController.ReadFirstImage)
	mux.HandleFunc("/apiv1/readSecondImage/", proofController.ReadSecondImage)

	server := &http.Server{
		Addr:              baseAddr,
		Handler:           h2c.NewHandler(middleware.WithCORS(mux), &http2.Server{}),
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       90 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
