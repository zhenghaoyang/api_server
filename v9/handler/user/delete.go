package user

import (
	. "api_server/v9/handler"
	"api_server/v9/model"
	"api_server/v9/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	//log.Info("Delete function called.",
	//	lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
