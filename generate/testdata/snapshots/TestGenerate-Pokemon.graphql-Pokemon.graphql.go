// Code generated by github.com/obitech/genqlient, DO NOT EDIT.

package test

import (
	"github.com/obitech/genqlient/graphql"
	"github.com/obitech/genqlient/internal/testutil"
)

// GetPokemonSiblingsResponse is returned by GetPokemonSiblings on success.
type GetPokemonSiblingsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User GetPokemonSiblingsUser `json:"user"`
}

// GetUser returns GetPokemonSiblingsResponse.User, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsResponse) GetUser() GetPokemonSiblingsUser { return v.User }

// GetPokemonSiblingsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type GetPokemonSiblingsUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id               string                                   `json:"id"`
	Roles            []string                                 `json:"roles"`
	Name             string                                   `json:"name"`
	Pokemon          []testutil.Pokemon                       `json:"pokemon"`
	GenqlientPokemon []GetPokemonSiblingsUserGenqlientPokemon `json:"genqlientPokemon"`
}

// GetId returns GetPokemonSiblingsUser.Id, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetId() string { return v.Id }

// GetRoles returns GetPokemonSiblingsUser.Roles, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetRoles() []string { return v.Roles }

// GetName returns GetPokemonSiblingsUser.Name, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetName() string { return v.Name }

// GetPokemon returns GetPokemonSiblingsUser.Pokemon, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetPokemon() []testutil.Pokemon { return v.Pokemon }

// GetGenqlientPokemon returns GetPokemonSiblingsUser.GenqlientPokemon, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUser) GetGenqlientPokemon() []GetPokemonSiblingsUserGenqlientPokemon {
	return v.GenqlientPokemon
}

// GetPokemonSiblingsUserGenqlientPokemon includes the requested fields of the GraphQL type Pokemon.
type GetPokemonSiblingsUserGenqlientPokemon struct {
	Species string `json:"species"`
	Level   int    `json:"level"`
}

// GetSpecies returns GetPokemonSiblingsUserGenqlientPokemon.Species, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUserGenqlientPokemon) GetSpecies() string { return v.Species }

// GetLevel returns GetPokemonSiblingsUserGenqlientPokemon.Level, and is useful for accessing the field via an interface.
func (v *GetPokemonSiblingsUserGenqlientPokemon) GetLevel() int { return v.Level }

// __GetPokemonSiblingsInput is used internally by genqlient
type __GetPokemonSiblingsInput struct {
	Input testutil.Pokemon `json:"input"`
}

// GetInput returns __GetPokemonSiblingsInput.Input, and is useful for accessing the field via an interface.
func (v *__GetPokemonSiblingsInput) GetInput() testutil.Pokemon { return v.Input }

// The query or mutation executed by GetPokemonSiblings.
const GetPokemonSiblings_Operation = `
query GetPokemonSiblings ($input: PokemonInput!) {
	user(query: {hasPokemon:$input}) {
		id
		roles
		name
		pokemon {
			species
			level
		}
		genqlientPokemon: pokemon {
			species
			level
		}
	}
}
`

func GetPokemonSiblings(
	client graphql.Client,
	input testutil.Pokemon,
) (*GetPokemonSiblingsResponse, error) {
	req := &graphql.Request{
		OpName: "GetPokemonSiblings",
		Query:  GetPokemonSiblings_Operation,
		Variables: &__GetPokemonSiblingsInput{
			Input: input,
		},
	}
	var err error

	var data GetPokemonSiblingsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

