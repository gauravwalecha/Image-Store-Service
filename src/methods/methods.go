package methods

import (
	"domain"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/emicklei/go-restful"
)

var (
	fileName    string
	fullURLFile string
)

// CreateAlbum -
func CreateAlbum(req *restful.Request, resp *restful.Response) {
	fmt.Println("Creating Album...")
	_, err := os.Stat("public")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("public/pics", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
		fmt.Println("Album Created..")
		resp.WriteEntity("Album Created..")
	} else {
		fmt.Println("Album already present..")
		resp.WriteEntity("Album already present..")
	}
}

// DeleteAlbum - 
func DeleteAlbum(req *restful.Request, resp *restful.Response) {
	fmt.Println("Deleting Album...")
	_, err := os.Stat("public")

	if os.IsNotExist(err) {
		fmt.Println("Already Deleted Album...")
		resp.WriteEntity("Already Deleted Album...")
		return
	} 
	os.RemoveAll("./public")
	fmt.Println("Album Deleted..")
	resp.WriteEntity("Album Deleted..")
	return
}

// CreateImage - 
func CreateImage(req *restful.Request, resp *restful.Response) {
	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first...")
		resp.WriteEntity("Sorry, Album is not present. Kindly make Album first...")
		return
	}

	addImage := domain.Image{}
	req.ReadEntity(&addImage)

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
	resp.WriteEntity(fileName + " Image Added to Album.")
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

// DeleteImage -
func DeleteImage(req *restful.Request, resp *restful.Response) {
	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first and insert some Images...")
		return
	}
	fmt.Println("Deleting Image from Album...")

	imageName := domain.ImageName{}
	req.ReadEntity(&imageName)

	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(directory + "/public/pics/" + imageName.Name)

	if err != nil {
		fmt.Println(err)
		resp.WriteEntity("Image not Present...")
		return
	}
	fmt.Println("Image " + imageName.Name + " successfully deleted")
	resp.WriteEntity("Image " + imageName.Name + " successfully deleted")
}

func fileExists(filename string) bool {
	_, err := os.Stat("public/pics/" + filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// GetAllImages -
func GetAllImages(req *restful.Request, resp *restful.Response) {

	_, err := os.Stat("public")
	if os.IsNotExist(err) {
		fmt.Println("Sorry, Album is not present. Kindly make Album first...")
		resp.WriteEntity("Sorry, Album is not present. Kindly make Album first...")
		return
	}
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		resp.WriteEntity(err)
	}
	files, err := ioutil.ReadDir(directory + "/public/pics")
	if err != nil {
		log.Fatal(err)
		resp.WriteEntity(err)
	}
	flag := 0
	fmt.Println("Images Stored in Album are: ")
	listOfImages := domain.ListOfImages{}
	for _, file := range files {
		flag = 1
		fmt.Println(file.Name())
		listOfImages.Images = append(listOfImages.Images, file.Name())
	}
	if flag == 0 {
		fmt.Println("Album is Empty.")
		resp.WriteEntity("Album is Empty.")
		return
	}
	resp.WriteEntity(listOfImages)
}
