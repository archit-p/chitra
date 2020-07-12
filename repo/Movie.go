package repo

import (
	"crypto/md5"
	"log"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type VideoFile struct {
	FileName	string		`json:"filename"`
	BaseName	string		`json:"basename"`
	Key			string		`json:"key"`
}

var VideoDict map[string]VideoFile

func valid(a string) bool {
	validtypes := []string{".mp4", ".mkv", ".webm"}

	for _, b := range validtypes {
		if b == a {
			return true
		}
	}

	return false
}

func getHashString(word string) string {
	hashBytes := md5.Sum([]byte(word))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString[:10]
}

func GetMovieList(root string) []VideoFile {
	var movies []VideoFile

	VideoDict = make(map[string]VideoFile)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !valid(filepath.Ext(path)) {
			return nil
		}
		m := VideoFile {
			FileName: path,
			BaseName: info.Name(),
			Key: getHashString(path),
		}
		movies = append(movies, m)

		VideoDict[m.Key] = m
		return nil
	})

	if err != nil {
		panic(err)
	}

	return movies
}

func GetMovieDetails(key string) VideoFile {
	m, ok := VideoDict[key]

	if !ok {
		log.Println("Movie not found in database")
	}

	return m
}

func GetMoviePath(key string) (string, error) {
	m, ok := VideoDict[key]

	if !ok {
		return "", fmt.Errorf("filesys: %s not found in store", key)
	}

	return m.FileName, nil
}
