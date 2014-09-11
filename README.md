# WebTranslateIt GoLang Client

[![Build Status](https://secure.travis-ci.org/fromYukki/webtranslateit_go_client.png?branch=master)](http://travis-ci.org/fromYukki/webtranslateit_go_client)

This package is under active development. You can get more information about WebTranslateIt API using this link [https://webtranslateit.com/en/docs/api/](https://webtranslateit.com/en/docs/api/).

## Installation

To install this package, please, use default **go get** tool

    go get github.com/fromYukki/webtranslateit_go_client

## Getting started

Authentication is made by so-called API tokens, and of course you need specify this one of the token to the `WebTranslateIt` structure.

```go
import wti_client "github.com/fromYukki/webtranslateit_go_client"

func main() {
    wti := wti_client.NewWebTranslateIt("YOUR_TOKEN")
    project, err := wti.GetProject()
    if err != nil {
        panic(err)
    }
}
```

If you need to change API URL address or Token you can do it using next methods: `SetApiUrl` and `SetToken`.

### Project API

Project API section has only one method *Show Project*. You can read about it [here](https://webtranslateit.com/en/docs/api/project/).

#### Show Project

As shown in the example above, you can get the project and use it data as you wish. For more information, please, take a look on `Project` structure.

```go
import (
    "fmt"
    wti_client "github.com/fromYukki/webtranslateit_go_client"
)

func main() {
    wti := wti_client.NewWebTranslateIt("YOUR_TOKEN")
    project, err := wti.GetProject()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Project name: %q with %d files", project.Name, len(project.ProjectFiles))
}
```
    
### File API

Only one method is implemented in the File API section. The rest you can find [here](https://webtranslateit.com/en/docs/api/file/).

#### Zip File

The easiest method to get all the translation files - is to download them in Zip archive.

```go
import (
    "fmt"
    wti_client "github.com/fromYukki/webtranslateit_go_client"
)

func main() {
    var (
        err     error
        project wti_client.Project
        zipFile wti_client.ProjectZipFile
        data	map[string][]byte
    )
    
    wti := wti_client.NewWebTranslateIt("YOUR_TOKEN")
    if project, err = wti.GetProject(); err != nil {
        panic(err)
    }
    
    if zipFile, err = project.ZipFile(); err != nil {
        panic(err)
    }
    
    if data, err = zipFile.Extract(); err != nil {
        panic(err)
    }
    
    for fileName, fileData := range data {
        fmt.Printf("Extracted file %q with %d bytes length", fileName, len(fileData))
    }
}
```
