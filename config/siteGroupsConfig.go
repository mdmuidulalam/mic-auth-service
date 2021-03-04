package config

import (
	"github.com/tkanos/gonfig"
)

type SiteGroup struct {
	key      string
	database struct {
		name     string
		host     string
		port     string
		user     string
		password string
	}
}

type SiteGroupConfig struct {
	siteGroupsMap *map[string]*SiteGroup
}

func (sgConfig *SiteGroupConfig) SetSiteGroupConfig() {
	var siteGroups []*SiteGroup
	if err := gonfig.GetConf("siteGroups.json", siteGroups); err != nil {
		panic(err)
	}

	sgConfig.siteGroupsMap = &map[string]*SiteGroup{}
	for _, siteGroup := range siteGroups {
		(*sgConfig.siteGroupsMap)[siteGroup.key] = siteGroup
	}
}

func (sgConfig *SiteGroupConfig) GetSiteGroupConfig() *map[string]*SiteGroup {
	if len(*sgConfig.siteGroupsMap) == 0 {
		sgConfig.SetSiteGroupConfig()
	}

	return sgConfig.siteGroupsMap
}
