package main

import "fmt"
import "os"
import "io"
import "strings"
import "os/user"
import "net/http"
import "github.com/gorilla/mux"

// /elm-lang/core/zipball/3.0.0/

func localPath(root string, user string, packageName string, version string) string {
    return fmt.Sprintf("%s/%s/%s/%s/code.zip", root, user, packageName, version)
}

func mainElmHost(url string) string {
    return "https://github.com" + url
}

func downloadFile(path string, url string) {
    fmt.Printf("Downloading file %s to %s", mainElmHost(url), path)
    os.MkdirAll(path[0:strings.LastIndex(path, "/")], 0777)

    out, err := os.Create(path)
    defer out.Close()

    if err != nil {
        fmt.Printf("Failed to create..\n %s\n", err.Error())
    }

    resp, _ := http.Get(mainElmHost(url))
    defer resp.Body.Close()
    n, err := io.Copy(out, resp.Body)
    fmt.Printf("copied %d bytes", n)

    if err != nil {
        fmt.Printf("something went wrong..\n %s\n", err.Error())
    }
}

func PackageGetHandler(w http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    user := vars["user"]
    packageName := vars["package"]
    version := vars["version"]

    root := setup()

    path := localPath(root, user, packageName, version)

    if _, err := os.Stat(path); os.IsNotExist(err) {
        downloadFile(path, request.URL.Path)
    }

    fmt.Printf("file requested at %s\n", path)
    http.ServeFile(w, request, path)
}

func setup() string {
    user, _ := user.Current()
    return user.HomeDir + "/cache"
}


func main() {
    fmt.Printf("Hello, world.\n")

    r := mux.NewRouter()
    r.HandleFunc("/{user}/{package}/zipball/{version}/", PackageGetHandler)
    http.ListenAndServe(":80", r)
}
