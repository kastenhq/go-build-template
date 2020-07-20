package ship

import "time"

type Job struct {
	ID                    string   `json:"id"`
	JobID                 int      `json:"jobId"`
	RunID                 string   `json:"runId"`
	RunNumber             int      `json:"runNumber"`
	JobNumber             int      `json:"jobNumber"`
	ProjectID             string   `json:"projectId"`
	ProjectName           string   `json:"projectName"`
	SubscriptionID        string   `json:"subscriptionId"`
	SubscriptionOrgName   string   `json:"subscriptionOrgName"`
	BranchName            string   `json:"branchName"`
	GitTagName            string   `json:"gitTagName"`
	ReleaseName           string   `json:"releaseName"`
	PullRequestBaseBranch string   `json:"pullRequestBaseBranch"`
	IsPullRequest         bool     `json:"isPullRequest"`
	IsGitTag              bool     `json:"isGitTag"`
	IsRelease             bool     `json:"isRelease"`
	IsPrerelease          bool     `json:"isPrerelease"`
	IsPrivateRepository   bool     `json:"isPrivateRepository"`
	PullRequestNumber     int      `json:"pullRequestNumber"`
	JobLengthInMS         int      `json:"jobLengthInMS"`
	TimeoutMS             int      `json:"timeoutMS"`
	Version               string   `json:"version"`
	Env                   []string `json:"env"`
	BundlerArgs           []string `json:"bundlerArgs"`
	Services              []string `json:"services"`
	Addons                []struct {
		Firefox string `json:"firefox"`
	} `json:"addons"`
	Gemfile          string `json:"gemfile"`
	Compiler         string `json:"compiler"`
	Jdk              string `json:"jdk"`
	SubmoduleEnabled bool   `json:"submoduleEnabled"`
	AllowFailure     bool   `json:"allowFailure"`
	Build            struct {
		PreCi     []string `json:"pre_ci"`
		PreCiBoot struct {
			ImageName       string `json:"image_name"`
			ImageTag        string `json:"image_tag"`
			Envs            string `json:"envs"`
			Options         string `json:"options"`
			IsOfficialImage bool   `json:"isOfficialImage"`
			IsLegacyImage   bool   `json:"isLegacyImage"`
			ContainerName   string `json:"containerName"`
			SectionName     string `json:"sectionName"`
		} `json:"pre_ci_boot"`
		Ci        []string `json:"ci"`
		PostCi    []string `json:"post_ci"`
		OnSuccess []string `json:"on_success"`
		OnFailure []string `json:"on_failure"`
		Cache     bool     `json:"cache"`
	} `json:"build"`
	Deploy struct {
	} `json:"deploy"`
	Infra struct {
	} `json:"infra"`
	Steps []struct {
		ID         string `json:"id"`
		ExecOrder  int    `json:"execOrder"`
		ScriptType string `json:"scriptType"`
		Who        string `json:"who"`
		Script     string `json:"script"`
	} `json:"steps"`
	OnStartJobEnvs []struct {
		AccountIntegrationID string   `json:"accountIntegrationId"`
		EnvVars              []string `json:"envVars"`
	} `json:"onStartJobEnvs"`
	PostJobEnvs []struct {
		AccountIntegrationID string   `json:"accountIntegrationId"`
		EnvVars              []string `json:"envVars"`
	} `json:"postJobEnvs"`
	JobErrorMsgs            []interface{} `json:"jobErrorMsgs"`
	IsCompleted             bool          `json:"isCompleted"`
	IsConsoleArchived       bool          `json:"isConsoleArchived"`
	TotalTests              int           `json:"totalTests"`
	TestsFailed             int           `json:"testsFailed"`
	TestsPassed             int           `json:"testsPassed"`
	TestsSkipped            int           `json:"testsSkipped"`
	TestsErrors             int           `json:"testsErrors"`
	BranchCoveragePercent   int           `json:"branchCoveragePercent"`
	SequenceCoveragePercent int           `json:"sequenceCoveragePercent"`
	Type                    string        `json:"type"`
	StatusCode              int           `json:"statusCode"`
	StatusMessage           string        `json:"statusMessage"`
	Node                    string        `json:"node"`
	CommitSha               string        `json:"commitSha"`
	ScmURL                  string        `json:"scmUrl"`
	PullRequestSourceURL    string        `json:"pullRequestSourceUrl"`
	StartedAt               time.Time     `json:"startedAt"`
	CreatedAt               time.Time     `json:"createdAt"`
	EndedAt                 time.Time     `json:"endedAt"`
	ReleasedAt              time.Time     `json:"releasedAt"`
	CreatedBy               string        `json:"createdBy"`
	UpdatedBy               string        `json:"updatedBy"`
}
