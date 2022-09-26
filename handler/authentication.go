package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/dgrijalva/jwt-go.v3"

	"Saham-BE/authentication"
	"Saham-BE/config"
	"Saham-BE/constant"
)

var secretCode = []byte("AYANK EPI")

const tokenExpired = time.Hour * 2

type authenticationHandler struct {
	authenticationService authentication.Service
}

func NewHandler(authenticationService authentication.Service) *authenticationHandler {
	return &authenticationHandler{authenticationService}
}

func ConvertGetResponse(user authentication.User) authentication.AuthenticationResponse {
	return authentication.AuthenticationResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}
}

func GenerateToken(email string) (string, error) {
	// c := config.JwtClaim{
	// 	Email: email,
	// 	jwt.StandardClaims{

	// 	},
	// }
	c := config.JwtClaim{
		Email: email, // Custom field
		// Expired:
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpired).Unix(), // Expiration time
			Issuer:    "my-project",
			Audience:  "kami", // Issuer
		},
	}
	// Creates a signed object using the specified signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Use the specified secret signature and obtain the complete encoded string token
	return token.SignedString(secretCode)
}

func (h *authenticationHandler) LoginHandler(c *gin.Context) {
	var loginParams authentication.LoginParameter
	err := c.ShouldBindJSON(&loginParams)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.REQUIRED_LOGIN_PARAM,
			"message": "Parameter tidak sesuai",
			"data":    errorMessages,
		})

		return
		// log.Fatal(err)
	}

	user, err := h.authenticationService.Login(loginParams)
	if err != nil {
		// errorToPrint := c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.WRONG_COMBINATION_EMAIL_PASSWORD,
			"message": "Username atau password salah/tidak ditemukan",
		})
		return
	}
	jwt, err := GenerateToken(loginParams.Email)
	if err != nil {
		errorToPrint := c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.JWT_ERROR,
			"message": errorToPrint,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESS_LOGIN,
		"message": "Success login",
		"data":    user,

		"token": jwt,
	})

	return
}

func (h *authenticationHandler) RegisterHandler(c *gin.Context) {
	var registerParam authentication.RegisterParameter
	err := c.ShouldBindJSON(&registerParam)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.REQUIRED_REGISTER_PARAM,
			"message": errorMessages,
		})

		return
	}

	// idUser, _ := updateParam.ID.Int64()
	emailExist, err := h.authenticationService.CheckingEmail(registerParam.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.CHECKING_EMAIL_FAILED,
			"message": "CHECKING EMAIL GAGAL",
		})
		return
	}

	if emailExist.Email == "" {
		user, err := h.authenticationService.Register(registerParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    constant.REGISTRATION_FAILED,
				"message": "REGISTER GAGAL",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    constant.SUCCESS_REGISTER,
			"message": "REGISTER BERHASIL",
			"data":    user,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.EXISTING_EMAIL,
			"message": "EMAIL SUDAH TERDAFTAR",
		})
		return
	}

}

func (h *authenticationHandler) ForgetPasswordHandler(c *gin.Context) {
	var loginParams authentication.ForgetPasswordParameter
	err := c.ShouldBindJSON(&loginParams)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "SH4001",
			"message": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SH2001",
		"message": "Success login",
		"email":   loginParams.Email,
	})
	// }

}

func (h *authenticationHandler) ResetPasswordHandler(c *gin.Context) {
	var loginParams authentication.ResetPasswordParameter
	err := c.ShouldBindJSON(&loginParams)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "SH4001",
			"message": errorMessages,
			"data":    errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     "SH2001",
		"message":  "Success login",
		"password": loginParams.Password,
	})
	// }

}

func (h *authenticationHandler) FindAll(c *gin.Context) {
	list_user, err := h.authenticationService.FindAll()

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_LIST_USER_FAILED,
			"message": "Gagal mendapatkan list user",
			"data":    errorMessages,
		})
		return
	}

	var authsResponse []authentication.AuthenticationResponse

	for _, res := range list_user {
		authResponse := ConvertGetResponse(res)

		authsResponse = append(authsResponse, authResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESS_GET_LIST_USER,
		"message": "Success mengambil data user",
		"data":    authsResponse,
	})
}

func (h *authenticationHandler) FindById(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	selected_user, err := h.authenticationService.FindById(id)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_USER_FAILED,
			"message": "Gagal mendapatkan list user",
			"data":    errorMessages,
		})

		return
	}

	if selected_user.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_USER_NOT_FOUND,
			"message": "User tidak ditemukan",
		})

		return
	}

	authResponse := ConvertGetResponse(selected_user)

	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESS_GET_USER,
		"message": "Success mengambil data user",
		"data":    authResponse,
	})
}

func (h *authenticationHandler) UpdateHandler(c *gin.Context) {
	var updateParam authentication.UpdateUserParamater
	err := c.ShouldBindJSON(&updateParam)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.REQUIRED_UPDATE_USER_PARAM,
			"message": errorMessages,
		})

		return
	}

	idUser, _ := updateParam.ID.Int64()
	selectedUser, err := h.authenticationService.FindById(int(idUser))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_USER_FAILED,
			"message": "MENGAMBIL DATA USER GAGAL",
			"data":    err,
		})
		return
	}

	if updateParam.Email == selectedUser.Email {
		emailExist, err := h.authenticationService.CheckingEmail(updateParam.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    constant.CHECKING_EMAIL_FAILED,
				"message": "CHECKING EMAIL GAGAL",
			})
			return
		} else {
			if emailExist.Email == updateParam.Email {
				user, err := h.authenticationService.UpdateUser(updateParam)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    constant.UPDATE_USER_FAILED,
						"message": "UPDATE USER GAGAL",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"code":    constant.SUCCESS_UPDATE_USER,
					"message": "UPDATE USER BERHASIL",
					"data":    user,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    constant.EXISTING_EMAIL,
					"message": "EMAIL SUDAH DIGUNAKAN",
				})
			}
		}

	} else {
		user, err := h.authenticationService.UpdateUser(updateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    constant.UPDATE_USER_FAILED,
				"message": "UPDATE USER GAGAL",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    constant.SUCCESS_UPDATE_USER,
			"message": "UPDATE USER BERHASIL",
			"data":    user,
		})
	}
}

func (h *authenticationHandler) DeleteHandler(c *gin.Context) {
	var deleteUserParam authentication.DeleteUserParameter
	err := c.ShouldBindJSON(&deleteUserParam)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.ID_REQUIRED_DELETE_USER,
			"message": errorMessages,
		})

		return
	}

	idUser, _ := deleteUserParam.ID.Int64()
	selectedUser, err := h.authenticationService.FindById(int(idUser))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_USER_FAILED,
			"message": "MENGAMBIL DATA USER GAGAL",
			"data":    err,
		})
		return
	} else if selectedUser.ID != 0 {
		user, err := h.authenticationService.DeleteUser(selectedUser.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    constant.DELETE_FAILED,
				"message": "DELETE USER GAGAL",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    constant.SUCCESS_DELETE_USER,
			"message": "BERHASIL MENGHAPUS USER",
			"data":    user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_USER_NOT_FOUND,
			"message": "DATA USER TIDAK DITEMUKAN",
			"data":    err,
		})
		return
	}
}
