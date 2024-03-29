package usecase_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xfiendx4life/media_tel_test/pkg/models"
	"github.com/xfiendx4life/media_tel_test/pkg/usecase"
)

func TestAdd(t *testing.T) {
	uc := usecase.New()
	testData := [][2]string{
		{"M", "K"},
		{"P", "L"},
		{"M", "P"},
		{"M", "K"},
	}
	uc.Add(testData)
	g := uc.GetGraph().Data
	assert.EqualValues(t, []models.Com{
		{Name: "K", Num: 2},
		{Name: "P", Num: 1},
	}, g["M"])
	assert.EqualValues(t, []models.Com{
		{Name: "M", Num: 2},
	},
		g["K"])
	assert.EqualValues(t, []models.Com{
		{Name: "L", Num: 1},
		{Name: "M", Num: 1},
	},
		g["P"])
	assert.EqualValues(t, []models.Com{
		{Name: "P", Num: 1},
	},
		g["L"])
}

// not cleant tests =(
func TestGetGraph(t *testing.T) {
	uc := usecase.New()
	testData := [][2]string{
		{"M", "K"},
		{"P", "L"},
		{"M", "P"},
		{"M", "K"},
	}
	uc.Add(testData)
	g := uc.GetGraph()
	assert.Less(t, math.Abs(g.Info.AverageCommunications-2.0), 0.001)
	assert.Equal(t, 1, g.Info.MinCommunications)
	assert.Equal(t, 3, g.Info.MaxCommunications)
	res, _ := json.Marshal(testData)
	fmt.Println(string(res))
}
