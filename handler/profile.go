package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Saham-BE/constant"
	"Saham-BE/profile"
)

type profileHandler struct {
	profileService profile.Service
}

func ProfileHandler(profileService profile.Service) *profileHandler {
	return &profileHandler{profileService}
}

func (h *profileHandler) GetBankList(c *gin.Context) {
	bank, err := h.profileService.GetBankList()

	if err != nil {
		// errorMessages := []string{}
		// for _, e := range err.() {
		// 	errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		// 	errorMessages = append(errorMessages, errorMessage)
		// }
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    constant.GET_BANK_LIST_FAILED,
			"message": "MENGAMBIL DATA USER GAGAL",
			"data":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESS_UPDATE_USER,
		"message": "UPDATE USER BERHASIL",
		"data":    bank,
	})
}
