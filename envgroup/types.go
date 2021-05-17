package envgroup

import "github.com/cloudfoundry-community/go-cfclient"

//go:generate counterfeiter -o fakes/fake_cfclient.go . CFClient
// CFClient -
type CFClient interface {
	GetRunningEnvironmentVariableGroup() (cfclient.EnvironmentVariableGroup, error)
	SetRunningEnvironmentVariableGroup(evg cfclient.EnvironmentVariableGroup) error
	GetStagingEnvironmentVariableGroup() (cfclient.EnvironmentVariableGroup, error)
	SetStagingEnvironmentVariableGroup(evg cfclient.EnvironmentVariableGroup) error
}

type Manager interface {
	GetEnvironmentVariableGroup(running bool) (cfclient.EnvironmentVariableGroup, error)
	SetEnvironmentVariableGroup(running bool) error
}
