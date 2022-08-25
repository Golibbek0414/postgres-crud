package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"postgres-gin-crud/entity"
	"time"
)

type PostgreRepository struct {
	db *sql.DB
}

func NewDb(db *sql.DB) PostgreRepository {
	return PostgreRepository{
		db: db,
	}
}

func (p PostgreRepository) CreateBook(ctx context.Context, pr entity.Book) error {
	// rows, err := p.db.Query(
	// 	"SELECT * FROM 	books",
	// )
	// if err != nil {
	// 	return err
	// }
	// var books []entity.Book
	// for rows.Next() {
	// 	book := entity.Book{}
	// 	if err := rows.Scan(book.ID, book.Name, book.Author); err != nil {
	// 		return err
	// 	}
	// 	books = append(books, book)
	// }
	// for _, val := range books {
	// 	if val.Name == pr.Name {
	// 		return err
	// 	}
	// }
	s := fmt.Sprintf("INSERT INTO books VALUES('%s', '%s', '%s')", pr.ID, pr.Name, pr.Author.ID)
	_, err := p.db.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgreRepository) CreateAuthor(ctx context.Context, pr entity.Author) error {
	s := fmt.Sprintf("INSERT INTO authors VALUES('%s', '%s')", pr.ID, pr.Name)
	_, err := p.db.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgreRepository) GetAuthor(ctx context.Context, id string) (entity.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond*10)
	defer cancel()
	rows, err := p.db.Query("SELECT * FROM authors WHERE id=$1", id)
	if err != nil {
		return entity.Author{}, err
	}
	var author entity.Author
	for rows.Next() {
		if err := rows.Scan(&author.Name, &author.ID); err != nil {
			return entity.Author{}, err
		}
	}
	return author, err
}

func (p PostgreRepository) GetBook(ctx context.Context, name string) (entity.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Nanosecond*1)
	defer cancel()
	rows, err := p.db.QueryContext(ctx,
		"SELECT * FROM books WHERE title=$1", name,
	)
	if err != nil {
		return entity.Book{}, err
	}
	var books entity.Book
	for rows.Next() {
		if err := rows.Scan(&books.ID, &books.Name, &books.Author.ID); err != nil {
			return entity.Book{}, err
		}
	}
	return books, err
}

func (p PostgreRepository) ListBooks(ctx context.Context) ([]entity.Book, error) {
	rows, err := p.db.Query("SELECT * FROM books")
	if err != nil {
		return []entity.Book{}, err
	}
	var listbook []entity.Book
	for rows.Next() {
		book := entity.Book{}
		if err := rows.Scan(&book.ID, &book.Name, &book.Author.ID); err != nil {
			return []entity.Book{}, err
		}
		listbook = append(listbook, book)
	}
	return listbook, nil
}

func (p PostgreRepository) ListAuthors(ctx context.Context) ([]entity.Author, error) {
	rows, err := p.db.Query("SELECT * FROM authors")
	if err != nil {
		return []entity.Author{}, err
	}
	var authors []entity.Author
	for rows.Next() {
		author := entity.Author{}
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return []entity.Author{}, err
		}
		authors = append(authors, author)
	}
	return authors, err
}

func (p PostgreRepository) ListBooksByAuthor(ctx context.Context, authorID string) ([]entity.Book, error) {
	rows, err := p.db.Query("SELECT * FROM books WHERE author_id=$1", authorID)
	if err != nil {
		return []entity.Book{}, err
	}
	var books []entity.Book
	for rows.Next() {
		book := entity.Book{}
		if err := rows.Scan(&book.ID, &book.Name, &book.Author.ID); err != nil {
			return []entity.Book{}, err
		}
		books = append(books, book)
	}
	return books, err
}

func (p PostgreRepository) DeleteBook(ctx context.Context, id string) error {
	if _, err := p.db.Exec("DELETE FROM books WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}

func (p PostgreRepository) DeleteAuthor(ctx context.Context, id string) error {
	if _, err := p.db.Exec("DELETE FROM authors WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}

func (p PostgreRepository) Updatelibrary(ctx context.Context, name string, id string) error {
	if _, err := p.db.Exec("UPDATE books SET title=$1 WHERE id=$2", name, id); err != nil {
		return err
	}
	return nil
}
