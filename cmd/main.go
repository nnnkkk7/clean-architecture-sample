package main

import (
	"log"
	"net/http"

	"database/sql"

	"github/nnnkkk7/clean-architecture-sample/cmd/handler"
	"github/nnnkkk7/clean-architecture-sample/internal/domain/service"
	"github/nnnkkk7/clean-architecture-sample/internal/infra/mysql"
	"github/nnnkkk7/clean-architecture-sample/internal/usecase"

	"github.com/gorilla/mux"
)

const driverName = "mysql"
const dsn = ""

func main() {
	r := mux.NewRouter()
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	repo := mysql.NewUserRepository(conn)
	service := service.NewService(repo)
	us := usecase.NewUserCase(service)
	handler := handler.NewUserHandler(us)

	r.HandleFunc("/user", handler.UserHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
