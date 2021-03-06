package graphql

import (
	"context"
	"fmt"
	"github.com/hunter1271/todos/database"
	"strconv"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	queries *database.Queries
}

func NewResolver(queries *database.Queries) *Resolver {
	return &Resolver{queries: queries}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*database.Todo, error) {
	todo, err := r.queries.CreateTodo(ctx, input.Text)

	return &todo, err
}

func (r *mutationResolver) UpdateTodoDone(ctx context.Context, id string, isDone bool) (*database.Todo, error) {
	intId, err := parseID(&id)
	if err != nil {
		return nil, err
	}
	todo, err := r.queries.UpdateTodoDone(ctx, database.UpdateTodoDoneParams{ID: intId, IsDone: isDone})

	return &todo, err
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*database.Todo, error) {
	intId, err := parseID(&id)
	if err != nil {
		return nil, err
	}
	todo, err := r.queries.DeleteTodo(ctx, intId)

	return &todo, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*database.Todo, error) {
	todos, err := r.queries.ListTodos(ctx)
	var list []*database.Todo
	for key, _ := range todos {
		list = append(list, &todos[key])
	}

	return list, err
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) ID(ctx context.Context, obj *database.Todo) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

func parseID(id *string) (int32, error) {
	int64Id, err := strconv.ParseInt(*id, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(int64Id), nil
}
