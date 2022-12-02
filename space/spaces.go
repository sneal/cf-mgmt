package space

import (
	"context"
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/vmwarepivotallabs/cf-mgmt/config"
	"github.com/vmwarepivotallabs/cf-mgmt/organizationreader"
	"github.com/vmwarepivotallabs/cf-mgmt/uaa"
	"github.com/xchapter7x/lo"
)

// NewManager -
func NewManager(spaceClient CFSpaceClient, spaceFeatureClient CFSpaceFeatureClient, orgClient CFOrganizationClient, jobClient CFJobClient,
	uaaMgr uaa.Manager,
	orgReader organizationreader.Reader,
	cfg config.Reader, peek bool) Manager {
	return &DefaultManager{
		Cfg:                cfg,
		UAAMgr:             uaaMgr,
		SpaceClient:        spaceClient,
		OrgClient:          orgClient,
		SpaceFeatureClient: spaceFeatureClient,
		JobClient:          jobClient,
		OrgReader:          orgReader,
		Peek:               peek,
	}
}

// DefaultManager -
type DefaultManager struct {
	Cfg                config.Reader
	SpaceClient        CFSpaceClient
	SpaceFeatureClient CFSpaceFeatureClient
	OrgClient          CFOrganizationClient
	JobClient          CFJobClient
	UAAMgr             uaa.Manager
	OrgReader          organizationreader.Reader
	Peek               bool
	spaces             []*resource.Space
}

func (m *DefaultManager) IsSSHEnabled(spaceGUID string) (bool, error) {
	return m.SpaceFeatureClient.IsSSHEnabled(context.Background(), spaceGUID)
}

func (m *DefaultManager) UpdateSpaceSSH(sshAllowed bool, space *resource.Space) error {
	return m.SpaceFeatureClient.EnableSSH(context.Background(), space.GUID, sshAllowed)
}

func (m *DefaultManager) init() error {
	if m.spaces == nil {
		spaces, err := m.SpaceClient.ListAll(context.Background(), nil)
		if err != nil {
			return err
		}
		m.spaces = spaces
	}
	return nil
}

// UpdateSpaces -
func (m *DefaultManager) UpdateSpaces() error {
	m.spaces = nil
	spaceConfigs, err := m.Cfg.GetSpaceConfigs()
	if err != nil {
		return err
	}
	err = m.init()
	if err != nil {
		return err
	}
	for _, input := range spaceConfigs {
		space, err := m.FindSpace(input.Org, input.Space)
		if err != nil {
			continue // TODO silent failure
		}
		spaceAllowsSSH, err := m.SpaceFeatureClient.IsSSHEnabled(context.Background(), space.GUID)
		if err != nil {
			continue // TODO silent failure
		}

		lo.G.Debug("Processing space", space.Name)
		if input.AllowSSHUntil != "" {
			allowUntil, err := time.Parse(time.RFC3339, input.AllowSSHUntil)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("Unable to parse %s with format %s", input.AllowSSHUntil, time.RFC3339))
			}
			if allowUntil.After(time.Now()) && !spaceAllowsSSH {
				if m.Peek {
					lo.G.Infof("[dry-run]: temporarily enabling sshAllowed for org/space %s/%s until %s", input.Org, space.Name, input.AllowSSHUntil)
					continue
				}
				lo.G.Infof("temporarily enabling sshAllowed for org/space %s/%s until %s", input.Org, space.Name, input.AllowSSHUntil)
				if err := m.UpdateSpaceSSH(true, space); err != nil {
					return err
				}
			}
			if allowUntil.Before(time.Now()) && spaceAllowsSSH {
				if m.Peek {
					lo.G.Infof("[dry-run]: removing temporarily enabling sshAllowed for org/space %s/%s as past %s", input.Org, space.Name, input.AllowSSHUntil)
					continue
				}
				lo.G.Infof("removing temporarily enabling sshAllowed for org/space %s/%s as past %s", input.Org, space.Name, input.AllowSSHUntil)
				if err := m.UpdateSpaceSSH(false, space); err != nil {
					return err
				}
			}
		} else {
			if input.AllowSSH != spaceAllowsSSH {
				if m.Peek {
					lo.G.Infof("[dry-run]: setting sshAllowed to %v for org/space %s/%s", input.AllowSSH, input.Org, space.Name)
					continue
				}
				lo.G.Infof("setting sshAllowed to %v for org/space %s/%s", input.AllowSSH, input.Org, space.Name)
				if err := m.UpdateSpaceSSH(input.AllowSSH, space); err != nil {
					return err
				}
			}
		}
	}
	m.spaces = nil
	return nil
}

func (m *DefaultManager) ListSpaces(orgGUID string) ([]*resource.Space, error) {
	if m.spaces == nil {
		err := m.init()
		if err != nil {
			return nil, err
		}
	}
	spaces := []*resource.Space{}
	for _, space := range m.spaces {
		if strings.EqualFold(space.Relationships.Organization.Data.GUID, orgGUID) {
			spaces = append(spaces, space)
		}
	}
	return spaces, nil

}

// FindSpace -
func (m *DefaultManager) FindSpace(orgName, spaceName string) (*resource.Space, error) {
	orgGUID, err := m.OrgReader.GetOrgGUID(orgName)
	if err != nil {
		return nil, err
	}
	spaces, err := m.ListSpaces(orgGUID)
	if err != nil {
		return nil, err
	}
	for _, theSpace := range spaces {
		if theSpace.Name == spaceName {
			return theSpace, nil
		}
	}
	if m.Peek {
		return &resource.Space{
			Name: spaceName,
			GUID: fmt.Sprintf("%s-dry-run-space-guid", spaceName),
			Relationships: &resource.SpaceRelationships{
				Organization: &resource.ToOneRelationship{
					Data: &resource.Relationship{
						GUID: fmt.Sprintf("%s-dry-run-org-guid", orgName),
					},
				},
			},
		}, nil
	}
	return nil, fmt.Errorf("space [%s] not found in org [%s]", spaceName, orgName)
}

func (m *DefaultManager) CreateSpace(spaceName, orgName, orgGUID string) error {
	if m.Peek {
		lo.G.Infof("[dry-run]: create space %s for org %s", spaceName, orgName)
		return nil
	}
	lo.G.Infof("create space %s for org %s", spaceName, orgName)
	space, err := m.SpaceClient.Create(context.Background(), resource.NewSpaceCreate(spaceName, orgGUID))
	m.spaces = append(m.spaces, space)
	return err
}

func (m *DefaultManager) RenameSpace(originalSpaceName, spaceName, orgName string) error {
	if m.Peek {
		lo.G.Infof("[dry-run]: rename space %s for org %s to %s", originalSpaceName, orgName, spaceName)
		return nil
	}
	lo.G.Infof("rename space %s for org %s to %s", originalSpaceName, orgName, spaceName)

	space, err := m.FindSpace(orgName, originalSpaceName)
	if err != nil {
		return err
	}
	_, err = m.SpaceClient.Update(context.Background(), space.GUID, &resource.SpaceUpdate{
		Name: spaceName,
	})
	space.Name = spaceName
	return err
}

// CreateSpaces -
func (m *DefaultManager) CreateSpaces() error {
	m.spaces = nil
	configSpaceList, err := m.Cfg.GetSpaceConfigs()
	if err != nil {
		return err
	}
	err = m.init()
	if err != nil {
		return err
	}
	for _, space := range configSpaceList {
		orgGUID, err := m.OrgReader.GetOrgGUID(space.Org)
		if err != nil {
			return err
		}
		spaces, err := m.ListSpaces(orgGUID)
		if err != nil {
			continue
		}

		if m.doesSpaceExist(spaces, space.Space) {
			lo.G.Debugf("[%s] space already exists in org [%s]", space.Space, space.Org)
			continue
		} else if doesSpaceExistFromRename(space.OriginalSpace, spaces) {
			lo.G.Debugf("renamed space [%s] already exists as [%s]", space.Space, space.OriginalSpace)
			if err = m.RenameSpace(space.OriginalSpace, space.Space, space.Org); err != nil {
				return err
			}

			continue
		} else {
			lo.G.Debugf("[%s] space doesn't exist in [%v]", space.Space, spaces)
		}
		if err = m.CreateSpace(space.Space, space.Org, orgGUID); err != nil {
			lo.G.Error(err)
			return err
		}
	}
	m.spaces = nil
	return nil
}

func (m *DefaultManager) doesSpaceExist(spaces []*resource.Space, spaceName string) bool {
	for _, space := range spaces {
		if strings.EqualFold(space.Name, spaceName) {
			return true
		}
	}
	return false
}

func doesSpaceExistFromRename(spaceName string, spaces []*resource.Space) bool {
	for _, space := range spaces {
		if strings.EqualFold(space.Name, spaceName) {
			return true
		}
	}
	return false
}

func (m *DefaultManager) DeleteSpaces() error {
	m.spaces = nil
	configSpaceList, err := m.Cfg.Spaces()
	if err != nil {
		return err
	}

	err = m.init()
	if err != nil {
		return err
	}
	for _, input := range configSpaceList {
		if !input.EnableDeleteSpaces {
			lo.G.Debugf("Space deletion is not enabled for %s.  Set enable-delete-spaces: true in spaces.yml", input.Org)
			continue //Skip all orgs that have not opted-in
		}
		renamedSpaces := make(map[string]string)
		configuredSpaces := make(map[string]bool)
		for _, spaceName := range input.Spaces {
			spaceCfg, err := m.Cfg.GetSpaceConfig(input.Org, spaceName)
			if err != nil {
				return err
			}
			if spaceCfg.OriginalSpace != "" {
				renamedSpaces[spaceCfg.OriginalSpace] = spaceName
			}
			configuredSpaces[spaceName] = true
		}

		org, err := m.OrgReader.FindOrg(input.Org)
		if err != nil {
			return err
		}
		spaces, err := m.ListSpaces(org.GUID)
		if err != nil {
			return err
		}

		spacesToDelete := make([]*resource.Space, 0)
		for _, space := range spaces {
			if _, exists := configuredSpaces[space.Name]; !exists {
				if _, renamed := renamedSpaces[space.Name]; !renamed {
					spacesToDelete = append(spacesToDelete, space)
				}
			}
		}

		for _, space := range spacesToDelete {
			if err := m.DeleteSpace(space, input.Org); err != nil {
				return err
			}
		}
	}
	m.spaces = nil
	return nil
}

func (m *DefaultManager) ClearMetadata(space *resource.Space, orgName string) error {
	if m.Peek {
		lo.G.Infof("[dry-run]: removing space metadata from space %s in org %s", space.Name, orgName)
		return nil
	}
	lo.G.Infof("removing space metadata from space %s in org %s", space.Name, orgName)
	r := &resource.SpaceUpdate{
		Metadata: space.Metadata,
	}
	r.Metadata.Clear()
	_, err := m.SpaceClient.Update(context.Background(), space.GUID, r)
	return err
}

// DeleteSpace - deletes a space based on GUID
func (m *DefaultManager) DeleteSpace(space *resource.Space, orgName string) error {
	if m.Peek {
		lo.G.Infof("[dry-run]: delete space with %s from org %s", space.Name, orgName)
		return nil
	}
	lo.G.Infof("delete space with %s from org %s", space.Name, orgName)
	jobGUID, err := m.SpaceClient.Delete(context.Background(), space.GUID)
	if err != nil {
		return err
	}
	return m.JobClient.PollComplete(context.Background(), jobGUID, nil)
}

func (m *DefaultManager) UpdateSpacesMetadata() error {
	spaceConfigs, err := m.Cfg.GetSpaceConfigs()
	if err != nil {
		return err
	}

	globalCfg, err := m.Cfg.GetGlobalConfig()
	if err != nil {
		return err
	}

	for _, spaceConfig := range spaceConfigs {
		if spaceConfig.Metadata != nil {
			space, err := m.FindSpace(spaceConfig.Org, spaceConfig.Space)
			if err != nil {
				continue
			}
			metadata := &resource.Metadata{}
			if spaceConfig.Metadata.Labels != nil {
				for key, value := range spaceConfig.Metadata.Labels {
					if len(value) > 0 {
						metadata.SetLabel(globalCfg.MetadataPrefix, key, value)
					} else {
						metadata.RemoveLabel(globalCfg.MetadataPrefix, key)
					}
				}
			}
			if spaceConfig.Metadata.Annotations != nil {
				for key, value := range spaceConfig.Metadata.Annotations {
					if len(value) > 0 {
						metadata.SetAnnotation(globalCfg.MetadataPrefix, key, value)
					} else {
						metadata.RemoveAnnotation(globalCfg.MetadataPrefix, key)
						// For bug in CC API that removal doesn't include prefix
						metadata.RemoveAnnotation("", key)
					}
				}
			}
			err = m.UpdateSpaceMetadata(spaceConfig.Org, space, metadata)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *DefaultManager) UpdateSpaceMetadata(org string, space *resource.Space, metadata *resource.Metadata) error {
	if m.Peek {
		lo.G.Infof("[dry-run]: update org/space %s/%s metadata", org, space.Name)
		return nil
	}
	lo.G.Infof("update org/space %s/%s metadata", org, space.Name)
	r := &resource.SpaceUpdate{
		Metadata: metadata,
	}
	_, err := m.SpaceClient.Update(context.Background(), space.GUID, r)
	return err
}

func (m *DefaultManager) DeleteSpacesForOrg(orgGUID, orgName string) (err error) {
	spaces, err := m.ListSpaces(orgGUID)
	if err != nil {
		return err
	}
	for _, space := range spaces {
		err := m.DeleteSpace(space, orgName)
		if err != nil {
			return err
		}
	}
	return nil
}
