package delivery

import (
	"context"
	"customer/domain"
	"customer/pkg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	CustUCase domain.CustomerUsecase
}

type ReponseFetchCustomer struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    []domain.Customer `json:"data"`
	Meta    pkg.Pagination    `json:"meta"`
}

type ReponseGetCustomer struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    domain.Customer `json:"data"`
}

func NewCustomerHandler(e *echo.Echo, custUC domain.CustomerUsecase) {
	handler := &CustomerHandler{
		CustUCase: custUC,
	}

	e.GET("/cust/:id", handler.GetCustomerById)
	e.GET("/cust", handler.Fetch)
}

func (ch *CustomerHandler) Fetch(c echo.Context) error {
	// get pagination from params
	pg := pkg.Pagination{}
	cursor := c.QueryParam("cursor")
	order := c.QueryParam("order")
	perPage, _ := strconv.Atoi((c.QueryParam("per_page")))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	if perPage == 0 {
		perPage = 10
	}

	if cursor == "" {
		cursor = "customerid"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	pg.Limit = perPage
	pg.Page = page
	pg.Sort = fmt.Sprintf("%s %s", cursor, order)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	cust, nextCursor, err := ch.CustUCase.Fetch(ctx, pg)
	if err != nil {
		return c.NoContent(getStatusCode(err))
	}

	return c.JSON(http.StatusOK, &ReponseFetchCustomer{
		Status:  true,
		Message: "succes get data",
		Data:    cust,
		Meta:    nextCursor,
	})
}

func (ch *CustomerHandler) GetCustomerById(c echo.Context) error {
	idFromParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(getStatusCode(err))
	}

	id := int64(idFromParam)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	cust, err := ch.CustUCase.GetById(ctx, id)
	if err != nil {
		return c.NoContent(getStatusCode(err))
	}

	return c.JSON(http.StatusOK, &ReponseGetCustomer{
		Status:  true,
		Message: "success get customer data",
		Data:    cust,
	})
}

func getStatusCode(err error) int {
	logrus.Error(err)
	switch err {
	case domain.ErrCustomerNoFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
