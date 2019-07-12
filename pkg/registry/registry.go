package registry

type Service struct {
}

type Target struct {
	Address string
}

type ServiceUpdateType int

const (
	ServiceUpdateTypeUnkown ServiceUpdateType = 0
	ServiceUpdateTypeAdd    ServiceUpdateType = 1
	ServiceUpdateTypeRemove ServiceUpdateType = 2
)

type ServiceUpdate struct {
	Type    ServiceUpdateType
	Service *Service
	Target  *Target
}

type Registry interface {
	Name() string
	Init() error
	WatchServiceUpdates() chan ServiceUpdate
	Shutdown() error
}
