package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xfiendx4life/media_tel_test/pkg/delivery"
	"github.com/xfiendx4life/media_tel_test/pkg/models"
)

type mockUsecase struct {
	d models.Graph
}

func new() *mockUsecase {
	return &mockUsecase{
		d: models.Graph{
			Data: map[string][]models.Com{
				"K": {{Name: "M", Num: 2}},
				"L": {{Name: "P", Num: 1}},
				"M": {{Name: "K", Num: 2}, models.Com{Name: "P", Num: 1}},
				"P": {{Name: "L", Num: 1}, models.Com{Name: "M", Num: 1}},
			},
			Info: models.Info{
				MinCommunications:     1,
				MaxCommunications:     3,
				AverageCommunications: 2},
		},
	}
}

func (mu *mockUsecase) Add(list [][2]string) {

}
func (mu *mockUsecase) GetGraph() *models.Graph {
	return &mu.d
}

var (
	jsn       = `{"graph":{"K":[{"name":"M","num":2}],"L":[{"name":"P","num":1}],"M":[{"name":"K","num":2},{"name":"P","num":1}],"P":[{"name":"L","num":1},{"name":"M","num":1}]},"info":{"minComs":1,"maxComs":3,"average":2}}`
	jsn_array = `[["M","K"],["P","L"],["M","P"],["M","K"]]`
)

func TestAdd(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsn_array))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	d := delivery.New(new())
	assert.NoError(t, d.Add(c))
}

func TestAddErr(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsn))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	d := delivery.New(new())
	assert.Error(t, d.Add(c))
}

func TestGetGraph(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	d := delivery.New(new())
	assert.NoError(t, d.Graph(c))
	assert.Equal(t, jsn+"\n", rec.Body.String())
}
