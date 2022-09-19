package main

import (
	login_hdl "gokomodo_test/internal/app/login/handler"
	login_repo "gokomodo_test/internal/app/login/repository"
	login_uc "gokomodo_test/internal/app/login/usecase"
	"log"

	product_hdl "gokomodo_test/internal/app/product/handler"
	product_repo "gokomodo_test/internal/app/product/repository"
	product_uc "gokomodo_test/internal/app/product/usecase"

	order_hdl "gokomodo_test/internal/app/order/handler"
	order_repo "gokomodo_test/internal/app/order/repository"
	order_uc "gokomodo_test/internal/app/order/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	dbUrl := "root@tcp(localhost:3306)/gokomodo_test?parseTime=true"
	dbConn, err := sqlx.Connect("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	// login
	loginRepo := login_repo.NewDb(dbConn)
	loginUsecase := login_uc.NewUsecase(loginRepo)
	login_hdl.NewHandler(e, loginUsecase)

	// Product
	productRepo := product_repo.NewDb(dbConn)
	productUsecase := product_uc.NewUsecase(productRepo)
	product_hdl.NewHandler(e, productUsecase)

	// Order
	orderRepo := order_repo.NewDb(dbConn)
	orderUsecase := order_uc.NewUsecase(orderRepo, productRepo, loginRepo)
	order_hdl.NewHandler(e, orderUsecase)

	e.Logger.Fatal(e.Start(":8000"))
}
