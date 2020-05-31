package controllers

import (
	"github.com/fausto4911/GO-USER-API/domain/users"
	"github.com/fausto4911/GO-USER-API/services"
	"github.com/fausto4911/GO-USER-API/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	user := users.User{}
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: handle error
	//	fmt.Println(err)
	//	c.String(http.StatusBadRequest, "Error getting http body")
	//}
	//if err := json.Unmarshal(bytes, &user); err !=nil {
	//	//TODO: handle json parse error
	//	fmt.Println(err)
	//	c.String(http.StatusBadRequest, "Invalid Json")
	//}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		//TODO: handle json parse error
		r := errors.NewBadRequestError("Ivalid Json")
		c.JSON(r.Status,r)
		return
	}
   u, e := services.CreateUser(user)
   if e != nil {
   	//TODO: handle user creation error
   	c.JSON(e.Status, e)
   	return
   }
	c.JSON(http.StatusCreated, u)
   return
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
