package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"social-network/internal/models"

	"github.com/gofrs/uuid/v5"
)

type contextKey string

const UserIDKey contextKey = "user"

const UserCookie contextKey = "cookie"

func WriteJson(w http.ResponseWriter, statuscode int, Data any) error {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Data)
	if err != nil {
		return err
	}
	return nil
}

func ParseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func SetSessionCookie(w http.ResponseWriter, uuid string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    uuid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   31536000,
	})
}

func DeleteSessionCookie(w http.ResponseWriter, uuid string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  uuid,
		Path:   "/",
		MaxAge: -1,
	})
}

func GetUserFromContext(ctx context.Context) (user models.User, uuidCookie string, ok bool) {
	user, ok = ctx.Value(UserIDKey).(models.User)
	if !ok {
		return
	}
	uuidCookie, ok = ctx.Value(UserCookie).(string)
	return
}

func StoreThePic(UploadDir string, file multipart.File, handler *multipart.FileHeader) (string, error) {
	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		os.Mkdir(UploadDir, os.ModePerm)
	}

	randomstr := GenerateUuid()
	fmt.Println(randomstr + handler.Filename)

	extensions := []string{".png", ".jpg", ".jpeg", ".gif"}
	extIndex := slices.IndexFunc(extensions, func(ext string) bool {
		return strings.HasSuffix(handler.Filename, ext)
	})
	if extIndex == -1 {
		return "", fmt.Errorf("err")
	}

	filePath := filepath.Join(UploadDir, randomstr+extensions[extIndex])
	dst, err := os.Create(filePath)
	if err != nil {
		return "", errors.New("could not save file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", errors.New("failed to save file")
	}

	return randomstr + extensions[extIndex], nil
}

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
