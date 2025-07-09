package api_interfaces

import "net/http"

type HealthCheckApi interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	EmailCheck(w http.ResponseWriter, r *http.Request)
}
