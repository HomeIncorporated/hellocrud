// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/choonkeat/hellocrud/golang/example/dbmodel"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// SchemaTypePost is the GrpahQL schema type for Post
var SchemaTypePost = `
# Post is a resource type
type Post {
	
		# Convenient GUID for ReactJS component @key attribute
		rowId: String!
		id: ID!
		title: String!
		author: String!
		body: String!
		notes: String
		createdAt: Time
		updatedAt: Time
		# comments has a many-to-one relationship with Post
		comments: CommentsCollection
}

type PostsCollection {
	nodes: [Post!]!
}
`

// Post is an object to back GraphQL type
type Post struct {
	model dbmodel.Post
	db    boil.Executor
}

// GraphQL friendly getters

// RowID is the Post GUID
func (o Post) RowID() string {
	return fmt.Sprintf("Post%d", o.model.ID) // int
}

// ID is the Post id
func (o Post) ID() graphql.ID {
	return graphql.ID(fmt.Sprintf("%d", o.model.ID)) // int
}

// Title is the Post title
func (o Post) Title() string {
	return o.model.Title // string
}

// Author is the Post author
func (o Post) Author() string {
	return o.model.Author // string
}

// Body is the Post body
func (o Post) Body() string {
	return o.model.Body // string
}

// Notes is the Post notes
func (o Post) Notes() *string {
	return o.model.Notes.Ptr() // null.String
}

// CreatedAt is the Post created_at
func (o Post) CreatedAt() *graphql.Time {
	if o.model.CreatedAt.Valid {
		return &graphql.Time{Time: o.model.CreatedAt.Time}
	}
	return nil // null.Time
}

// UpdatedAt is the Post updated_at
func (o Post) UpdatedAt() *graphql.Time {
	if o.model.UpdatedAt.Valid {
		return &graphql.Time{Time: o.model.UpdatedAt.Time}
	}
	return nil // null.Time
}

// Comments returns the list of Comments that has a foreign key pointing to Post
func (o Post) Comments() *CommentsCollection {
	if o.model.R == nil || o.model.R.Comments == nil {
		return nil
	}

	result := &CommentsCollection{}
	for _, it := range o.model.R.Comments {
		result.nodes = append(result.nodes, Comment{model: *it, db: o.db})
	}
	return result
}

// PostsCollection is the struct representing a collection of GraphQL types
type PostsCollection struct {
	nodes []Post
	// future meta data goes here, e.g. count
}

// Nodes returns the list of GraphQL types
func (c PostsCollection) Nodes(ctx context.Context) []Post {
	return c.nodes
}

// SchemaCreatePostInput is the schema create input for Post
var SchemaCreatePostInput = `
# CreatePostInput is a create input type for Post resource
input CreatePostInput {
	
	  title: String!
	  author: String!
	  body: String!
	  notes: String
}
`

// createPostInput is an object to back Post mutation (create) input type
type createPostInput struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Body   string  `json:"body"`
	Notes  *string `json:"notes"`
}

// SchemaUpdatePostInput is the schema update input for Post
var SchemaUpdatePostInput = `
input UpdatePostInput {
	
	  title: String!
	  author: String!
	  body: String!
	  notes: String
}
`

// updatePostInput is an object to back Post mutation (update) input type
type updatePostInput struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Body   string  `json:"body"`
	Notes  *string `json:"notes"`
}

// searchPostArgs is an object to back Post search arguments type
type searchPostArgs struct {
	Title     *string       `json:"title"`
	Author    *string       `json:"author"`
	Body      *string       `json:"body"`
	Notes     *string       `json:"notes"`
	CreatedAt *graphql.Time `json:"created_at"`
	UpdatedAt *graphql.Time `json:"updated_at"`
}

// QueryMods returns a list of QueryMod based on the struct values
func (s *searchPostArgs) QueryMods() []qm.QueryMod {
	mods := []qm.QueryMod{}
	// Get reflect value
	v := reflect.ValueOf(s).Elem()
	// Iterate struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i) // StructField
		value := v.Field(i)        // Value
		if value.IsNil() || !value.IsValid() {
			// Skip if field is nil
			continue
		}

		// Get column name from tags
		column, hasColumnName := field.Tag.Lookup("json")
		// Skip if no DB definition
		if !hasColumnName {
			continue
		}

		operator := "="
		val := value.Elem().Interface()
		if dataType := field.Type.String(); (dataType == "string" || dataType == "*string") &&
			val.(string) != "" {
			operator = "LIKE"
			val = fmt.Sprintf("%%%s%%", val)
		}
		mods = append(mods, qm.And(fmt.Sprintf("%s %s ?", column, operator), val))
	}
	return mods
}

// SchemaSearchPostInput is the schema search input for Post
var SchemaSearchPostInput = `
# SearchPostInput is a search input/arguments type for Post resources
input SearchPostInput {
	
	  title: String
	  author: String
	  body: String
	  notes: String
}
`

// searchPostInput is an object to back Post search arguments input type
type searchPostInput struct {
	Title     *string       `json:"title"`
	Author    *string       `json:"author"`
	Body      *string       `json:"body"`
	Notes     *string       `json:"notes"`
	CreatedAt *graphql.Time `json:"created_at"`
	UpdatedAt *graphql.Time `json:"updated_at"`
}

// QueryMods returns a list of QueryMod based on the struct values
func (s *searchPostInput) QueryMods() []qm.QueryMod {
	mods := []qm.QueryMod{}
	// Get reflect value
	v := reflect.ValueOf(s).Elem()
	// Iterate struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i) // StructField
		value := v.Field(i)        // Value
		if value.IsNil() || !value.IsValid() {
			// Skip if field is nil
			continue
		}

		// Get column name from tags
		column, hasColumnName := field.Tag.Lookup("json")
		// Skip if no DB definition
		if !hasColumnName {
			continue
		}

		operator := "="
		val := value.Elem().Interface()
		if dataType := field.Type.String(); (dataType == "string" || dataType == "*string") &&
			val.(string) != "" {
			operator = "LIKE"
			val = fmt.Sprintf("%%%s%%", val)
		}
		mods = append(mods, qm.And(fmt.Sprintf("%s %s ?", column, operator), val))
	}
	return mods
}

// AllPosts retrieves Posts based on the provided search parameters
func (r *Resolver) AllPosts(ctx context.Context, args struct {
	Since    *graphql.ID
	PageSize int32
	Search   *searchPostInput
}) (PostsCollection, error) {
	result := PostsCollection{}
	mods := []qm.QueryMod{
		qm.Limit(int(args.PageSize)),
		// TODO: Add eager loading based on requested fields
		qm.Load("Comments"),
	}

	if args.Since != nil {
		s := string(*args.Since)
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return result, err
		}
		mods = append(mods, qm.Offset(int(i)))
	}
	if args.Search != nil {
		mods = append(mods, args.Search.QueryMods()...)
	}
	slice, err := dbmodel.Posts(r.db, mods...).All()
	if err != nil {
		return result, errors.Wrapf(err, "allPosts(%#v)", args)
	}
	for _, m := range slice {
		result.nodes = append(result.nodes, Post{model: *m, db: r.db})
	}

	return result, nil
}

// PostByID retrieves Post by ID
func (r *Resolver) PostByID(ctx context.Context, args struct {
	ID graphql.ID
}) (Post, error) {
	result := Post{}
	i, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return result, errors.Wrapf(err, "PostByID(%#v)", args)
	}
	id := int(i)

	mods := []qm.QueryMod{
		qm.Where("id = ?", id),
		// TODO: Add eager loading based on requested fields
		qm.Load("Comments")}

	m, err := dbmodel.Posts(r.db, mods...).One()
	if err != nil {
		return result, errors.Wrapf(err, "PostByID(%#v)", args)
	} else if m == nil {
		return result, errors.New("not found")
	}
	return Post{model: *m, db: r.db}, nil
}

// CreatePost creates a Post based on the provided input
func (r *Resolver) CreatePost(ctx context.Context, args struct {
	Input createPostInput
}) (Post, error) {
	result := Post{}
	m := dbmodel.Post{}
	data, err := json.Marshal(args.Input)
	if err != nil {
		return result, errors.Wrapf(err, "json.Marshal(%#v)", args.Input)
	}
	if err = json.Unmarshal(data, &m); err != nil {
		return result, errors.Wrapf(err, "json.Unmarshal(%s)", data)
	}

	if err := m.Insert(r.db); err != nil {
		return result, errors.Wrapf(err, "createPost(%#v)", m)
	}
	return Post{model: m, db: r.db}, nil
}

// UpdatePostByID updates a Post based on the provided ID and input
func (r *Resolver) UpdatePostByID(ctx context.Context, args struct {
	ID    graphql.ID
	Input updatePostInput
}) (Post, error) {
	result := Post{}
	i, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return result, errors.Wrapf(err, "PostByID(%#v)", args)
	}
	id := int(i)

	m, err := dbmodel.FindPost(r.db, id)
	if err != nil {
		return result, errors.Wrapf(err, "updatePostByID(%#v)", args)
	} else if m == nil {
		return result, errors.New("not found")
	}
	data, err := json.Marshal(args.Input)
	if err != nil {
		return result, errors.Wrapf(err, "json.Marshal(%#v)", args.Input)
	}
	if err = json.Unmarshal(data, &m); err != nil {
		return result, errors.Wrapf(err, "json.Unmarshal(%s)", data)
	}

	if err := m.Update(r.db); err != nil {
		return result, errors.Wrapf(err, "updatePost(%#v)", m)
	}
	return Post{model: *m, db: r.db}, nil
}

// DeletePostByID deletes a Post based on the provided ID
func (r *Resolver) DeletePostByID(ctx context.Context, args struct {
	ID graphql.ID
}) (Post, error) {
	result := Post{}
	i, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return result, errors.Wrapf(err, "PostByID(%#v)", args)
	}
	id := int(i)

	m, err := dbmodel.FindPost(r.db, id)
	if err != nil {
		return result, errors.Wrapf(err, "updatePostByID(%#v)", args)
	} else if m == nil {
		return result, errors.New("not found")
	}
	if err := m.Delete(r.db); err != nil {
		return result, errors.Wrapf(err, "deletePostByID(%#v)", m)
	}
	return Post{model: *m, db: r.db}, nil
}
