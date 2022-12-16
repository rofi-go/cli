package generator

import "github.com/rofi-go/cli/common"

type ModelGen interface {
	CreateModel(rofiConfig common.RofiConfig)
}

type ModelGenerator struct {
}

func (m ModelGenerator) CreateModel(rofiConfig common.RofiConfig) {

}

func NewModelGen() ModelGen {
	return &ModelGenerator{}
}
