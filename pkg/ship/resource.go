package ship

import "time"

type Resource struct {
	ID                        int       `json:"id"`
	Name                      string    `json:"name"`
	SourceName                string    `json:"sourceName"`
	SubscriptionIntegrationID int       `json:"subscriptionIntegrationId"`
	TypeCode                  int       `json:"typeCode"`
	ArchTypeCode              int       `json:"archTypeCode"`
	OperatingSystemCode       int       `json:"operatingSystemCode"`
	LastVersionID             int       `json:"lastVersionId"`
	LastVersionName           string    `json:"lastVersionName"`
	LastVersionNumber         int       `json:"lastVersionNumber"`
	SubscriptionID            string    `json:"subscriptionId"`
	ProjectID                 string    `json:"projectId"`
	ParentResourceID          int       `json:"parentResourceId"`
	IsJob                     bool      `json:"isJob"`
	IsDeleted                 bool      `json:"isDeleted"`
	IsConsistent              bool      `json:"isConsistent"`
	NextTriggerTime           time.Time `json:"nextTriggerTime"`
	PropertyBag               struct {
		Yml struct {
			Name        string   `json:"name"`
			Type        string   `json:"type"`
			Integration string   `json:"integration"`
			Flags       []string `json:"flags"`
			Steps       []struct {
				IN     string `json:"IN"`
				OUT    string `json:"OUT"`
				NOTIFY string `json:"NOTIFY"`
				TASK   string `json:"TASK"`
			} `json:"steps"`
		} `json:"yml"`
		SysIntegrationName string `json:"sysIntegrationName"`
	} `json:"propertyBag"`
	CreatedBy time.Time `json:"createdBy"`
	UpdatedBy time.Time `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
