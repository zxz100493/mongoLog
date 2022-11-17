package service

import "os"

func SyncLog() {
	// read file
	pwd := "/Logs"
	os.ReadDir(pwd)
	// write db

}
