package generator

type HandlerGen interface {
}

type HandlerGenerator struct {
}

func NewHandlerGen() HandlerGen {
	return &HandlerGenerator{}
}
