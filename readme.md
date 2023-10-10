# GoUploader

[![Static Badge](https://img.shields.io/badge/Go-blue.svg)](https://go.dev/) [![Static Badge](https://img.shields.io/badge/v1.0.0-blue.svg)](https://go.dev/)


Go-uploader is a Go package for handling file uploads in HTTP requests. It provides functions to upload single or multiple files and save them to a specified directory.

## Usage

### Importing the Package

Import the package in your Go code:

```go
import "github.com/celpung/go-uploader"
```

### Single file upload
To handle the upload of a single file from an HTTP request, you can use the "Single" function:
```go
uploadedFile, err := gouploader.Single(request, directory, fieldName)
if err != nil {
    // Handle the error
}
if uploadedFile != nil {
    // The file was successfully uploaded
    fmt.Printf("Uploaded file: %s\n", uploadedFile.Filename)
}
```

### Multiple file upload
To handle the upload of multiple files from an HTTP request, you can use the "Multiple" function:
```go
uploadedFiles, err := gouploader.Multiple(request, directory, fieldName)
if err != nil {
    // Handle the error
}
if len(uploadedFiles) > 0 {
    // Files were successfully uploaded
    for _, file := range uploadedFiles {
        fmt.Printf("Uploaded file: %s\n", file.Filename)
    }
}
```

### Example
Here's a simple example of how to use go-uploader in a web server:
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/yourusername/gouploader"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Handle single file upload
    uploadedFile, err := gouploader.Single(r, "./uploads", "file")
    if err != nil {
        http.Error(w, "Upload failed", http.StatusInternalServerError)
        return
    }
    if uploadedFile != nil {
        fmt.Fprintf(w, "Uploaded file: %s\n", uploadedFile.Filename)
    }

    // Handle multiple file upload
    uploadedFiles, err := gouploader.Multiple(r, "./uploads", "files")
    if err != nil {
        http.Error(w, "Upload failed", http.StatusInternalServerError)
        return
    }
    if len(uploadedFiles) > 0 {
        for _, file := range uploadedFiles {
            fmt.Fprintf(w, "Uploaded file: %s\n", file.Filename)
        }
    }
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Contributing

Contributions are welcome! To contribute to cexam, follow these steps:

1. Fork this repository.
2. Create a new branch: git checkout -b new-feature
3. Make changes and commit: git commit -m 'Add new feature'
4. Push to your forked repository: git push origin new-feature
5. Create a pull request explaining your changes.
6. Thank you for contributing!

## License

This package is distributed under the [MIT License](https://opensource.org/license/mit/).




