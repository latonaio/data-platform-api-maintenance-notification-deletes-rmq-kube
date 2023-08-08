package requests

type Item struct {
	MaintenanceNotification     int   `json:"MaintenanceNotification"`
	MaintenanceNotificationItem int   `json:"MaintenanceNotificationItem"`
	IsMarkedForDeletion         *bool `json:"IsMarkedForDeletion"`
}
