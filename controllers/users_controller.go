package controllers

import (
	"github.com/fausto4911/GO-USER-API/domain/users"
	"github.com/fausto4911/GO-USER-API/services"
	"github.com/fausto4911/GO-USER-API/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
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
		r := errors.GetBadRequest("Ivalid Json")
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
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
