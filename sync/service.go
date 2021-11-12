package sync

// Service is an interface that provides the common API for all sync services.
type Service interface {
	Run() error
}
