package generator

import (
	"errors"
	"github.com/rofi-go/cli/common"
	"github.com/rofi-go/cli/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

type StructureGen interface {
	CreateFolders(rofiConfig common.RofiConfig)
}

type StructureGenerator struct {
	log *log.Entry
}

func (s StructureGenerator) CreateFolders(rofiConfig common.RofiConfig) {

	projectPath := utils.GetProjectPath(rofiConfig.ProjectName)

	if _, err := os.Stat(projectPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(projectPath, 0755)
		if err != nil {
			s.log.Errorln(err)
		}
	}

	for _, microservice := range rofiConfig.MicroserviceComponents {
		err := os.MkdirAll(projectPath+"/build/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

		err = os.MkdirAll(projectPath+"/cmd/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

		err = os.MkdirAll(projectPath+"/deployment/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

		err = os.MkdirAll(projectPath+"/internal/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

		err = os.MkdirAll(projectPath+"/config/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

		err = os.MkdirAll(projectPath+"/docs/"+microservice.Name, 0755)
		if err != nil {
			s.log.Errorln(err)
		}

	}

}

func NewStructureGen(log *log.Entry) StructureGen {
	return &StructureGenerator{log: log}
}
