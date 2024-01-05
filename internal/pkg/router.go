package webAPIUsers

import (
	"net/http"
)

func NewRouting(h *Handler) {
	http.HandleFunc("/users", h.loggingMD(h.GeneralHandler))
	http.HandleFunc("/users/all", h.loggingMD(h.GetAllUsers))
	http.ListenAndServe(":8080", nil)
}
