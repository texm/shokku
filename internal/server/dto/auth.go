package dto

type PasswordLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TotpCode string `json:"totp"`
}

type PasswordLoginResponse struct {
	Success   bool `json:"success"`
	NeedsTotp bool `json:"needs_totp"`
}

type GithubAuthRequest struct {
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_url"`
}

type GetGithubSetupStatus struct {
	AppCreated bool `json:"created"`
}

type CreateGithubAppRequest struct {
	Code string `json:"code"`
}

type CreateGithubAppResponse struct {
	Slug string `json:"slug"`
}

type InstallGithubAppResponse struct {
	InstallURL string `json:"install_url"`
}

type CompleteGithubSetupRequest struct {
	Code           string `json:"code"`
	InstallationId int64  `json:"installation_id"`
}

type GetGithubAuthInfoResponse struct {
	ClientID string `json:"client_id"`
}

type GenerateTotpResponse struct {
	Secret       string `json:"secret"`
	Image        string `json:"image"`
	RecoveryCode string `json:"recovery_code"`
}

type ConfirmTotpRequest struct {
	Secret string `json:"secret"`
	Code   string `json:"code"`
}

type ConfirmTotpResponse struct {
	Valid bool `json:"valid"`
}

type CompletePasswordSetupRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Enable2FA    bool   `json:"enable_2fa"`
	TotpSecret   string `json:"totp_secret"`
	RecoveryCode string `json:"recovery_code"`
}
