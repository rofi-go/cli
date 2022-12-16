package service

import (
	"github.com/rofi-go/cli/common"
	"github.com/rofi-go/cli/config/cli"
	"github.com/rofi-go/cli/internal/cli/domain/ports/input"
	"github.com/rofi-go/cli/pkg/generator"
	log "github.com/sirupsen/logrus"
)

type AppSvc struct {
	log     *log.Entry
	configs *cli.AppConfig
}

func (s AppSvc) StartApp() {
	var folderGen generator.StructureGen
	folderGen = generator.NewStructureGen(s.log)

	var microservices []common.MicroserviceComponents

	api := common.MicroserviceComponents{
		Name:      "api",
		Routes:    nil,
		Databases: nil,
	}

	auth := common.MicroserviceComponents{
		Name:      "auth",
		Routes:    nil,
		Databases: nil,
	}

	microservices = append(microservices, api, auth)

	folderGen.CreateFolders(common.RofiConfig{
		ProjectName:            "test",
		MicroserviceComponents: microservices,
	})

	var modGen generator.ModGen
	modGen = generator.NewModGen(s.log)

	modGen.InitMod(common.RofiConfig{
		ProjectName:            "test",
		MicroserviceComponents: microservices,
	})

}

func NewAppService(log *log.Entry, configs *cli.AppConfig) input.AppUseCase {
	return &AppSvc{log: log, configs: configs}
}
