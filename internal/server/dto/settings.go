package dto

import "github.com/texm/dokku-go"

type GetVersionsResponse struct {
	Dokku  string `json:"dokku"`
	Shokku string `json:"shokku"`
}

type GetLetsEncryptStatusResponse struct {
	Installed bool   `json:"installed"`
	Email     string `json:"email"`
}

type User struct {
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	SSHKeys []string `json:"ssh_keys"`
}
type GetUsersResponse struct {
	Users []User `json:"users"`
}

type GetSSHKeysResponse struct {
	Keys []dokku.SSHKey `json:"keys"`
}

type GetGlobalDomainsResponse struct {
	Domains []string `json:"domains"`
	Enabled bool     `json:"enabled"`
}
type AlterGlobalDomainRequest struct {
	Domain string `json:"domain"`
}
type DeleteGlobalDomainRequest struct {
	Domain string `query:"domain"`
}

type AddGitAuthRequest struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RemoveGitAuthRequest struct {
	Host string `json:"host"`
}

type SetDockerRegistryRequest struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetDockerRegistryReportResponse struct {
	Server        string `json:"server"`
	PushOnRelease bool   `json:"push_on_release"`
}

type SetEventLoggingEnabledRequest struct {
	Enabled bool `json:"enabled"`
}
type GetEventLogsListResponse struct {
	Events []string `json:"events"`
}
type GetEventLogsResponse struct {
	Logs string `json:"logs"`
}

type PluginInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}
type ListPluginsResponse struct {
	Plugins []PluginInfo `json:"plugins"`
}
