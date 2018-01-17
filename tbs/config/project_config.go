package config

import (
	"fmt"
	"template-builder/tbs/config/business"
	"template-builder/tbs/config/database"
	"template-builder/tbs/config/shared"
	"template-builder/tbs/config/station"
)

var logger = shared.ConfigLog

type ProjectConfig struct {
	Business business.BusinessConfig `json:"business"`
	Database database.DatabaseConfig `json:"database"`
	Station  station.StationConfig   `json:"station"`
}

func (p ProjectConfig) String() string {
	return fmt.Sprintf("config:[%s][%s][%s]", p.Database, p.Station, p.Business)
}

func (p *ProjectConfig) Check() {
	p.Station.Check()
	p.Database.Check()
	p.Business.Check()
}
