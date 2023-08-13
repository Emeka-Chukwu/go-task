package directives

import (
	"context"
	"go-task/middlewares"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData, _ := middlewares.GetCurrentUserID(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Unauthorized",
		}
	}

	return next(ctx)
}
