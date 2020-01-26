package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/rs/cors"
)

func MapFileHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	vars := mux.Vars(r)
	fileName := vars["mapfile"]

	filePath := path.Join(
		"tilesets/",
		fileName,
	)

	contentType := strings.ToLower(r.Header.Get("Accept"))

	if strings.Contains(contentType, "image/png") && strings.Contains(reqPath, ".png") {
		tileMap, err := tiled.LoadFromFile(filePath)

		if err != nil {
			sendError(err, w)
			return
		}

		renderer, err := render.NewRenderer(tileMap)

		if err != nil {
			sendError(err, w)
			return
		}

		err = renderer.RenderVisibleLayers()

		if err != nil {
			sendError(err, w)
			return
		}

		buff := bytes.NewBuffer(*new([]byte))

		err = renderer.SaveAsPng(buff)

		if err != nil {
			sendError(err, w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "image/png")
		w.Write(buff.Bytes())
		return
	}

	fs.ServeHTTP(w, r)
}

func sendError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%e", err)))
}

var fs http.Handler

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tilesets/{mapfile}", MapFileHandler)

	wd, _ := os.Getwd()

	fs = http.FileServer(http.Dir(wd))

	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		cors.Default().Handler(r),
	)

	if err != nil {
		panic(err)
	}
}
