package dto

type GetSetupStatusResponse struct {
	IsSetup bool   `json:"is_setup"`
	Method  string `json:"method"`
}
