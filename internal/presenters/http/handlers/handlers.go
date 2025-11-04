package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
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
	log.Printf("[license] call License method")

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
		log.Printf("[license][error] err=%v", err)

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

func (h *Handler) GetLicense(w http.ResponseWriter, r *http.Request) {
	log.Printf("[license] call get License method")

	kidQuery := r.URL.Query().Get("kid")
	if kidQuery == "" {
		http.Error(w, "missing kid", http.StatusBadRequest)
		return
	}

	var req = servicepb.LicenseRequest{
		Kids: []string{kidQuery},
	}

	resp, err := h.svc.License(r.Context(), &req)
	if err != nil {
		log.Printf("[license][error] err=%v", err)

		returnJSON(w, http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})

		return
	}

	if resp.GetKeys() != nil {
		k := resp.GetKeys()[0].GetK()

		// Decode base64 URL-encoded key
		k64, err := base64.RawURLEncoding.DecodeString(k)
		if err != nil {
			log.Printf("[license][error] error when decode string: %v", err)
			http.Error(w, "invalid base64 key", http.StatusBadRequest)
			return
		}

		// Trả về raw binary (16 bytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(k64)
		if err != nil {
			log.Printf("[license][error] failed to write response: %v", err)
		}

		log.Printf("[debug][license] return raw key (len=%d): %x", len(k64), k64)
		return
	}

	w.WriteHeader(http.StatusNotFound)
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
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(data)
	_, _ = w.Write(resp)
}
