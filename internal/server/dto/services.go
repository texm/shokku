package dto

type ManageServiceRequest struct {
	Name string `json:"name" validate:"appName"`
	Type string `json:"type" validate:"alphanum"`
}

type GenericServiceRequest struct {
	Name string `query:"name" validate:"appName"`
	Type string `query:"type" validate:"alphanum"`
}

type GenericServiceCreationConfig struct {
	ConfigOptions *string `json:"config-options"`
	// validate inner pairs are len=2
	CustomEnv        *[][]string `json:"custom-env"`
	Image            *string     `json:"image"`
	ImageVersion     *string     `json:"image-version"`
	MemoryLimit      *string     `json:"memory"`
	Password         *string     `json:"password"`
	RootPassword     *string     `json:"root-password"`
	SharedMemorySize *string     `json:"shm-size"`
}

type CreateGenericServiceRequest struct {
	Name        string                       `json:"name" validate:"appName"`
	ServiceType string                       `json:"type"`
	Config      GenericServiceCreationConfig `json:"config"`
}

type CloneServiceRequest struct {
	Name    string `json:"name" validate:"appName"`
	NewName string `json:"newName" validate:"appName"`
}

type ServiceInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ListServicesResponse struct {
	Services []ServiceInfo `json:"services"`
}

type GetServiceInfoResponse struct {
	Info map[string]string `json:"info"`
}

type GetServiceTypeRequest struct {
	Name string `query:"name" validate:"appName"`
}
type GetServiceTypeResponse struct {
	Type string `json:"type"`
}

type LinkGenericServiceToAppRequest struct {
	ServiceName string `json:"service_name" validate:"appName"`
	AppName     string `json:"app_name" validate:"appName"`
	Alias       string `json:"alias"`
	QueryString string `json:"query_string"`
}

type GetServiceLinkedAppsResponse struct {
	Apps []string `json:"apps"`
}

type GetServiceLogsResponse struct {
	Logs []string `json:"logs"`
}

type GetServiceBackupReportRequest struct {
	Name string `query:"name" validate:"appName"`
}

type ServiceBackupReport struct {
	AuthSet       bool   `json:"auth_set"`
	EncryptionSet bool   `json:"encryption_set"`
	Bucket        string `json:"bucket"`
	Schedule      string `json:"schedule"`
}
type GetServiceBackupReportResponse struct {
	Report ServiceBackupReport `json:"report"`
}

type RunServiceBackupRequest struct {
	Name string `query:"name" validate:"appName"`
}

type BackupsAuthConfig struct {
	AccessKeyId      string `json:"access_key_id"`
	SecretKey        string `json:"secret_key"`
	Region           string `json:"region"`
	SignatureVersion string `json:"signature_version"`
	EndpointUrl      string `json:"endpoint_url"`
}
type SetServiceBackupsAuthRequest struct {
	Name   string            `json:"name" validate:"appName"`
	Config BackupsAuthConfig `json:"config"`
}

type SetServiceBackupsBucketRequest struct {
	Name   string `json:"name" validate:"appName"`
	Bucket string `json:"bucket"`
}

type SetServiceBackupsScheduleRequest struct {
	Name     string `json:"name" validate:"appName"`
	Schedule string `json:"schedule"`
}

type SetServiceBackupsEncryptionRequest struct {
	Name       string `json:"name" validate:"appName"`
	Passphrase string `json:"passphrase"`
}
