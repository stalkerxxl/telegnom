package types

import "time"

// WebhookInfo describes the current status of a webhook.
//
// See https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                          string       `json:"url"`
	HasCustomCertificate         bool         `json:"has_custom_certificate"`
	PendingUpdateCount           int          `json:"pending_update_count"`
	IPAddress                    string       `json:"ip_address,omitempty"`
	LastErrorDate                int64        `json:"last_error_date,omitempty"`
	LastErrorMessage             string       `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64        `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int          `json:"max_connections,omitempty"`
	AllowedUpdates               []UpdateType `json:"allowed_updates,omitempty"`
}

// LastErrorTime returns the time of the last error.
func (wi *WebhookInfo) LastErrorTime() time.Time {
	return time.Unix(wi.LastErrorDate, 0)
}

// LastSynchronizationErrorTime returns the time of the last synchronization error.
func (wi *WebhookInfo) LastSynchronizationErrorTime() time.Time {
	return time.Unix(wi.LastSynchronizationErrorDate, 0)
}
