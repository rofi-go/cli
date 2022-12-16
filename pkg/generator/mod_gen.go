package generator

import (
	"github.com/rofi-go/cli/common"
	"github.com/rofi-go/cli/utils"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type ModGen interface {
	InitMod(rofiConfig common.RofiConfig)
	TidyMod(rofiConfig common.RofiConfig)
}

type ModGenerator struct {
	log *log.Entry
}

func (m ModGenerator) InitMod(rofiConfig common.RofiConfig) {
	projectPath := utils.GetProjectPath(rofiConfig.ProjectName)

	app := "go"
	arg0 := "mod"
	arg1 := "init"

	cmd := exec.Command(app, arg0, arg1)
	cmd.Dir = projectPath

	_, err := cmd.Output()

	if err != nil {
		m.log.Errorln(err.Error())
	}
}

func (m ModGenerator) TidyMod(rofiConfig common.RofiConfig) {
	projectPath := utils.GetProjectPath(rofiConfig.ProjectName)

	app := "go"
	arg0 := "mod"
	arg1 := "tidy"

	cmd := exec.Command(app, arg0, arg1)
	cmd.Dir = projectPath

	_, err := cmd.Output()

	if err != nil {
		m.log.Errorln(err.Error())
	}
}

func NewModGen(log *log.Entry) ModGen {
	return &ModGenerator{log: log}
}
