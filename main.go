package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	BASE_URL = "https://apicovidv2.reynadi.xyz/api/indonesia"
	LIST_URL = map[string]string{
		"kumulatif":       BASE_URL + "/",
		"kumulatif_more":  BASE_URL + "/more",
		"harian":          BASE_URL + "/harian",
		"provinsi":        BASE_URL + "/provinsi",
		"provinsi_more":   BASE_URL + "/provinsi/more",
		"provinsi_harian": BASE_URL + "/provinsi/harian",
	}
)

func main() {
	for k, v := range LIST_URL {
		fmt.Println("Downloading", k, "...")
		DownloadFile(fmt.Sprintf("results/%s.json", k), v)
	}
	WriteTimeLog()
}

func WriteTimeLog() {
	out, err := os.Create("results/last_update.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	out.WriteString(strconv.Itoa(int(time.Now().Unix())))
	fmt.Println("writed time log")
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
