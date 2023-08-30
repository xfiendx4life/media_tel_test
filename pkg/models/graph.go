package models

type Info struct {
	// Min communication is a field to show minimal number
	// of communications between two persons
	MinCommunications int `json:"minComs"`
	// Max communications is a field to show maximum number
	// of communications between two persons
	MaxCommunications int `json:"maxComs"`
	// Total number of communications between people in graph
	AverageCommunications float64 `json:"average"`
}

type Com struct {
	// name of subject of communication
	Name string `json:"name"`
	// number of communication acts with the person
	Num int `json:"num"`
}

type Graph struct {
	// graph data
	Data map[string][]Com `json:"graph"`
	// common information
	Info Info `json:"info"`
}
