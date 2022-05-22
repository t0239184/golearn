package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/t0239184/golearn/internal/database"
	"github.com/t0239184/golearn/internal/model"
	"github.com/t0239184/golearn/internal/router/api/v1/request"
	"github.com/t0239184/golearn/internal/router/api/v1/response"
)

type UserHandler struct {
	DB *database.GormDatabase
}

/* Inject database implement */
func NewUserHandler(db *database.GormDatabase) *UserHandler {
	return &UserHandler{DB: db}
}

func (u *UserHandler) QueryAllUser(c *gin.Context) {
	logrus.Info("[QueryAllUser] Start.")
	users, err := u.DB.FindAllUser()
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserHandler) FindUserById(c *gin.Context) {
	logrus.Info("[FindUserById] Start.")
	/* Convert id string to id int64 */
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	/* Query user from database */
	user, err := u.DB.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	logrus.Info("[CreateUser] Start.")

	/* Convert request json to request object */
	request := &request.CreateUserRequest{}
	c.BindJSON(request)

	/* Valide */
	if request.Account == "" || request.Password == "" {
		c.JSON(http.StatusOK,  response.FailResponse(400, "Account or Password should not be blank."))
		return
	}

	/* Convert request object to entity */
	newUser := model.NewUser(request)

	/* Save to database */
	user, err := u.DB.CreateUser(newUser)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	/* Return success response */
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	logrus.Info("UpdateUser")

	/* Convert request json to request object */
	request := &request.UpdateUserRequest{}
	c.BindJSON(request)

	updateUser := model.UpdateUser(request)
	user, err := u.DB.UpdateUser(updateUser)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	logrus.Info("DeleteUser")
}
