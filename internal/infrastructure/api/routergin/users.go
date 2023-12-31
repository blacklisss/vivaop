package routergin

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"vivaop/internal/entities/userentity"
	"vivaop/internal/infrastructure/notificator"
	"vivaop/internal/usecases/app/repos/userrepo"
	"vivaop/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type createUserRequest struct {
	CountryID string `json:"country_id" binding:"required"`
	Fname     string `json:"fname" binding:"required"`
	Mname     string `json:"mname"`
	Lname     string `json:"lname"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Birthdate string `json:"birthdate"`
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

	birthdate, err := time.Parse(time.RFC3339, req.Birthdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c_id, err := strconv.Atoi(req.CountryID)
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
		CountryID: sql.NullInt32{Int32: int32(c_id), Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	emailToken := util.GetMD5Hash(user.ID.String())
	for {
		_, err = router.hs.CheckEmailToken(ctx, emailToken)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				break
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	expireAt := util.CreateExpireDate(router.config.EmailTokenExpire)
	emailVerification, err := router.hs.CreateEmailVerification(ctx, &userrepo.CreateEmailVerificationParams{
		UserID:    user.ID,
		Token:     emailToken,
		ExpiredAt: expireAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = router.notificatorer.Notificator.Send(ctx, &notificator.EmailData{
		Token:   emailVerification.Token,
		To:      user.Email,
		Subject: "Email Verification Link",
		Name:    fmt.Sprintf("%s %s %s", user.FName, user.MName, user.LName),
		Host:    router.config.HTTPServerAddress,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}
