package user

import (
	. "api_server/v7/handler"
	"api_server/v7/model"
	"api_server/v7/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	//log.Info("Delete function called.",
	//	lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("想要删除的ID为%+v\n", userId)
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
