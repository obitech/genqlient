// Code generated by github.com/obitech/genqlient, DO NOT EDIT.

package test

import (
	"github.com/obitech/genqlient/graphql"
	"github.com/obitech/genqlient/internal/testutil"
)

// QueryWithAliasResponse is returned by QueryWithAlias on success.
type QueryWithAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithAliasUser `json:"User"`
}

// GetUser returns QueryWithAliasResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithAliasResponse) GetUser() QueryWithAliasUser { return v.User }

// QueryWithAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID testutil.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	OtherID testutil.ID `json:"otherID"`
}

// GetID returns QueryWithAliasUser.ID, and is useful for accessing the field via an interface.
func (v *QueryWithAliasUser) GetID() testutil.ID { return v.ID }

// GetOtherID returns QueryWithAliasUser.OtherID, and is useful for accessing the field via an interface.
func (v *QueryWithAliasUser) GetOtherID() testutil.ID { return v.OtherID }

// The query or mutation executed by QueryWithAlias.
const QueryWithAlias_Operation = `
query QueryWithAlias {
	User: user {
		ID: id
		otherID: id
	}
}
`

func QueryWithAlias(
	client graphql.Client,
) (*QueryWithAliasResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithAlias",
		Query:  QueryWithAlias_Operation,
	}
	var err error

	var data QueryWithAliasResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

