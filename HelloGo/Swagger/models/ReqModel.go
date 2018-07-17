package models

type GetUsrReq struct {
	ID string `json:"id,omitempty"`
}

type GetUsrResp struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateUsrReq struct {
	ID  string `json:"id,omitempty"`
	Act string `json:"act,omitempty"`
	Pwd string `json:"pwd,omitempty"`
}

type CreateUsrResp struct {
	ID  string `json:"id,omitempty"`
	Act string `json:"act,omitempty"`
	Pwd string `json:"pwd,omitempty"`
}
