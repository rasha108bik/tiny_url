package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rasha108bik/tiny_url/config"
	storage "github.com/rasha108bik/tiny_url/internal/storage/db"
	storagefile "github.com/rasha108bik/tiny_url/internal/storage/file"
)

func TestHandlers(t *testing.T) {
	db := storage.NewStorage()
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", cfg)

	fileName := cfg.FileStoragePath
	strgFile, err := storagefile.NewFileStorage(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer strgFile.Close()

	handler := NewHandler(&cfg, db, strgFile)

	var shortenURL string
	var originalURL string

	t.Run("save", func(t *testing.T) {
		originalURL = "http://jqymby.biz/wruxoh/eii7bbkvbz4oj"

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(originalURL))
		w := httptest.NewRecorder()
		h := http.HandlerFunc(handler.CreateShortLink)
		h(w, request)
		result := w.Result()

		assert.Equal(t, http.StatusCreated, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

		userResult, err := io.ReadAll(result.Body)
		require.NoError(t, err)
		err = result.Body.Close()
		require.NoError(t, err)

		shortenURL = string(userResult)

		// проверяем URL на валидность
		_, urlParseErr := url.Parse(shortenURL)
		assert.NoErrorf(t, urlParseErr, "cannot parsee URL: %s ", shortenURL, err)
	})

	t.Run("get", func(t *testing.T) {
		uri, err := url.Parse(shortenURL)
		require.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/"+uri.Path[1:], nil)
		w := httptest.NewRecorder()
		h := http.HandlerFunc(handler.GetOriginalURL)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", uri.Path[1:])
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

		h(w, request)
		result := w.Result()
		err = result.Body.Close()
		require.NoError(t, err)

		assert.Equal(t, http.StatusTemporaryRedirect, result.StatusCode)
		assert.Equalf(t, originalURL, result.Header.Get("Location"),
			"Несоответствие URL полученного в заголовке Location ожидаемому",
		)
	})

	t.Run("save shorten", func(t *testing.T) {
		reqBody, err := json.Marshal(map[string]string{
			"url": "http://fsdkfkldshfjs.ru/test",
		})
		require.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		h := http.HandlerFunc(handler.CreateShorten)
		h(w, request)
		result := w.Result()

		err = result.Body.Close()
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, result.StatusCode)
		assert.Equal(t, "application/json", result.Header.Get("Content-Type"))

		m := RespReqCreateShorten{}
		err = json.NewDecoder(result.Body).Decode(&m)
		require.NoError(t, err)

		// проверяем URL на валидность
		_, urlParseErr := url.Parse(m.Result)
		assert.NoErrorf(t, urlParseErr, "cannot parsee URL: %s ", m.Result, err)
	})
}
