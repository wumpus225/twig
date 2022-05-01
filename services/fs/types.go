package fs

type PlantDataFile struct {
	OwnedPlants []Plant `json:"owned_plants"`
}

type Plant struct {
	Name string `json:"name"`
}