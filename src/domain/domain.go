package domain

// Image -
type Image struct {
	PathOfImage string `json:"pathOfImage" description:"Sample Image Path: https://i.pinimg.com/originals/2e/8a/41/2e8a4149d974c46e2da30589f03ecdf8.jpg"`
}

// ImageName -
type ImageName struct {
	Name string `json:"name" description:"Sample Image Name: 2e8a4149d974c46e2da30589f03ecdf8.jpg"`
}

// ListOfImages -
type ListOfImages struct {
	Images []string `json:"images" description:"List all the Images present in the Album."`
}
