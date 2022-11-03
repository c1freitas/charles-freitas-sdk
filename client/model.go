package client

//Data is the base wrapper all endpoints return
//Each specific type implements this
type Data struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
}

type BookData struct {
	Data
	Books []Book `json:"docs"`
}
type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type ChapterData struct {
	Data
	Chapters []Chapter `json:"docs"`
}
type Chapter struct {
	Id          string `json:"_id"`
	ChapterName string `json:"chapterName"`
}

type MovieData struct {
	Data
	Movies []Movie `json:"docs"`
}
type Movie struct {
	Id                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           int     `json:"runtimeInMinutes"`
	BudgetInMillions           float32 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float32 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float32 `json:"rottenTomatoesScore"`
}

type CharacterData struct {
	Data
	Characters []Character `json:"docs"`
}
type Character struct {
	Id      string  `json:"_id"`
	Name    string  `json:"name"`
	Height  string  `json:"height"`
	Race    string  `json:"race"`
	Gender  string  `json:"gender"`
	Spouse  *string `json:"spouse,omitempty"`
	Birth   string  `json:"birth"`
	Death   string  `json:"death"`
	Realm   string  `json:"realm"`
	Hair    string  `json:"hair"`
	WikiUrl string  `json:"wikiUrl"`
}
