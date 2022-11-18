package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type LogContent struct {
	Datetime       string `json:"datetime"`
	Timestamp      string `json:"timestamp"`
	NginxRequestId string `json:"nginx_request_id"`
	UniqueRemark   string `json:"unique_remark"`
	CnRemark       string `json:"cn_remark"`
	UserId         string `json:"user_id"`
	Project        string `json:"project"`
	Types          string `json:"type"`
	Path           string `json:"path"`
	Module         string `json:"module"`
	Host           string `json:"host"`
	Url            string `json:"url"`
	Level          string `json:"level"`
	Context        string `json:"context"`
	Backtrace      string `json:"backtrace"`
	PostData       string `json:"postData"`
	GetData        string `json:"getData"`
}

func SyncLog(dir string) {
	for _, entry := range dirents(dir) {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			SyncLog(path)
		} else {
			readDir(path)
		}
	}
}

func readDir(path string) {
	fmt.Println("path:", path)
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	fmt.Println("-------------")
	var lc []LogContent
	err = json.Unmarshal(jsonData, &lc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", jsonData)

	log.Fatal(lc)
	for _, lv := range lc {
		fmt.Println(lv.Datetime)
		log.Fatal()
	}
}

// write db

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("read dir:%v\n", err)
		return nil
	}
	return entries
}
