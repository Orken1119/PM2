package admin

import (
	"net/http"

	"github.com/Orken1119/PM2/internal/controllers/tokenutil"
	models "github.com/Orken1119/PM2/internal/models"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	UserRepository models.UserRepository
}

// @Tags admin-controller
// @Summary Add user
//	    @Accept		json
//	    @Produce	json
//	    @Param request body models.AddUser true "query params"
//	    @Success	200		{object}	models.SuccessResponse
//		@Failure	default	{object}	models.ErrorResponse
//	    @Router		/admin/add-user [post]
func (ad AdminController) AddUser(c *gin.Context) {
	var request models.AddUser

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signin",
				},
			},
		})
		return
	}

	_, err = ad.UserRepository.AddUser(c, request.Name, request.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_USERS",
					Message: "Couldn't create user",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	user, err := ad.UserRepository.GetUserByEmail(c, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email doesn't found",
				},
			},
		})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, `access-secret-key`, 50) //
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})
}
