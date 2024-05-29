package data

import (
	"fmt"
	"main/tick"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Init() error {

	tick.EveryMinute(RemoveDeprecatedFromCache)
	RemoveDeprecatedFromCache()
	return nil

}

func AddToCache(cid string, text string, t int64) {

	// Check params

	if len(cid) == 0 || len(text) == 0 {
		return
	}

	// Create cache folder if needed

	folderPath := "./cache/" + cid + "/"
	createFolder(folderPath)

	// Write text to cache file

	os.WriteFile(folderPath+toFilename(t)+".txt", []byte(text), 0644)

}

func GetFromCache(cid string) string {

	// Check params

	if len(cid) == 0 {
		return ""
	}

	// Get cache

	folderPath := "./cache/" + cid + "/"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return ""
	}

	// Find cached file in folder

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			pathfile := folderPath + filename

			// Check filename

			_, t, err := fromFilename(strings.TrimSuffix(filename, filepath.Ext(filename)))
			if err != nil {
				continue
			}

			// Get cached content

			content, _ := os.ReadFile(pathfile)

			// Renew filename

			os.Rename(pathfile, folderPath+toFilename(t)+".txt")

			return string(content)
		}
	}

	return ""

}

func RemoveFromCache(cid string) {

	// Check params

	if len(cid) == 0 {
		return
	}

	// Remove cache

	folderPath := "./cache/" + cid + "/"
	os.RemoveAll(folderPath)

}

func RemoveDeprecatedFromCache() {

	var paths []string
	err := filepath.Walk("./cache", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	for _, path := range paths {
		filename := filepath.Base(path)
		name := strings.TrimSuffix(filename, filepath.Ext(filename))
		timestamp, t, err := fromFilename(name)
		if err != nil {
			os.RemoveAll(filepath.Dir(path))
			continue
		}
		if timestamp+t < time.Now().UTC().Unix() {
			os.RemoveAll(filepath.Dir(path))
		}
	}

}

// Helpers

func createFolder(folderPath string) error {

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return err

}

func toFilename(t int64) string {

	return fmt.Sprintf("%d:%d", time.Now().UTC().Unix(), t)

}

func fromFilename(name string) (int64, int64, error) {

	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("cache error: invalid filename")
	}
	timestamp, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("cache error: %v", err)
	}
	time, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("cache error: %v", err)
	}
	return timestamp, time, nil

}
