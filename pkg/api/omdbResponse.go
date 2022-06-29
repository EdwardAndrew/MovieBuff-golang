package api

type OMDBSearchResultRating struct {
	Source string
	Value  string
}

type OMDBSearchResult struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []OMDBSearchResultRating
	Metascore  string
	IMDBRating string `json:"imdbRating"`
	IMDBVotes  string `json:"imdbVotes"`
	IMDBId     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}
