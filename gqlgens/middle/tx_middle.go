package middle

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shiqiyue/gof/gorms"
	"gorm.io/gorm"
)

func NewTxMiddle(db *gorm.DB) graphql.ResponseMiddleware {
	return func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		operationContext := graphql.GetOperationContext(ctx)
		if operationContext != nil && operationContext.Doc != nil && operationContext.Doc.Operations != nil {
			operationDefinition := operationContext.Doc.Operations.ForName(operationContext.OperationName)
			if operationDefinition.Operation == "query" {
				return next(ctx)
			} else {
				withTx, tx := gorms.CreateTx(ctx, db)
				response := next(withTx)
				errors := graphql.GetErrors(ctx)
				if errors != nil {
					fmt.Println("rollback")
					tx.Rollback()
				} else {
					fmt.Println("commit")
					tx.Commit()
				}
				return response
			}
		} else {
			return next(ctx)
		}
	}
}
