package handlers

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/hashfyre/sample-go-app/app/context"
	"github.com/hashfyre/sample-go-app/app/controllers"
	"github.com/hashfyre/sample-go-app/app/models"
	"github.com/hashfyre/sample-go-app/app/types"
)

// RegisterNoAuthUserRoutes - unauthenticated routes
func RegisterNoAuthUserRoutes(router *gin.RouterGroup) {
	router.POST("/signup/", usersRegistration)
}

// RegisterUsersRoutes - registers gin user routes that neeed auth
func RegisterUsersRoutes(router *gin.RouterGroup) {
	router.GET("/", getUserDetails)
	router.PATCH("/", updateUser)
	router.DELETE("/", deleteUser)
	router.PATCH("/password/", updatePassword)
}

// UsersRegistration goDoc
// @Summary user registration
// @Description registers user and returns a access token along with response
// @Accept json
// @Produce json
// @Param value body types.RegisterRequestDto true "user registration"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {string} string "Could not bind request data, may contain unknown or missing fields"
// @Failure 422 {string} string "user not created"
// @Router /api/users/signup/ [post]
// UsersRegistration does user registration
func usersRegistration(c *gin.Context) {
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span.SetBaggageItem("os", runtime.GOOS)
		span.SetBaggageItem("arch", runtime.GOARCH)
	}

	var request types.RegisterRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ResponseError{
			Code:    "ERR_BAD_REQUEST",
			Message: "Could not bind request data, may contain unknown or missing fields",
		})
		return
	}

	_, err := controllers.UsersRegistration(request.FirstName, request.LastName, request.Email, request.Password)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ResponseError{
			Code:    "ERR_INTERNAL_SERVER_ERROR",
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// getUserDetails godoc
// @Summary fetch all users from the database
// @Description get all users from the database
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Success 200 {object} types.GetUserResponseDto
// @Failure 400 {string} string "User info not present in context"
// @Failure 404 {string} string "User not found"
// @Router /api/v1/users/ [get]
// getUserDetails returns current user details
func getUserDetails(c *gin.Context) {
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span.SetBaggageItem("os", runtime.GOOS)
		span.SetBaggageItem("arch", runtime.GOARCH)
	}

	userID, err := context.GetCtxUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errUserInfoContext)
		return
	}

	user, err := controllers.GetUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ResponseError{
			Code:    "ERR_RECORD_NOT_FOUND",
			Message: "user not found",
		})
		return
	}

	data := types.GetUserResponseDto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	c.JSON(http.StatusOK, data)
}

// updateUser godoc
// @Summary updates user based on details provided
// @Description update a cluster by changing any fields -> firstname | lastname
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param value body types.BasicProfileDto true "update cluster"
// @Success 200 {string} string "uUser details updated successfully"
// @Failure 400 {string} string "User info not present in context"
// @Failure 422 {string} string "Unable to update user"
// @Failure 404 {string} string "User not found"
// @Router /api/v1/users/ [patch]
// updateUser update allowed fields of a user/profile
func updateUser(c *gin.Context) {
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span.SetBaggageItem("os", runtime.GOOS)
		span.SetBaggageItem("arch", runtime.GOARCH)
	}

	var payloadJSON types.BasicProfileDto
	if err := c.ShouldBindJSON(&payloadJSON); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ResponseError{
			Code:    "ERR_BAD_REQUEST",
			Message: "Could not bind request data, may contain unknown or missing fields",
		})
		return
	}

	userID, err := context.GetCtxUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errUserInfoContext)
		return
	}

	var user models.User
	if payloadJSON.FirstName != "" {
		user.FirstName = payloadJSON.FirstName
	}

	if payloadJSON.LastName != "" {
		user.LastName = payloadJSON.LastName
	}

	updatedUser, err := controllers.UpdateUser(userID, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ResponseError{
			Code:    "ERR_INTERNAL_SERVER_ERROR",
			Message: "Failed to update user",
		})
		return
	}

	if updatedUser == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ResponseError{
			Code:    "ERR_RECORD_NOT_FOUND",
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User details updated successfully",
	})
}

// deleteUser godoc
// @Summary delete a user based on userID provided
// @Param Authorization header string true "Authorization Header"
// @Produce json
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "User info not present in context"
// @Failure 422 {string} string "Unable to delete user"
// @Failure 404 {string} string "User Not Found"
// @Router /api/v1/users/ [delete]
// deleteUser disables current user
func deleteUser(c *gin.Context) {
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span.SetBaggageItem("os", runtime.GOOS)
		span.SetBaggageItem("arch", runtime.GOARCH)
	}

	userID, err := context.GetCtxUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errUserInfoContext)
		return
	}

	numOfRows, err := controllers.DeleteUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ResponseError{
			Code:    "ERR_INTERNAL_SERVER_ERROR",
			Message: "Failed to delete user",
		})
		return
	}

	if numOfRows == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ResponseError{
			Code:    "ERR_RECORD_NOT_FOUND",
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// UpdatePassword godoc
// @Summary updates password with new password provided
// @Description update a password by providing  -> current-password | new-password
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param value body types.ChangePasswordDto true "update password"
// @Success 200 {string} string "password updated sucessfully"
// @Failure 400 {string} string "User info not present in context"
// @Failure 422 {string} string "unable to change password"
// @Failure 404 {string} string "invalid user name or password"
// @Router /api/v1/users/password/ [patch]
// UpdatePassword update user password
func updatePassword(c *gin.Context) {
	var payloadJSON types.ChangePasswordDto
	if err := c.ShouldBindJSON(&payloadJSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ResponseError{
			Code:    "ERR_BAD_REQUEST",
			Message: "Could not bind request data, may contain unknown or missing fields",
		})
		return
	}

	userID, err := context.GetCtxUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errUserInfoContext)
		return
	}

	numOfRows, err := controllers.UpdatePassword(userID, payloadJSON.CurrentPassword, payloadJSON.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ResponseError{
			Code:    "ERR_INTERNAL_SERVER_ERROR",
			Message: "Failed to change password",
		})
		return
	}

	if numOfRows == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ResponseError{
			Code:    "ERR_RECORD_NOT_FOUND",
			Message: "Invalid user name or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
