package common

const (
	DbTypeNote   = 1
	DbTypeUser   = 2
	DbTypeUpload = 3

	CurrentUser = "current_user"
)

const (
	TopicPostCreated = "TopicPostCreated"
)

const (
	PortGRPCUserService = ":50051"
	HostRPCUserService  = "localhost" + PortGRPCUserService
	PortGinUserService  = ":8081"

	PortGinPostService = ":8082"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	//GetRole() string
}
