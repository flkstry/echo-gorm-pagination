package main

import (
	delivery "customer/app/customer/delivery/http"
	"customer/app/customer/repository"
	"customer/app/customer/usecase"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=wap12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	e := echo.New()
	custRepo := repository.NewCustomerRepository(db)
	timeoutContext := time.Duration(10 * time.Second)
	cUC := usecase.NewCustomerUsecase(timeoutContext, custRepo)
	delivery.NewCustomerHandler(e, cUC)

	log.Fatal(e.Start("127.0.0.1:3000"))
}
