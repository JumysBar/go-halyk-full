package delivery

import (
	"example_project/api/usecase"
	"fmt"
	"net/http"
	"strconv"
)

type GetInfoHandler struct {
	uc usecase.ProcessUserUsecase
}

func (h *GetInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("user").(string)

	user, err := h.uc.ProcessUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Process user error: " + err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("User info: %+v", user)))
}

func NewGetInfoHandler(mux *http.ServeMux, uc usecase.ProcessUserUsecase) {
	handler := &GetInfoHandler{
		uc: uc,
	}
	mux.Handle("/", handler)
}

type UpdateAgeHandler struct {
	uc usecase.UpdateUserAgeUsecase
}

func (h *UpdateAgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("user").(string)

	newAgeStr := r.FormValue("age")
	if newAgeStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'age' value not found"))
		return
	}
	newAge, err := strconv.ParseInt(newAgeStr, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'age' value not valid"))
		return
	}
	h.uc.UpdateUserAge(username, newAge)
}

func NewUpdateAgeHandler(mux *http.ServeMux, uc usecase.UpdateUserAgeUsecase) {
	handler := &UpdateAgeHandler{
		uc: uc,
	}
	mux.Handle("/update", handler)
}
