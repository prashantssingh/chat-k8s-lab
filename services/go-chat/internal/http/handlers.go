package httpapi

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go-chat/internal/chat"
)

type Handler struct {
	PythonChatURL string
	Client        *http.Client
	PodName       string
}

func NewHandler() *Handler {
	return &Handler{
		PythonChatURL: getEnv("PYTHON_CHAT_URL", "http://python-chat/reply"),
		Client: &http.Client{
			Timeout: 2 * time.Second,
		},
		PodName: getEnv("POD_NAME", "local"),
	}
}

func (h *Handler) Healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func (h *Handler) Readyz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ready"}`))
}

func (h *Handler) Send(w http.ResponseWriter, _ *http.Request) {
	msg := chat.PickMessage()

	payload := map[string]string{
		"from":     "go-chat",
		"message":  msg,
		"trace_id": time.Now().Format(time.RFC3339Nano),
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", h.PythonChatURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "request creation failed", 500)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		http.Error(w, "python-chat unreachable", 502)
		return
	}
	defer resp.Body.Close()

	log.Printf("pod=%s sent=%s status=%d", h.PodName, msg, resp.StatusCode)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sent"}`))
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
