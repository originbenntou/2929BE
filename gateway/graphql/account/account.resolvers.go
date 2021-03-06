package account

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/originbenntou/2929BE/gateway/graphql/account/generated"
	"github.com/originbenntou/2929BE/gateway/graphql/account/model"
	redis "github.com/originbenntou/2929BE/gateway/infrastructure/redis/client"
	"github.com/originbenntou/2929BE/gateway/interfaces/support"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
	"github.com/originbenntou/2929BE/shared/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"time"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, user model.User) (ok bool, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}

		registerCode(ctx, err)
	}()

	pbUser, err := r.accountClient.RegisterUser(ctx, &pbAccount.RegisterUserRequest{
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		CompanyId: uint64(user.CompanyID),
	})
	if err != nil {
		return false, err
	}

	if pbUser == nil {
		return false, nil
	}

	return pbUser.UserId > 0, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user model.User) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VerifyUser(ctx context.Context, email string, password string) (token string, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}

		registerCode(ctx, err)
	}()

	pbRes, err := r.accountClient.VerifyUser(ctx, &pbAccount.VerifyUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	token = pbRes.Token
	uid := pbRes.User.Id
	cid := pbRes.User.CompanyId

	// set id, company_id in Redis to session
	if err = redis.TokenClient.HSet(token, map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}).Err(); err != nil {
		return "", err
	}
	// expire 1 hour
	if err = redis.TokenClient.Expire(token, time.Hour*1).Err(); err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) RecoveryUser(ctx context.Context, email string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func registerCode(ctx context.Context, err error) {
	if err != nil {
		graphql.RegisterExtension(ctx, "code", status.Code(err))
	} else {
		graphql.RegisterExtension(ctx, "code", 0)
	}

	return
}
