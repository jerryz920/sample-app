package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	jhttp "github.com/jerryz920/utils/goutils/http"
	"github.com/sirupsen/logrus"
)

const (
	DefaultPort = 19850
	DefaultPage = "index.html"
)

func DocRoot(w http.ResponseWriter, r *http.Request) {
	findex, err := os.Open(DefaultPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(w, findex)
}

func main() {
	server := jhttp.NewAPIServer(DocRoot)
	if err := server.ListenAndServe(fmt.Sprintf(":%d", DefaultPort)); err != nil {
		logrus.Fatal("listen and serve: ", err)
	}
}
