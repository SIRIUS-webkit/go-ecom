package user

import (
	"fmt"
	"net/http"

	"github.com/SIRIUS-webkit/go-ecom/services/auth"
	"github.com/SIRIUS-webkit/go-ecom/types"
	"github.com/go-playground/validator/v10"

	"github.com/SIRIUS-webkit/go-ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){
	
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validator.Struct(payload); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check existing user
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exist", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	// if doesn't exist, create user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

    utils.WriteJSON(w, http.StatusCreated, nil)

}