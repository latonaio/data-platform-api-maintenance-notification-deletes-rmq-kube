package requests

type Header struct {
	MaintenanceNotification int   `json:"MaintenanceNotification"`
	IsMarkedForDeletion     *bool `json:"IsMarkedForDeletion"`
}
