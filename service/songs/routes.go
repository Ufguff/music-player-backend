package songs

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ufguff/config"
	"github.com/ufguff/types"
	"github.com/ufguff/utils"
)

type Handler struct {
	store types.ISong
}

func NewHandler(store types.ISong) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) { // это мы закрепили методы
	router.HandleFunc("/", h.handleInfoSongs).Methods(http.MethodGet)
	router.HandleFunc("/songs/{id}", h.handleSong).Methods(http.MethodGet)
	router.HandleFunc("/image/{id}", h.handleImage).Methods(http.MethodGet)

}

func (h *Handler) handleInfoSongs(w http.ResponseWriter, r *http.Request) {

	payload, err := h.store.GetInfoSongs()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, payload)
}

func (h *Handler) handleSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	pathSong, err := h.store.GetPathSong(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	log.Println(pathSong)
	newPathSong := config.Envs.StaticPathAudio

	errCopy := utils.CopyFile(pathSong, newPathSong)
	if errCopy != nil {
		utils.WriteError(w, http.StatusInternalServerError, errCopy)
	}
	/*
		w.Header().Del("Content-Type")
		w.Header().Set("Content-Type", "audio/mpeg3")
		//	w.Header().Set("Content-Type", "text/html
		r.Header.Del("Content-Type")

		//	w.Header().Set("Content-Type", "text/html")
	*/
	w.Header().Set("Content-Type", "audio/mpeg3")
	// r.Header.Set("Content-Type", "audio/mpeg3")
	http.ServeFile(w, r, newPathSong)

	errDel := utils.DeleteFile(newPathSong)
	if errDel != nil {
		utils.WriteError(w, http.StatusInternalServerError, errDel)
	}
	//utils.WriteJSON(w, http.StatusOK, "ok")
}

func (h *Handler) handleImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// проблема в беке, надо channels применить
	log.Println(id)

	pathSong, err := h.store.GetPathImage(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	log.Println(pathSong)
	newPathSong := config.Envs.StaticPathImage

	errCopy := utils.CopyFile(pathSong, newPathSong)
	if errCopy != nil {
		utils.WriteError(w, http.StatusInternalServerError, errCopy)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, newPathSong)

	errDel := utils.DeleteFile(newPathSong)
	if errDel != nil {
		utils.WriteError(w, http.StatusInternalServerError, errDel)
	}
}
