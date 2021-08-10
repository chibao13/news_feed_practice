package userfriendbiz

import "context"

type ListFriendStorage interface {
	ListFriendIdsWithCondition(ctx context.Context, cond map[string]interface{}) ([]uint32, error)
}

type listFriendBussiess struct {
	storage ListFriendStorage
}

func NewListFriendBusiness(storage ListFriendStorage) *listFriendBussiess {
	return &listFriendBussiess{storage: storage}
}

func (biz *listFriendBussiess) ListFriendIds(ctx context.Context, conditions map[string]interface{}) ([]uint32, error) {
	return biz.storage.ListFriendIdsWithCondition(ctx, conditions)
}
