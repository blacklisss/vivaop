package routergin

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"vivaop/internal/entities/userentity"
	"vivaop/internal/usecases/app/repos/userrepo"
	"vivaop/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	Fname     string `json:"fname" binding:"required"`
	Mname     string `json:"mname"`
	Lname     string `json:"lname"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Birthdate string `json:"birthdate"`
	CountryID int32  `json:"country_id" binding:"required"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	Fname     string    `json:"fname"`
	Mname     string    `json:"mname"`
	Lname     string    `json:"lname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CountryID int32     `json:"country_id"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user *userentity.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Fname:     user.FName,
		Mname:     user.MName,
		Lname:     user.LName,
		Email:     user.Email,
		Phone:     user.Phone,
		CountryID: user.CountryID,
		CreatedAt: user.CreatedAt,
	}
}

func (router *RouterGin) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	birthdate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := router.hs.CreateUser(ctx, &userrepo.CreateUserParams{
		ID:        uid,
		Fname:     sql.NullString{String: req.Fname, Valid: true},
		Mname:     sql.NullString{String: req.Mname, Valid: true},
		Lname:     sql.NullString{String: req.Lname, Valid: true},
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPassword,
		Birthdate: sql.NullTime{Time: birthdate, Valid: true},
		CountryID: sql.NullInt32{Int32: req.CountryID, Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println(user)
	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}
