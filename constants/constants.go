package constants

const PanicLogIdentifier = "DEVTRON_PANIC_RECOVER"

// metrics name constants
const (
	NATS_PUBLISH_COUNT          = "nats_publish_count"
	NATS_CONSUMPTION_COUNT      = "nats_consumption_count"
	NATS_CONSUMING_COUNT        = "nats_consuming_count"
	NATS_EVENT_CONSUMPTION_TIME = "nats_event_consumption_time"
	NATS_EVENT_PUBLISH_TIME     = "nats_event_publish_time"
	NATS_EVENT_DELIVERY_COUNT   = "nats_event_delivery_count"
	PANIC_RECOVERY_COUNT        = "panic_recovery_count"
)

// metrics lables constant
const (
	PANIC_TYPE = "panicType"
	HOST       = "host"
	METHOD     = "method"
	PATH       = "path"
	TOPIC      = "topic"
	STATUS     = "status"
)
