package webAPIUsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sskrill/WebAPI-Proj/pkg/cache"
	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
type Handler struct {
	crud  CRUD
	cache *cache.Cache
}

func NewHandler(crud CRUD, cache *cache.Cache) *Handler {
	return &Handler{crud: crud, cache: cache}
}
func (h *Handler) GeneralHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUser(w, r)
	case http.MethodPut:
		h.UpdateUser(w, r)
	case http.MethodDelete:
		h.DeleteUser(w, r)
	case http.MethodPost:
		h.CreateUser(w, r)
	}
}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("cant conv to int (Get)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, ok := h.cache.Get(id)
	if ok {
		log.Println("from cache (Get)")
		_ = json.NewEncoder(w).Encode(user)
		return
	}
	user, err = h.crud.Get(id)

	if err != nil {
		log.Println("not found user(Get)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("cant conv to json (Get)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.cache.Set(id, user)
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("cant conv to int (Update)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("cant read body (Update)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("cant unmarshal body (Update)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.crud.Update(id, user)
	if err != nil {
		log.Println("not found user (update)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Updated User")
}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("cant conv to int (Delete)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.crud.Delete(id)
	if err != nil {
		log.Println("not found user (Delete)")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprint(w, "Deleted User")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("cant to read body (Create)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("cant cant unmarshal body (Create)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.crud.Insert(user)
	if err != nil {
		log.Println("not found user (Create)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Created User")
}
func (h *Handler) loggingMD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
		}).Info("Request Procesed")
		next(w, r)

	}
}
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.crud.GetAll()
	if len(users) == 0 {
		err := errors.New("not found users(GetAll)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("cant give json (GetAll) ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
