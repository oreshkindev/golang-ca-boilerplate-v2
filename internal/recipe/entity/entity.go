package entity

type Recipe struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type (
	Repository interface {
		Get() (string, error)
	}

	Usecase interface {
		Get() (string, error)
	}
)
