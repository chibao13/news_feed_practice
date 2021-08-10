package main

import (
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/memcache"
	"github.com/chibao13/news_feed_practice/pubsub/pblocal"
	"github.com/chibao13/news_feed_practice/services/apigateway/cache"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/userstorage"
	subscriber "github.com/chibao13/news_feed_practice/subscribe"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	sysSecret := os.Getenv("SYSTEM_SECRET")
	dns := os.Getenv("DB_CONN")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, sysSecret); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, secretKey string) error {
	r := gin.Default()
	pb := pblocal.NewPubSub()

	appCtx := appctx.New(db, secretKey, pb)
	//rtEngine := skio.NewEngine()
	//_ = rtEngine.Run(appCtx, r)
	//subscriber.NewEngine(appCtx).Start(rtEngine)
	subscriber.NewEngine(appCtx).Start()

	userStore := userstorage.NewSQLStore(db)
	userCachingStore := cache.NewUserCaching(memcache.NewCaching(), userStore)
	setupRouter(r, appCtx, userCachingStore)

	go func() {
		if err := runRpcService(appCtx); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := r.Run(); err != nil {
		return err
	}
	return nil
}

func runRpcService(ctx appctx.AppContext) error {
	//lis, err := net.Listen("tcp", common.PortGRPCUserService)
	//if err != nil {
	//	return err
	//}
	//s := grpc.NewServer()
	//api.RegisterUserRPCServer(s, grpcuser.NewGrpcUser(ctx))
	//if err := s.Serve(lis); err != nil {
	//	return err
	//}
	return nil
}
