package frontend_app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// getServerPort returns the server port as an integer
func getServerPort() int {
	if os.Getenv("PORT")!= "" {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err!= nil {
			log.Fatal(err)
		}
		return p
	}
	return 3000
}

// getServerHost returns the server host as a string
func getServerHost() string {
	if os.Getenv("HOST")!= "" {
		return os.Getenv("HOST")
	}
	if runtime.GOOS == "windows" {
		return "localhost"
	}
	return "127.0.0.1"
}

// getServerAddress returns the server address as a string
func getServerAddress() string {
	return fmt.Sprintf("%s:%d", getServerHost(), getServerPort())
}

// getAbsPath returns the absolute path to the given path
func getAbsPath(path string) string {
	return filepath.Abs(path)
}

// getRelPath returns the relative path from the given path to the current directory
func getRelPath(path string) string {
	return filepath.Rel(getAbsPath(""), path)
}

// getDirName returns the directory name of the given path
func getDirName(path string) string {
	return filepath.Dir(path)
}

// getFileName returns the file name of the given path
func getFileName(path string) string {
	return filepath.Base(path)
}

// getExt returns the file extension of the given path
func getExt(path string) string {
	return filepath.Ext(path)
}

// getQueryValue returns the query value for the given key
func getQueryValue(r *http.Request, key string) string {
	q := r.URL.Query()
	if val, ok := q[key]; ok {
		return val[0]
	}
	return ""
}

// getQueryInt returns the query value for the given key as an integer
func getQueryInt(r *http.Request, key string) int {
	q := r.URL.Query()
	if val, ok := q[key]; ok {
		p, err := strconv.Atoi(val[0])
		if err!= nil {
			return 0
		}
		return p
	}
	return 0
}

// getQueryBool returns the query value for the given key as a boolean
func getQueryBool(r *http.Request, key string) bool {
	q := r.URL.Query()
	if val, ok := q[key]; ok {
		switch strings.ToLower(val[0]) {
		case "true":
			return true
		case "false":
			return false
		}
	}
	return false
}

// getQueryFloat returns the query value for the given key as a float64
func getQueryFloat(r *http.Request, key string) float64 {
	q := r.URL.Query()
	if val, ok := q[key]; ok {
		p, err := strconv.ParseFloat(val[0], 64)
		if err!= nil {
			return 0
		}
		return p
	}
	return 0
}