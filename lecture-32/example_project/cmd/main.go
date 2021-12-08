package main

import (
	"example_project/api/delivery"
	"example_project/api/delivery/middleware"
	"example_project/api/repository/postgres"
	"example_project/api/usecase"
	"log"
	"net/http"
)

func main() {
	mux := &http.ServeMux{}

	dbConn, err := postgres.NewPostgresDbInterface("localhost", 5433, "postgres")
	if err != nil {
		log.Fatalf("Db interface create error: %v", err)
	}

	processUserUsecase := usecase.NewProcessUserUsecase(dbConn)

	updateAgeUsecase := usecase.NewUpdateUserAgeUsecase(dbConn)

	delivery.NewGetInfoHandler(mux, processUserUsecase)

	delivery.NewUpdateAgeHandler(mux, updateAgeUsecase)

	handler := middleware.AuthMiddleware(mux)

	http.ListenAndServe(":8080", handler)
}
