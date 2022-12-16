package generator

type UsecaseGen interface {
}

type UsecaseGenerator struct {
}

func NewUsecaseGen() UsecaseGen {
	return &UsecaseGenerator{}
}
