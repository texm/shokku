package dto

import (
	"github.com/texm/dokku-go"
	"time"
)

type GetAppOverviewRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppOverviewResponse struct {
	Name            string `json:"name,omitempty"`
	IsSetup         bool   `json:"is_setup"`
	SetupMethod     string `json:"setup_method"`
	GitDeployBranch string `json:"git_deploy_branch"`
	GitLastUpdated  string `json:"git_last_updated"`
	IsDeployed      bool   `json:"is_deployed"`
	IsRunning       bool   `json:"is_running"`
	NumProcesses    int    `json:"num_processes"`
	CanScale        bool   `json:"can_scale"`
	Restore         bool   `json:"restore"`
}

type GetAllAppsOverviewResponse struct {
	Apps []GetAppOverviewResponse `json:"apps"`
}

type GetAppsListItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type GetAppsListResponse struct {
	Apps []GetAppsListItem `json:"apps"`
}

type GetAppInfoRequest struct {
	Name string `query:"name" validate:"appName"`
}

type DestroyAppRequest struct {
	Name string `json:"name"`
}

type AppInfo struct {
	Name                 string    `json:"name"`
	Directory            string    `json:"directory"`
	DeploySource         string    `json:"deploy_source"`
	DeploySourceMetadata string    `json:"deploy_source_metadata"`
	CreatedAt            time.Time `json:"created_at"`
	IsLocked             bool      `json:"is_locked"`
}

type GetAppInfoResponse struct {
	Info AppInfo `json:"info"`
}

type ManageAppRequest struct {
	Name string `json:"name" validate:"appName"`
}

type GetAppSetupStatusRequest struct {
	Name string `json:"name" validate:"appName"`
}
type GetAppSetupStatusResponse struct {
	IsSetup bool   `json:"is_setup"`
	Method  string `json:"method"`
}

type GetAppSetupConfigRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppSetupConfigResponse struct {
	IsSetup          bool   `json:"is_setup"`
	Method           string `json:"method"`
	DeploymentBranch string `json:"deployment_branch,omitempty"`
	RepoURL          string `json:"repo_url,omitempty"`
	RepoGitRef       string `json:"repo_git_ref,omitempty"`
	Image            string `json:"image,omitempty"`
}

type SetupAppNewRepoRequest struct {
	Name             string `json:"name" validate:"appName"`
	DeploymentBranch string `json:"deployment_branch"`
}

type SetupAppSyncRepoRequest struct {
	Name          string `json:"name" validate:"appName"`
	RepositoryURL string `json:"repository_url"`
	GitRef        string `json:"git_ref"`
}

type SetupAppPullImageRequest struct {
	Name  string `json:"name" validate:"appName"`
	Image string `json:"image"`
}

type SetupAppUploadArchiveRequest struct {
	Name string `form:"name" validate:"appName"`
}

type RenameAppRequest struct {
	CurrentName string `json:"current_name" validate:"appName"`
	NewName     string `json:"new_name" validate:"appName"`
}

/*
		methods = ["Git Push", "Git Repository", "Archive File", "Dockerfile", "Docker Image"]
		options = [["deploymentBranch", "envVar"], ["repositoryURL", "gitRef"], ["file"],
	 			   ["dockerfilePath", "usingBuildkit"], ["image"]]
*/
type DeployAppRequest struct {
	Name    string            `json:"name" validate:"appName"`
	Method  string            `json:"method" validate:"alpha"`
	Options map[string]string `json:"options" validate:"alpha"`
}

type GetAppServicesRequest struct {
	Name string `query:"name" validate:"appName"`
}

type GetAppDeployChecksRequest struct {
	Name string `query:"name" validate:"appName"`
}

type GetAppDeployChecksResponse struct {
	AllDisabled       bool     `json:"all_disabled"`
	AllSkipped        bool     `json:"all_skipped"`
	DisabledProcesses []string `json:"disabled_processes"`
	SkippedProcesses  []string `json:"skipped_processes"`
}

type SetAppDeployChecksRequest struct {
	Name string `json:"name" validate:"appName"`
	// enabled, disabled, skipped
	State string `json:"state" validate:"alpha"`
}

type SetAppProcessDeployChecksRequest struct {
	Name    string `json:"name" validate:"appName"`
	Process string `json:"process" validate:"processName"`
	// enabled, disabled, skipped
	State string `json:"state" validate:"alpha"`
}

type GetAppProcessesRequest struct {
	Name string `query:"name" validate:"appName"`
}

type GetAppProcessesResponse struct {
	Processes []string `json:"processes"`
}

type GetAppProcessReportRequest struct {
	Name string `query:"name" validate:"appName"`
}

type AppProcessInfo struct {
	Scale     int                    `json:"scale"`
	Resources dokku.ResourceSettings `json:"resources"`
}

type GetAppProcessReportResponse struct {
	ResourceDefaults dokku.ResourceSettings    `json:"resource_defaults"`
	Processes        map[string]AppProcessInfo `json:"processes"`
}

type AppResources struct {
	CPU        *int    `json:"cpu"`
	Memory     *int    `json:"memory"`
	MemoryUnit *string `json:"memory_unit"`
}

type SetAppProcessResourcesRequest struct {
	Name                 string       `json:"name" validate:"appName"`
	Process              string       `json:"process" validate:"processName"`
	ResourceLimits       AppResources `json:"limits"`
	ResourceReservations AppResources `json:"reservations"`
}

type GetAppProcessScaleRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppProcessScaleResponse struct {
	ProcessScale map[string]int `json:"process_scale"`
}

type SetAppProcessScaleRequest struct {
	Name       string `json:"name" validate:"appName"`
	Process    string `json:"process" validate:"processName"`
	Scale      int    `json:"scale" validate:"numeric"`
	SkipDeploy bool   `json:"skip_deploy"`
}

type GetAppDomainsReportRequest struct {
	Name string `query:"name" validate:"appName"`
}

type GetAppDomainsReportResponse struct {
	Domains []string `json:"domains"`
	Enabled bool     `json:"enabled"`
}

type SetAppDomainsEnabledRequest struct {
	Name    string `json:"name" validate:"appName"`
	Enabled bool   `json:"enabled"`
}

type GetAppLetsEncryptEnabledRequest struct {
	Name string `query:"name" validate:"appName"`
}
type SetAppLetsEncryptEnabledRequest struct {
	Name    string `json:"name" validate:"appName"`
	Enabled bool   `json:"enabled"`
}

type GetAppDomainsRequest struct {
	Name string `query:"name" validate:"appName"`
}
type AlterAppDomainRequest struct {
	Name   string `json:"name" validate:"appName"`
	Domain string `json:"domain" validate:"hostname_rfc1123"`
}

type AlterNetworkRequest struct {
	Network string `query:"network"`
}

type ListNetworksResponse struct {
	Networks []string `json:"networks"`
}

type GetAppNetworksReportRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppNetworksReportResponse struct {
	AttachInitial     string `json:"attach_initial"`
	AttachPostCreate  string `json:"attach_post_create"`
	AttachPostDeploy  string `json:"attach_post_deploy"`
	BindAllInterfaces bool   `json:"bind_all_interfaces"`
	TLD               string `json:"tld"`
	WebListeners      string `json:"web_listeners"`
}

type SetAppNetworksRequest struct {
	Name string `query:"name" validate:"appName"`

	Initial           *string `json:"attach_initial"`
	PostCreate        *string `json:"attach_post_create"`
	PostDeploy        *string `json:"attach_post_deploy"`
	BindAllInterfaces *bool   `json:"bind_all_interfaces"`
	TLD               *string `json:"tld"`
}

type GetAppLogsRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppLogsResponse struct {
	Logs []string `json:"logs"`
}

type GetAppConfigRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppConfigResponse struct {
	Config map[string]string `json:"config"`
}

type SetAppConfigRequest struct {
	Name string `json:"name" validate:"appName"`
	Config map[string]string `json:"config"`
}

type GetAppStorageRequest struct {
	Name string `query:"name" validate:"appName"`
}
type StorageMount struct {
	HostDir       string `json:"hostDir"`
	ContainerDir  string `json:"mountDir"`
	IsBuildMount  bool   `json:"isBuildMount"`
	IsRunMount    bool   `json:"isRunMount"`
	IsDeployMount bool   `json:"isDeployMount"`
}
type GetAppStorageResponse struct {
	Mounts []StorageMount `json:"mounts"`
}

type AlterAppStorageRequest struct {
	Name       string `json:"name" validate:"appName"`
	RestartApp bool   `json:"restart"`

	HostDir      string `json:"hostDir"`
	ContainerDir string `json:"mountDir"`
}

type GetAppBuilderRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppBuilderResponse struct {
	Selected string `json:"selected"`
}

type SetAppBuilderRequest struct {
	Name    string `json:"name" validate:"appName"`
	Builder string `json:"builder" validate:"alphanum"`
}

type GetAppBuildDirectoryRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetAppBuildDirectoryResponse struct {
	Directory string `json:"directory"`
}

type SetAppBuildDirectoryRequest struct {
	Name      string `json:"name" validate:"appName"`
	Directory string `json:"directory" validate:"alphanum"`
}

type ClearAppBuildDirectoryRequest struct {
	Name string `query:"name" validate:"appName"`
}
