package v1beta1

//Title is a location centric title that describes
//a group of images by their metadata
type Title string

type TitleWrapper struct {
	Title Title `json:"Title"`
}

type TitlesWrapper struct {
	Titles []Title `json:"Titles"`
}
