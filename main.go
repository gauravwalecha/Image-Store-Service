package main

import (
	"domain"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

func main() {
	fmt.Println("Starting on 8080...")
	http.HandleFunc("/createAlbum", createAlbum)
	http.HandleFunc("/deleteAlbum", deleteAlbum)
	http.HandleFunc("/createImage", createImage)
	http.HandleFunc("/deleteImage", deleteImage)
	http.HandleFunc("/getAllImages", getAllImages)
	http.ListenAndServe(":8080", nil)
}

func deleteAlbum(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Deleting Public Folder of Album...")
	os.RemoveAll("./public")
}

func createAlbum(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating Public Folder of Album...")
	_, err := os.Stat("public")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("public/pics", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
		fmt.Println("Created Public Folder of Album..")
	} else {
		fmt.Println("Public Folder of Album already present..")
	}
}

func createImage(w http.ResponseWriter, req *http.Request) {
	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first...")
		return
	}
	decoder := json.NewDecoder(req.Body)

	addImage := domain.Image{}
	err = decoder.Decode(&addImage)
	if err != nil {
		panic(err)
	}

	fmt.Println("Adding Image into the Album...")
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullURLFile = addImage.PathOfImage
	buildFileName()

	file, err := os.Create(directory + "/public/pics/" + fileName)
	if err != nil {
		panic(err)
	}

	putFile(file, httpClient())
}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullURLFile)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Just Downloaded a file" + fileName)
}

func buildFileName() {
	fileURL, err := url.Parse(fullURLFile)
	fmt.Println(fileURL)
	if err != nil {
		panic(err)
	}
	path := fileURL.Path
	fmt.Println(path)
	segments := strings.Split(path, "/")
	fmt.Println(segments)
	fmt.Println(len(segments))
	fileName = segments[len(segments)-1]
	fmt.Println(fileName)
}

// to omit leading slashes
func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func deleteImage(w http.ResponseWriter, req *http.Request) {
	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first and insert some Images...")
		return
	}
	fmt.Println("Deleting Image from Album...")
	decoder := json.NewDecoder(req.Body)

	imageName := domain.ImageName{}
	err = decoder.Decode(&imageName)
	if err != nil {
		panic(err)
	}

	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(directory + "/public/pics/" + imageName.Name)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Image " + imageName.Name + " successfully deleted")
}

func fileExists(filename string) bool {
	_, err := os.Stat("public/pics/" + filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getAllImages(w http.ResponseWriter, req *http.Request) {

	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first...")
		return
	}
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(directory + "/public/pics")
	if err != nil {
		log.Fatal(err)
	}
	flag := 0
	fmt.Println("Images Stored in Album are: ")
	for _, file := range files {
		flag = 1
		fmt.Println(file.Name())
	}
	if flag == 0 {
		fmt.Println("Album is Empty.")
	}

}
