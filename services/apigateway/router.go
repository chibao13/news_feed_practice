package main

import (
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/middleware"
	"github.com/chibao13/news_feed_practice/services/apigateway/authenmiddleware"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/usertransport/ginuser"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/userfriend/userfriendtransport/ginuserfriend"
	"github.com/gin-gonic/gin"
)

//API gateway
func setupRouter(r *gin.Engine, appCtx appctx.AppContext, authenStore authenmiddleware.AuthenStore) {
	r.Use(middleware.Recover(appCtx)) // global middleware

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	users := v1.Group("users")
	{
		//users.GET("/newsfeed", userMiddleware.RequiredAuth(appCtx, authenStore), ginpost2.ListNewsFeed(appCtx))
		//users.GET("/:user-id/posts", userMiddleware.RequiredAuth(appCtx, authenStore), ginpost2.CreatePost(appCtx))
		users.GET("/:user-id/friends", ginuserfriend.ListFriend(appCtx))
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.StaticFile("/demo/", "./demo.html")
}
