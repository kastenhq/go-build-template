package ship

import "time"

type BuildJob struct {
	ID                string    `json:"id"`
	BuildJobID        int       `json:"buildJobId"`
	BuildJobNumber    int       `json:"buildJobNumber"`
	BuildID           string    `json:"buildId"`
	StatusCode        int       `json:"statusCode"`
	SubscriptionID    string    `json:"subscriptionId"`
	ResourceID        int       `json:"resourceId"`
	VersionID         int       `json:"versionId"`
	IsSerial          bool      `json:"isSerial"`
	IsCompleted       bool      `json:"isCompleted"`
	IsConsoleArchived bool      `json:"isConsoleArchived"`
	StartedAt         time.Time `json:"startedAt"`
	EndedAt           time.Time `json:"endedAt"`
	JobLengthMS       int       `json:"jobLengthMS"`
	NodeID            string    `json:"nodeId"`
	//TimeoutMS         string    `json:"timeoutMS"`
	ProjectID   string    `json:"projectId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	PropertyBag struct {
		Payload struct {
			ProjectID      string `json:"projectId"`
			ResourceID     int    `json:"resourceId"`
			Name           string `json:"name"`
			SourceName     string `json:"sourceName"`
			SubscriptionID string `json:"subscriptionId"`
			Type           string `json:"type"`
			Force          bool   `json:"force"`
			PropertyBag    struct {
				Yml struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Steps []struct {
						IN string `json:"IN"`
					} `json:"steps"`
					Flags []string `json:"flags"`
				} `json:"yml"`
			} `json:"propertyBag"`
			Dependencies []struct {
				Operation    string `json:"operation"`
				ResourceID   int    `json:"resourceId"`
				Name         string `json:"name"`
				SourceName   string `json:"sourceName"`
				ProjectID    string `json:"projectId"`
				IsConsistent bool   `json:"isConsistent"`
				Type         string `json:"type"`
				PropertyBag  struct {
					Yml struct {
						Name string `json:"name"`
						Type string `json:"type"`
						Seed struct {
							VersionName string `json:"versionName"`
						} `json:"seed"`
						Flags []string `json:"flags"`
					} `json:"yml"`
				} `json:"propertyBag"`
				IsJob                        bool `json:"isJob"`
				VersionDependencyPropertyBag struct {
				} `json:"versionDependencyPropertyBag"`
				Version struct {
					VersionID     int    `json:"versionId"`
					VersionNumber int    `json:"versionNumber"`
					VersionName   string `json:"versionName"`
					PropertyBag   struct {
					} `json:"propertyBag"`
				} `json:"version"`
			} `json:"dependencies"`
			BuildID     string `json:"buildId"`
			BuildNumber int    `json:"buildNumber"`
		} `json:"payload"`
		SubscriptionIntegrationIds []int `json:"subscriptionIntegrationIds"`
	} `json:"propertyBag"`
}
