package config

import (
	"encoding/json"
	"io/ioutil"
)

type siteGroup struct {
	Key      string
	Database struct {
		Name     string
		Host     string
		Port     string
		User     string
		Password string
	}
}

type SiteGroupsConfig struct {
	siteGroupsMap *map[string]*siteGroup
}

func (sgConfig *SiteGroupsConfig) SetSiteGroupConfig() {
	fileContent, err := ioutil.ReadFile("config/siteGroups.json")
	if err != nil {
		panic(err)
	}

	var siteGroups []*siteGroup
	if err = json.Unmarshal(fileContent, &siteGroups); err != nil {
		panic(err)
	}

	sgConfig.siteGroupsMap = &map[string]*siteGroup{}
	for _, siteGroup := range siteGroups {
		(*sgConfig.siteGroupsMap)[siteGroup.Key] = siteGroup
	}
}

func (sgConfig *SiteGroupsConfig) GetSiteGroupConfig() *map[string]*siteGroup {
	if sgConfig.siteGroupsMap == nil || len(*sgConfig.siteGroupsMap) == 0 {
		sgConfig.SetSiteGroupConfig()
	}

	return sgConfig.siteGroupsMap
}
