package controllers

import (
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	var _ = []string{
		"john_kramnik",
		"magnus_carlsen",
		"naufal_ahza",
		"nakamura_hikaru",
		"hans_nieman",
	}

	helpers.Response(w, http.StatusOK, "My Profile", userResponse)
}
