package API

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todorokvaja1/DataStructures"
)

func (a *Controller) Login(c *gin.Context) {

	var userLogin DataStructures.UserLogin
	err := c.BindJSON(&userLogin)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	user, err := a.c.Login(c.Request.Context(), userLogin)

	c.String(http.StatusOK, "Prijava uspešna: %s", user)
}

func (a *Controller) InsertUser(c *gin.Context) {

	var user DataStructures.User
	err := c.BindJSON(&user)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	err = a.c.InsertUser(c.Request.Context(), user)

	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	c.String(http.StatusOK, "registracija novega userja")
}

func (a *Controller) GetUserByName(c *gin.Context) {

	userByName := c.Param("username")

	user, err := a.c.GetUserByName(c.Request.Context(), userByName)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, user)
}
