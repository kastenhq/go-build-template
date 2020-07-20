package ship

import "time"

type Run struct {
	ID                         string    `json:"id"`
	RunID                      int       `json:"runId"`
	RunNumber                  int       `json:"runNumber"`
	BranchName                 string    `json:"branchName"`
	LastCommitShortDescription string    `json:"lastCommitShortDescription"`
	CommitSha                  string    `json:"commitSha"`
	StatusCode                 int       `json:"statusCode"`
	StatusMessage              string    `json:"statusMessage"`
	StartedAt                  time.Time `json:"startedAt"`
	EndedAt                    time.Time `json:"endedAt"`
	RunLengthInMS              int       `json:"runLengthInMS"`
	IsReady                    bool      `json:"isReady"`
	ProviderID                 string    `json:"providerId"`
	ProviderName               string    `json:"providerName"`
	ProviderURL                string    `json:"providerUrl"`
	ProjectURL                 string    `json:"projectURL"`
	ProviderDomain             string    `json:"providerDomain"`
	ProjectHTMLURL             string    `json:"projectHtmlUrl"`
	ScmURL                     string    `json:"scmUrl"`
	SubscriptionID             string    `json:"subscriptionId"`
	ProjectID                  string    `json:"projectId"`
	ProjectName                string    `json:"projectName"`
	SubscriptionOrgName        string    `json:"subscriptionOrgName"`
	BeforeCommitSha            string    `json:"beforeCommitSha"`
	CommitRange                string    `json:"commitRange"`
	CommitMessage              string    `json:"commitMessage"`
	CommitURL                  string    `json:"commitUrl"`
	CompareURL                 string    `json:"compareUrl"`
	BaseCommitRef              string    `json:"baseCommitRef"`
	HeadCommitRef              string    `json:"headCommitRef"`
	HeadPROrgName              string    `json:"headPROrgName"`
	BranchHead                 string    `json:"branchHead"`
	IsFork                     bool      `json:"isFork"`
	IsPrivate                  bool      `json:"isPrivate"`
	IsOrg                      bool      `json:"isOrg"`
	IsPullRequest              bool      `json:"isPullRequest"`
	IsManualRun                bool      `json:"isManualRun"`
	IsReRun                    bool      `json:"isReRun"`
	PullRequestNumber          int       `json:"pullRequestNumber"`
	PullRequestBaseBranch      string    `json:"pullRequestBaseBranch"`
	PullRequestSourceURL       string    `json:"pullRequestSourceUrl"`
	ReRunBatchID               string    `json:"reRunBatchId"`
	Type                       string    `json:"type"`
	Committer                  struct {
		Login       string `json:"login"`
		DisplayName string `json:"displayName"`
		Email       string `json:"email"`
		AvatarURL   string `json:"avatarUrl"`
	} `json:"committer"`
	LastAuthor struct {
		Email       string `json:"email"`
		Login       string `json:"login"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
	} `json:"lastAuthor"`
	TriggeredBy struct {
		Login       string `json:"login"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
	} `json:"triggeredBy"`
	SkipDecryption bool `json:"skipDecryption"`
	CleanRunYml    struct {
		Language     string `json:"language"`
		Services     string `json:"services"`
		Integrations struct {
			Hub []struct {
				IntegrationName string `json:"integrationName"`
				Type            string `json:"type"`
				IsValid         bool   `json:"isValid"`
			} `json:"hub"`
			Deploy []struct {
				IntegrationName string `json:"integrationName"`
				Type            string `json:"type"`
				Target          string `json:"target"`
				Platform        string `json:"platform"`
				ApplicationName string `json:"application_name"`
				EnvName         string `json:"env_name"`
				Region          string `json:"region"`
				IsValid         bool   `json:"isValid"`
			} `json:"deploy"`
			Key []struct {
				IntegrationName string `json:"integrationName"`
				Type            string `json:"type"`
				IsValid         bool   `json:"isValid"`
			} `json:"key"`
			Notifications []struct {
				IntegrationName string   `json:"integrationName"`
				Type            string   `json:"type"`
				Recipients      []string `json:"recipients"`
				Branches        struct {
					Only   []string `json:"only"`
					Except []string `json:"except"`
				} `json:"branches"`
				OnSuccess     string `json:"on_success"`
				OnFailure     string `json:"on_failure"`
				OnPullRequest string `json:"on_pull_request"`
				OnStart       string `json:"on_start"`
				IsValid       bool   `json:"isValid"`
			} `json:"notifications"`
		} `json:"integrations"`
		Git struct {
			Submodules bool `json:"submodules"`
		} `json:"git"`
		Branches struct {
			Only   []string `json:"only"`
			Except []string `json:"except"`
		} `json:"branches"`
		Skip bool `json:"skip"`
		Env  struct {
			Global   []string `json:"global"`
			Matrix   []string `json:"matrix"`
			Injected []string `json:"injected"`
		} `json:"env"`
		Jdk         []string `json:"jdk"`
		Gemfile     string   `json:"gemfile"`
		BundlerArgs []string `json:"bundler_args"`
		Addons      []struct {
		} `json:"addons"`
		Compiler []string `json:"compiler"`
		Matrix   struct {
			Include struct {
			} `json:"include"`
			Exclude struct {
			} `json:"exclude"`
			AllowFailures struct {
			} `json:"allow failures"`
		} `json:"matrix"`
		Build struct {
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
		Versions []string `json:"versions"`
	} `json:"cleanRunYml"`
	Language                string    `json:"language"`
	ParallelizedTest        bool      `json:"parallelizedTest"`
	TotalTests              int       `json:"totalTests"`
	TestsFailed             int       `json:"testsFailed"`
	TestsPassed             int       `json:"testsPassed"`
	TestsSkipped            int       `json:"testsSkipped"`
	TestsErrors             int       `json:"testsErrors"`
	SequenceCoveragePercent int       `json:"sequenceCoveragePercent"`
	BranchCoveragePercent   int       `json:"branchCoveragePercent"`
	ErrorMsgs               []string  `json:"errorMsgs"`
	WarnMsgs                []string  `json:"warnMsgs"`
	IsGitTag                bool      `json:"isGitTag"`
	GitTagName              string    `json:"gitTagName"`
	GitTagMessage           string    `json:"gitTagMessage"`
	IsRelease               bool      `json:"isRelease"`
	IsPrerelease            bool      `json:"isPrerelease"`
	ReleaseName             string    `json:"releaseName"`
	ReleaseBody             string    `json:"releaseBody"`
	ReleasedAt              time.Time `json:"releasedAt"`
	CreatedBy               string    `json:"createdBy"`
	UpdatedBy               string    `json:"updatedBy"`
}
