package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph/generated"
	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph/model"
	"syreclabs.com/go/faker"
)

// DummyMutation returns a bunch of smart quotes.
func (r *mutationResolver) DummyMutation(ctx context.Context, input int) ([]string, error) {
	var res []string

	for i := 0; i < input; i++ {
		res = append(res, faker.Hacker().SaySomethingSmart())
	}

	return res, nil
}

// Health is the resolver for the health field.
func (r *queryResolver) Health(ctx context.Context, accessToken string) (*model.HealthPayload, error) {
	res := &model.HealthPayload{
		CanReachGraphQLProxy: true,
	}

	res.CanReachMicroservice, res.AccessTokenIsValid = getMicroHealth(accessToken)

	return res, nil
}

type microHealthResponse struct {
	CanReachMicroservice bool `json:"canReachMicroservice"`
	AccessTokenIsValid   bool `json:"accessTokenIsValid"`
}

func getMicroHealth(accessToken string) (canReachMicroservice, accessTokenIsValid bool) {
	url := fmt.Sprintf("http://localhost:8080/authorization/health?accessToken=%s", accessToken)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("error preparing http request: %v", err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error making http request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %v", err)
		}
	}(rsp.Body)

	rspBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	var microHealthRsp microHealthResponse

	if err = json.Unmarshal(rspBody, &microHealthRsp); err != nil {
		log.Fatalf("error unmarshalling response body: %v", err)
	}

	return microHealthRsp.CanReachMicroservice, microHealthRsp.AccessTokenIsValid
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
