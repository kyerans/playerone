package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	servicepb "github.com/kyerans/playerone/api/services/v1"
	"github.com/kyerans/playerone/internal/services"
)

func New(svc *services.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc *services.Service
}

func (h *Handler) License(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		returnJSON(w, http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	var req servicepb.LicenseRequest
	if err := json.Unmarshal(body, &req); err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	resp, err := h.svc.License(r.Context(), &req)
	if err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	returnJSON(w, http.StatusOK, resp)
}

func (h *Handler) LicenseRelease(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		returnJSON(w, http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	var req servicepb.LicenseReleaseRequest
	if err := json.Unmarshal(body, &req); err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	resp, err := h.svc.LicenseRelease(r.Context(), &req)
	if err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	returnJSON(w, http.StatusOK, resp)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		returnJSON(w, http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	var req servicepb.RegisterRequest
	if err := json.Unmarshal(body, &req); err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	resp, err := h.svc.Register(r.Context(), &req)
	if err != nil {
		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	returnJSON(w, http.StatusOK, resp)
}

func returnJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	resp, _ := json.Marshal(data)
	_, _ = w.Write(resp)
}
