package usecase

import "github.com/xfiendx4life/media_tel_test/pkg/models"

type Usecase struct {
	Graph models.Graph
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

func addCom(graph map[string][]models.Com, first, second string) (maxComs int) {
	maxComs = 1
	if _, ok := graph[first]; !ok {
		// graph[first] = make([]models.Com, 1)
		graph[first] = []models.Com{
			{
				Name: second,
				Num:  1},
		}
	} else {
		ind := getIndex(graph[first], second)
		if ind == -1 {
			graph[first] = append(graph[first], models.Com{
				Name: second,
				Num:  1,
			})
		} else {
			graph[first][ind].Num++
			maxComs = graph[first][ind].Num
		}
	}
	return
}

func (uc *Usecase) Add(list [][2]string) {
	for _, row := range list {
		fst, scnd := row[0], row[1]

		if max := addCom(uc.Graph.Data, fst, scnd); max > uc.Graph.Info.MaxCommunications {
			uc.Graph.Info.MaxCommunications = max
		}
		if max := addCom(uc.Graph.Data, scnd, fst); max > uc.Graph.Info.MaxCommunications {
			uc.Graph.Info.MaxCommunications = max
		}
	}
}
