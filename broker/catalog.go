package broker

import (
	"fmt"
)

type Catalog struct {
	Services []Service `json:"services,omitempty"`
}

type Service struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Bindable        bool             `json:"bindable,omitempty"`
	Tags            []string         `json:"tags,omitempty"`
	Metadata        *ServiceMetadata `json:"metadata,omitempty"`
	Requires        []string         `json:"requires,omitempty"`
	PlanUpdateable  bool             `json:"plan_updateable"`
	Plans           []ServicePlan    `json:"plans,omitempty"`
	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`
}

type ServiceMetadata struct {
	DisplayName         string `json:"displayName,omitempty"`
	ImageURL            string `json:"imageUrl,omitempty"`
	LongDescription     string `json:"longDescription,omitempty"`
	ProviderDisplayName string `json:"providerDisplayName,omitempty"`
	DocumentationURL    string `json:"documentationUrl,omitempty"`
	SupportURL          string `json:"supportUrl,omitempty"`
}

type ServicePlan struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Metadata      *ServicePlanMetadata `json:"metadata,omitempty"`
	Free          bool                 `json:"free"`
	ElastiCacheProperties ElastiCacheProperties        `json:"elasticache_properties,omitempty"`
}

type ServicePlanMetadata struct {
	Bullets     []string `json:"bullets,omitempty"`
	Costs       []Cost   `json:"costs,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
}

type DashboardClient struct {
	ID          string `json:"id,omitempty"`
	Secret      string `json:"secret,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
}

type Cost struct {
	Amount map[string]interface{} `json:"amount,omitempty"`
	Unit   string                 `json:"unit,omitempty"`
}

type ElastiCacheProperties struct {
	CacheInstanceClass          string   `json:"cache_instance_class"`
	Engine                      string   `json:"engine"`
	EngineVersion               string   `json:"engine_version"`
	AutoMinorVersionUpgrade     bool     `json:"auto_minor_version_upgrade,omitempty"`
	Port                        int64    `json:"port,omitempty"`
	NumCacheNodes								int64    `json:"num_cache_nodes,omitempty"`
	CacheSecurityGroups         []string `json:"cache_security_groups,omitempty"`
	CacheSubnetGroupName        string   `json:"cache_subnet_group_name,omitempty"`
}

func (c Catalog) Validate() error {
	for _, service := range c.Services {
		if err := service.Validate(); err != nil {
			return fmt.Errorf("Validating Services configuration: %s", err)
		}
	}

	return nil
}

func (c Catalog) FindService(serviceID string) (service Service, found bool) {
	for _, service := range c.Services {
		if service.ID == serviceID {
			return service, true
		}
	}

	return service, false
}

func (c Catalog) FindServicePlan(planID string) (plan ServicePlan, found bool) {
	for _, service := range c.Services {
		for _, plan := range service.Plans {
			if plan.ID == planID {
				return plan, true
			}
		}
	}

	return plan, false
}

func (s Service) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("Must provide a non-empty ID (%+v)", s)
	}

	if s.Name == "" {
		return fmt.Errorf("Must provide a non-empty Name (%+v)", s)
	}

	if s.Description == "" {
		return fmt.Errorf("Must provide a non-empty Description (%+v)", s)
	}

	for _, servicePlan := range s.Plans {
		if err := servicePlan.Validate(); err != nil {
			return fmt.Errorf("Validating Plans configuration: %s", err)
		}
	}

	return nil
}

func (sp ServicePlan) Validate() error {
	if sp.ID == "" {
		return fmt.Errorf("Must provide a non-empty ID (%+v)", sp)
	}

	if sp.Name == "" {
		return fmt.Errorf("Must provide a non-empty Name (%+v)", sp)
	}

	if sp.Description == "" {
		return fmt.Errorf("Must provide a non-empty Description (%+v)", sp)
	}

	if err := sp.ElastiCacheProperties.Validate(); err != nil {
		return fmt.Errorf("Validating ElastiCache Properties configuration: %s", err)
	}

	return nil
}

func (eq ElastiCacheProperties) Validate() error {

	return nil
}
