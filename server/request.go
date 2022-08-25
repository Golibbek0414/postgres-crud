package server

type CreateBookRequest struct {
	Title    string `json:"title"`
	AuthorID string `json:"author_id"`
}

type CreateAuthorRequest struct {
	ID string
	Name string `json: "name"`
}

type GetBookRequest struct {

}
