package v1beta1

type Title string

type TitleWrapper struct {
	Title Title `json:"Title"`
}

type TitlesWrapper struct {
	Titles []Title `json:"Titles"`
}
