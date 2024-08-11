package controllers

import (
  "encoding/json"
  "go-jwt/configs"
  "go-jwt/helpers"
  "go-jwt/models"
  "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
  var register models.Register

  err := json.NewDecoder(r.Body).Decode(&register)
  if err != nil {
    helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
    return
  }

  defer r.Body.Close()

  if register.Password != register.PasswordConfirm {
    helpers.Response(w, http.StatusBadRequest, "Passwords do not match", nil)
    return
  }

  passwordHash, err := helpers.HashPassword(register.Password)
  if err != nil {
    helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
    return
  }

  user := models.User{
    Name:     register.Name,
    Email:    register.Email,
    Password: passwordHash,
  }

  if configs.DB.Create(&user).Error != nil {
    helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
    return
  }

  helpers.Response(w, http.StatusCreated, "OK", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
  var login models.Login

  err := json.NewDecoder(r.Body).Decode(&login)
  if err != nil {
    helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
    return
  }

  defer r.Body.Close()

  var user models.User
  if err := configs.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
    helpers.Response(w, http.StatusNotFound, "Email is invalid", nil)
    return
  }

  if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
    helpers.Response(w, http.StatusUnauthorized, "Passwords do not match", nil)
    return
  }

  token, err := helpers.CreateToken(&user)
  if err != nil {
    helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
    return
  }

  helpers.Response(w, http.StatusOK, "Succesfully Logged In", token)
}
