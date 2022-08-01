package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/MumAroi/go-graphql-example/controller"
	"github.com/MumAroi/go-graphql-example/graph/generated"
	"github.com/MumAroi/go-graphql-example/graph/model"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string) (*model.Book, error) {
	// panic(fmt.Errorf("not implemented"))
	var book model.Book
	book.Title = title
	book.Author = &model.Author{
		ID: author,
	}
	id, err := controller.CreateBook(book)
	if err != nil {
		return nil, err
	}
	idStr := strconv.Itoa(int(id))
	createdBook, _ := controller.GetBooksByID(&idStr)
	return createdBook, nil
}

// CreatAuthor is the resolver for the creatAuthor field.
func (r *mutationResolver) CreatAuthor(ctx context.Context, firstName string, lastName string) (*model.Author, error) {
	// panic(fmt.Errorf("not implemented"))
	var author model.Author
	author.FirstName = firstName
	author.LastName = lastName
	id, err := controller.CreateAuthor(author)
	if err != nil {
		return nil, err
	}
	return &model.Author{ID: strconv.FormatInt(id, 10), FirstName: author.FirstName, LastName: author.LastName}, nil
}

// BookByID is the resolver for the bookByID field.
func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	// panic(fmt.Errorf("not implemented"))
	book, err := controller.GetBooksByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// AllBooks is the resolver for the allBooks field.
func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	// panic(fmt.Errorf("not implemented"))
	books, err := controller.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

// AuthorByID is the resolver for the authorByID field.
func (r *queryResolver) AuthorByID(ctx context.Context, id *string) (*model.Author, error) {
	// panic(fmt.Errorf("not implemented"))
	author, err := controller.GetAuthorByID(id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

// AllAuthors is the resolver for the allAuthors field.
func (r *queryResolver) AllAuthors(ctx context.Context) ([]*model.Author, error) {
	// panic(fmt.Errorf("not implemented"))
	authors, err := controller.GetAllAuthors()
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
