package user

import (
	"api_server/v8/model"
	"api_server/v8/pkg/errno"
	"api_server/v8/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "api_server/v8/handler"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

//func Create(c *gin.Context) {
//	log.Info("User Create function called.",
//		lager.Data{"X-Request-Id": util.GetReqID(c)})
//
//	var r CreateRequest
//	if err := c.Bind(&r); err != nil {
//		SendResponse(c, errno.ErrBind, nil)
//		return
//	}
//
//	u := model.UserModel{
//		Username: r.Username,
//		Password: r.Password,
//	}
//	// Validate the data.
//	if err := u.Validate(); err != nil {
//		SendResponse(c, errno.ErrValidation, nil)
//		return
//	}
//
//	// Encrypt the user password.
//	if err := u.Encrypt(); err != nil {
//		SendResponse(c, errno.ErrEncrypt, nil)
//		return
//	}
//	// Insert the user to the database.
//	if err := u.Create(); err != nil {
//		SendResponse(c, errno.ErrDatabase, nil)
//		return
//	}
//
//	rsp := CreateResponse{
//		Username: r.Username,
//	}
//
//
//	// Show the user information.
//	SendResponse(c, nil, rsp)
//
//}
