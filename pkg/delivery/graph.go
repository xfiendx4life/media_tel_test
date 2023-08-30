package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xfiendx4life/media_tel_test/pkg/usecase"
)

type Deliver struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) *Deliver {
	return &Deliver{
		uc: uc,
	}
}

func (d *Deliver) Graph(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusOK, d.uc.GetGraph()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Empty Graph")
	}
	return nil
}

func (d *Deliver) Add(ctx echo.Context) error {
	data := ctx.Request().Body
	a := make([][2]string, 0)
	err := json.NewDecoder(data).Decode(&a)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("can't decode your data %s", err))
	}
	d.uc.Add(a)
	return ctx.NoContent(http.StatusAccepted)
}
