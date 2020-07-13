package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/suapapa/sharefit-gql-server/graph/generated"
	"github.com/suapapa/sharefit-gql-server/graph/model"
	"github.com/suapapa/sharefit-gql-server/internal/database"
)

func (r *membershipResolver) Users(ctx context.Context, obj *model.Membership) ([]*model.User, error) {
	var users []database.User
	if err := database.SharefitDB.Where("card_id = ?", obj.ID).Find(&users).Error; err != nil {
		return nil, err
	}

	var ret []*model.User
	for _, u := range users {
		ret = append(ret, &model.User{
			ID:          fmt.Sprint(u.ID),
			Name:        u.Name,
			PhoneNumber: u.PhoneNumber,
		})
	}

	return ret, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := database.User{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	database.SharefitDB.Create(&user)

	ret := model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	return &ret, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID *string, user model.NewUser) (*model.User, error) {
	var u database.User
	if err := database.SharefitDB.Where("id = ?", userID).Find(&u).Error; err != nil {
		return nil, err
	}

	u.Name = user.Name
	u.PhoneNumber = user.PhoneNumber

	database.SharefitDB.Save(&u)

	return &model.User{
		ID:          fmt.Sprint(u.ID),
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*model.Membership, error) {
	var cards []database.Card
	database.SharefitDB.Find(&cards)

	var ret []*model.Membership
	for _, v := range cards {
		ret = append(ret, &model.Membership{
			ID:       fmt.Sprint(v.ID),
			Training: v.Training,
			CurrCnt:  v.CurrCnt,
			TotalCnt: v.TotalCnt,
			Expiry:   v.Expiry,
			// UserIDs: ,
		})
	}

	return ret, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []database.User
	database.SharefitDB.Find(&users)

	var ret []*model.User
	for _, v := range users {
		ret = append(ret, &model.User{
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
		})
	}
	return ret, nil
}

func (r *queryResolver) User(ctx context.Context, userID *string) (*model.User, error) {
	var user database.User
	if err := database.SharefitDB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (r *queryResolver) Centers(ctx context.Context) ([]*model.Center, error) {
	var centers []database.Center
	database.SharefitDB.Find(&centers)

	var ret []*model.Center
	for _, v := range centers {
		database.SharefitDB.Where("center_id = ?", v.ID).Find(&v.Cards)
		var cards []*model.Membership
		for _, c := range v.Cards {
			cards = append(cards, &model.Membership{
				Training: c.Training,
				CurrCnt:  c.CurrCnt,
				TotalCnt: c.TotalCnt,
				// TODO: ???? should retrive users ????
			})
		}

		ret = append(ret, &model.Center{
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
			Memberships: cards,
		})
	}

	return ret, nil
}

// Membership returns generated.MembershipResolver implementation.
func (r *Resolver) Membership() generated.MembershipResolver { return &membershipResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type membershipResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
