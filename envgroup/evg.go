package envgroup

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/vmwarepivotallabs/cf-mgmt/config"
)

type defaultManager struct {
	cf  CFClient
	cfg config.Manager
}

func NewManager(cf CFClient, cfg config.Manager) Manager {
	return &defaultManager{cf, cfg}
}

func (d *defaultManager) GetEnvironmentVariableGroup(running bool) (cfclient.EnvironmentVariableGroup, error) {
	var (
		evg cfclient.EnvironmentVariableGroup
		err error
	)

	fn := d.cf.GetStagingEnvironmentVariableGroup
	if running {
		fn = d.cf.GetRunningEnvironmentVariableGroup
	}

	evg, err = fn()
	if err != nil {
		return evg, err
	}

	if evg == nil {
		evg = make(cfclient.EnvironmentVariableGroup) // have an empty map instead of nil so the keys are always in the config files
	}

	return evg, err
}

func (d *defaultManager) SetEnvironmentVariableGroup(running bool) error {
	gCfg, err := d.cfg.GetGlobalConfig()
	if err != nil {
		return err
	}

	evg := gCfg.StagingEnvironmentVariableGroup
	fn := d.cf.SetStagingEnvironmentVariableGroup
	if running {
		evg = gCfg.RunningEnvironmentVariableGroup
		fn = d.cf.SetRunningEnvironmentVariableGroup
	}

	if err = fn(evg); err != nil {
		return err
	}

	return nil
}
