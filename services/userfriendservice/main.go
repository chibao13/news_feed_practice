package main

import (
	api "github.com/chibao13/news_feed_practice/api/v1/userfriend"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/pubsub/pblocal"
	"github.com/chibao13/news_feed_practice/services/userfriendservice/moudle/userfriend/userfriendtransport/grpcuserfriend"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
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
	pb := pblocal.NewPubSub()

	appCtx := appctx.New(db, secretKey, pb)
	lis, err := net.Listen("tcp", common.PortgRPCUserFriendService)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	log.Printf("server listening at %v", lis.Addr())
	api.RegisterUserFriendRPCServer(s, grpcuserfriend.NewGrpcUser(appCtx))
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
