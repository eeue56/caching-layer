package main

import "testing"

func TestLocalPath(t *testing.T) {
    root := "/"
    user := "elm-lang"
    packageName := "core"
    version := "3.0.0"

    finalLocation := "/elm-lang/core/3.0.0/code.zip"


    path := localPath(root, user, packageName, version)
    if path != finalLocation {
        t.Errorf("Path should match %s, got %s instead", finalLocation, path)
    }
}
