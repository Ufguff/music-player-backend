package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("body is nil")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// work with files

func CopyFile(pathSong, newPathSong string) error {
	file, err := os.OpenFile(pathSong, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := os.Create(newPathSong)
	if err != nil {
		return err
	}
	defer newFile.Close()

	bytesCop, err := io.Copy(newFile, file)

	if err != nil {
		return err
	}

	log.Printf("Bytes copied %d!", bytesCop)
	return nil
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
