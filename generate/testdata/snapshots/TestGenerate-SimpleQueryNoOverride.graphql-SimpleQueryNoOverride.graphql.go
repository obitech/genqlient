// Code generated by github.com/obitech/genqlient, DO NOT EDIT.

package test

import (
	"github.com/obitech/genqlient/graphql"
	"github.com/obitech/genqlient/internal/testutil"
)

// SimpleQueryNoOverrideResponse is returned by SimpleQueryNoOverride on success.
type SimpleQueryNoOverrideResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryNoOverrideUser `json:"user"`
}

// GetUser returns SimpleQueryNoOverrideResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideResponse) GetUser() SimpleQueryNoOverrideUser { return v.User }

// SimpleQueryNoOverrideUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryNoOverrideUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns SimpleQueryNoOverrideUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetId() testutil.ID { return v.Id }

// GetName returns SimpleQueryNoOverrideUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetName() string { return v.Name }

// The query or mutation executed by SimpleQueryNoOverride.
const SimpleQueryNoOverride_Operation = `
query SimpleQueryNoOverride {
	user {
		id
		name
	}
}
`

func SimpleQueryNoOverride(
	client graphql.Client,
) (*SimpleQueryNoOverrideResponse, error) {
	req := &graphql.Request{
		OpName: "SimpleQueryNoOverride",
		Query:  SimpleQueryNoOverride_Operation,
	}
	var err error

	var data SimpleQueryNoOverrideResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

