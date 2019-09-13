package main

import (
    "flag"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
  fmt.Println("File Upload Endpoint Hit")

  r.ParseMultipartForm(1 << 20)   // maximum upload size 1M
  file, handler, err := r.FormFile("myFile")

  if err != nil {
    fmt.Println("Error Retrieving the File")
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Printf("Uploaded File: %+v\n", handler.Filename)
  fmt.Printf("File Size: %+v\n", handler.Size)
//  fmt.Printf("MIME Header: %+v\n", handler.Handler)

  tmpFile, err := ioutil.TempFile("upload-dir", "upload-*.img")
  if err != nil {
    fmt.Println(err)
  }
  defer tmpFile.Close()

  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
  }

  tmpFile.Write(fileBytes)
  fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  port      := flag.String("p", "8080", "port to serve on")
  directory := flag.String("d", ".", "the directory of static file to host")
  flag.Parse()

  http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), http.FileServer(http.Dir(*directory))))

  //http.HandleFunc("/", handler)
  http.HandleFunc("/upload", uploadFile)
  fs := http.FileServer(http.Dir("./download-dir"))
  http.Handle("/files/", http.StripPrefix("/files", fs))

  log.Fatal(http.ListenAndServe(":" + *port, nil))
}
