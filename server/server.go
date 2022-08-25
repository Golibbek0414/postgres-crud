package server

import (
	"net/http"
	"postgres-gin-crud/entity"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repo Repository
}

func New(repo Repository) Server {
	return Server{
		repo: repo,
	}
}

func (s Server) CreateBook(c *gin.Context) {
	var request CreateBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})
		return
	}

	b := entity.NewBook(request.Title, entity.Author{
		ID: request.AuthorID,
	})
	if err := s.repo.CreateBook(c.Request.Context(), b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error2": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (s Server) CreateAuthor(c *gin.Context) {
	var request CreateAuthorRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	b := entity.NewAuthor(request.Name)
	if err := s.repo.CreateAuthor(c.Request.Context(), b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (s Server) GetAuthor(c *gin.Context) {
	h := c.Query("ID")
	author, err := s.repo.GetAuthor(c.Request.Context(), h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"author_name": author.Name,
	})
}

func (s Server) GetBook(c *gin.Context) {
	h := c.Query("name")
	book, err := s.repo.GetBook(c.Request.Context(), h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	book.Author, err = s.repo.GetAuthor(c.Request.Context(), book.Author.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Book name":   book.Name,
		"Book author": book.Author.Name,
	})
}

func (s Server) ListBooks(c *gin.Context) {
	books, err := s.repo.ListBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	for num, val := range books {
		c.JSON(http.StatusOK, gin.H{
			"Num":       num + 1,
			"book_name": val.Name,
		})
	}
}

func (s Server) ListAuthors(c *gin.Context) {
	author, err := s.repo.ListAuthors(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	for num, val := range author {
		c.JSON(http.StatusOK, gin.H{
			"Num":         num + 1,
			"author_name": val.Name,
		})
	}
}

func(s Server) ListBooksByAuthor(c *gin.Context){
	authorId := c.Query("authorID")
	books, err := s.repo.ListBooksByAuthor(c.Request.Context(), authorId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),			
		})
		return
	}
	books[0].Author, err = s.repo.GetAuthor(c.Request.Context(), books[0].Author.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, book := range books{
		c.JSON(http.StatusOK, gin.H{
			"Book_name": book.Name,
			"Book_author": books[0].Author.Name,
		})
	}
}

func(s Server) DeleteBook(c *gin.Context){
	bookId := c.Query("id")
	if err := s.repo.DeleteBook(c.Request.Context(), bookId); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Deleted": true,
	})
}

func (s Server) DeleteAuthor(c *gin.Context){
	authorId := c.Query("id")
	if err := s.repo.DeleteAuthor(c.Request.Context(), authorId); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} 
	c.JSON(http.StatusOK, gin.H{
		"Deleted_Author": true,
	})
}

func(s Server) Updatelibrary(c *gin.Context){
	id := c.Query("id")
	name := c.Query("name")
	if err := s.repo.Updatelibrary(c.Request.Context(), name, id); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Updated": true,
	})
}

func NewRouter(repo Repository) {
	r := gin.Default()
	h := Server{repo: repo}
	r.POST("/Author", h.CreateAuthor)
	r.POST("/Book", h.CreateBook)
	r.GET("/book", h.GetBook)
	r.GET("/ListB", h.ListBooks)
	r.GET("/ListA", h.ListAuthors)
	r.GET("/ListBooks", h.ListBooksByAuthor)
	r.DELETE("/bok", h.DeleteBook)
	r.DELETE("/author", h.DeleteAuthor)
	r.PUT("/update", h.Updatelibrary)
	r.Run(":6070")
}
