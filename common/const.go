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
	//PortGRPCUserService = ":50051"
	//HostRPCUserService  = "localhost" + PortGRPCUserService
	//PortGinUserService  = ":8081"

	//PortGRPCUAuthenService = ":50051"
	PortgRPCUserFriendService = ":50052"
	HostgRPCUserFriendService = "localhost" + PortgRPCUserFriendService
	//PortGinPostService        = ":8082"

	PortgRPCPostService = ":50051"
	HostgRPCPostService = "localhost" + PortgRPCPostService
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	//GetRole() string
}
