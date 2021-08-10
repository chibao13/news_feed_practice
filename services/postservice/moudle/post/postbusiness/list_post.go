package postbusiness

import (
	"context"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postmodel"
)

type ListPostStorage interface {
	ListPostWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging, moredata ...string) ([]postmodel.Post, error)
}

//type listPostBusiness struct {
//	store  ListPostStorage
//	appCtx appctx.AppContext
//}
//
//func NewListPostBusiness(storage ListPostStorage, appCtx appctx.AppContext) *listPostBusiness {
//	return &listPostBusiness{store: storage, appCtx: appCtx}
//}

//func (biz *listPostBusiness) ListPostWithCondition(ctx context.Context, conditions map[string]interface{}, paging *common.Paging) ([]postmodel.Post, error) {
//
//	result, err := biz.store.ListPostWithCondition(ctx, conditions, paging, "User")
//
//	if err != nil {
//		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
//	}
//	return result, nil
//}
