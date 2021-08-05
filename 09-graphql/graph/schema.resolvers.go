package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-go-demo/graph/generated"
	"graphql-go-demo/graph/model"
	"graphql-go-demo/repository"
	"strconv"
)

func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string) (*model.Book, error) {
	var book model.Book
	book.Title = title
	book.Author = &model.Author{
		ID: author,
	}

	id, err := repository.CreateBook(book)
	if err != nil {
		return nil, err
	}
	idStr := strconv.Itoa(int(id))
	createdBook, _ := repository.GetBooksByID(&idStr)
	return createdBook, nil
}

func (r *mutationResolver) CreatAuthor(ctx context.Context, firstName string, lastName string) (*model.Author, error) {
	var author model.Author
	author.FirstName = firstName
	author.LastName = lastName
	id, err := repository.CreateAuthor(author)
	if err != nil {
		return nil, err
	} else {
		return &model.Author{ID: strconv.FormatInt(id, 10), FirstName: author.FirstName, LastName: author.LastName}, nil
	}
}

func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	book, err := repository.GetBooksByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := repository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *queryResolver) AuthorByID(ctx context.Context, id *string) (*model.Author, error) {
	author, err := repository.GetAuthorByID(id)
	if err != nil {
		return nil, err
	} else {
		return author, nil
	}
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]*model.Author, error) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		return nil, err
	} else {
		return authors, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
