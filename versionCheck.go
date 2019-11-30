package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type VersionOL struct {
	Version string
}

func versionCheck() (v string, newest bool) {
	resp, err := http.Get("https://gitee.com/rocket049/secret-diary/raw/master/version.json")
	if err != nil {
		log.Println(err)
		newest = false
		return
	}
	defer resp.Body.Close()

	buf := bytes.NewBufferString("")
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Println(err)
		newest = false
		return
	}

	var ver VersionOL
	err = json.Unmarshal(buf.Bytes(), &ver)
	if err != nil {
		log.Println(err)
		newest = false
		return
	}

	if version == ver.Version {
		v = version
		newest = true
		return
	}

	v = ver.Version
	newest = false
	return
}
