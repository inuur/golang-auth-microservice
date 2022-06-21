package user

import (
	"authService/internal/domain/errors"
	"authService/internal/domain/handler"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.POST("/users", h.CreateUser)
	router.GET("/users/:id", h.GetUser)
}

func NewUserHandler() handler.Handler {
	return &Handler{
		service: NewUserService(),
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userDto := CreateUserDTO{}
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: err.Error(),
		})
		return
	}

	if err := userDto.OK(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(userDto)
	if err != nil {
		if strings.Contains(err.Error(), "E11000") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&errors.APIError{
				Error: "User is already registered!",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("id")

	user, err := h.service.FindUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
