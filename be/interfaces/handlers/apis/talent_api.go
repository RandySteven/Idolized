package api_interfaces

import "net/http"

type TalentApi interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	GetListTalents(w http.ResponseWriter, r *http.Request)
}
