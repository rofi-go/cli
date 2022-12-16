package common

type RofiConfig struct {
	ProjectName            string                   `json:"projectName"`
	MicroserviceComponents []MicroserviceComponents `json:"microserviceComponents"`
}

type MicroserviceComponents struct {
	Name      string                 `json:"name"`
	Routes    map[string]interface{} `json:"routes"`
	Databases map[string]interface{} `json:"databases"`
}
