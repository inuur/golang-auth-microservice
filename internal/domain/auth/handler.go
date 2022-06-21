package auth

import (
	"authService/internal/domain/errors"
	"authService/internal/domain/handler"
	"authService/internal/domain/token"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	service      *Service
	tokenService *token.Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.GET("/token/:uuid", h.Token)
	router.POST("/refresh", h.Refresh)
}

func NewAuthHandler() handler.Handler {
	return &Handler{
		service:      NewAuthService(),
		tokenService: token.NewTokenService(),
	}
}

func (h *Handler) Token(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	tokenDto, err := h.service.Authorize(params.ByName("uuid"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: "user not found",
		})
		return
	}
	json.NewEncoder(w).Encode(tokenDto)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tokenDto := &token.JWTTokenDTO{}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(tokenDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: err.Error(),
		})
		return
	}

	if err := tokenDto.OK(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: err.Error(),
		})
		return
	}

	newTokenDTO, err := h.tokenService.RefreshToken(tokenDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.APIError{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(newTokenDTO)
}
