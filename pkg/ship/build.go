package ship

import "time"

type Build struct {
	ID               string    `json:"id"`
	BuildID          int       `json:"buildId"`
	ExternalBuildID  string    `json:"externalBuildId"`
	StatusCode       int       `json:"statusCode"`
	ResourceID       int       `json:"resourceId"`
	VersionID        int       `json:"versionId"`
	SubscriptionID   string    `json:"subscriptionId"`
	ProjectID        string    `json:"projectId"`
	ExternalBuildURL string    `json:"externalBuildUrl"`
	EndedAt          time.Time `json:"endedAt"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
