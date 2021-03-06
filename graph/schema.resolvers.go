package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql-api-mysql-gorm/graph/generated"
	"github.com/go-graphql-api-mysql-gorm/graph/model"
)



func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {


	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount: input.OrderAmount,
		Items: mapItemsFromInput(input.Items),
	}

	 r.DB.Create(&order)


	return &order, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	
	updateOrder := model.Order{
		ID: orderID,
		CustomerName: input.CustomerName,
		OrderAmount: input.OrderAmount,
		Items: mapItemsFromInput(input.Items),
	}

	r.DB.Save(&updateOrder)

	return &updateOrder, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	r.DB.Where("order_id = ?", orderID).Delete(&model.Order{})
	r.DB.Where("order_id = ?", orderID).Delete(&model.Item{})

	return true, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {

	orders := []*model.Order{}
	r.DB.Set("gorm:auto_preload", true).Find(&orders)

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


func mapItemsFromInput(itemsInput []*model.ItemInput) []*model.Item {
	 items := []*model.Item{}
	
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
		})
	}
	return items
}

