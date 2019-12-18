package proto_nats

// QueueSubscriber is the interface generate to communicate with queue broker.
type QueueSubscriber interface {
	// Subscribe will register all the services into broker queue service.
	Subscribe() error

	// Unsubscribe will unregister all the services from the broker queue service.
	Unsubscribe() error
}
