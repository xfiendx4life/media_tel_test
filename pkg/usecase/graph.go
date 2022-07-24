package usecase

import (
	"math"

	"github.com/xfiendx4life/media_tel_test/pkg/models"
)

type Usecase struct {
	Graph models.Graph
	// to count every new person
	numOfPeople int
}

func New() *Usecase {
	return &Usecase{
		Graph: models.Graph{
			Data: make(map[string][]models.Com),
		},
	}
}

func getIndex(a []models.Com, data string) int {
	for ind, item := range a {
		if item.Name == data {
			return ind
		}
	}
	return -1
}

func (uc *Usecase) addCom(first, second string) {
	graph := uc.Graph.Data
	if _, ok := graph[first]; !ok {
		graph[first] = []models.Com{
			{
				Name: second,
				Num:  1},
		}
		// the only place where person adds
		uc.numOfPeople++
	} else {
		ind := getIndex(graph[first], second)
		if ind == -1 {
			graph[first] = append(graph[first], models.Com{
				Name: second,
				Num:  1,
			})
		} else {
			graph[first][ind].Num++
		}
	}
}

func (uc *Usecase) Add(list [][2]string) {
	for _, row := range list {
		fst, scnd := row[0], row[1]

		uc.addCom(fst, scnd)
		uc.addCom(scnd, fst)
	}
}

func (uc *Usecase) countInfo() {
	min := math.MaxInt
	fullSum := 0
	max := 0
	for _, node := range uc.Graph.Data {
		sum := 0
		for _, itm := range node {
			sum += itm.Num
		}
		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
		fullSum += sum
	}
	uc.Graph.Info.MinCommunications = min
	uc.Graph.Info.AverageCommunications = float64(fullSum) / float64(uc.numOfPeople)
	uc.Graph.Info.MaxCommunications = max
}

func (uc *Usecase) GetGraph() *models.Graph {
	// count min and average
	uc.countInfo()
	return &uc.Graph
}
