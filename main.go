package main	

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"rest_api_golang/app"
	"rest_api_golang/controller"
	"rest_api_golang/helper"
	"rest_api_golang/repository"
	"rest_api_golang/service"
	"rest_api_golang/middleware"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository :=repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.ErrorConditionCheck(err)
}