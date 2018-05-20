// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodel

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testPosts(t *testing.T) {
	t.Parallel()

	query := Posts(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPostsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = post.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Posts(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PostSlice{post}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPostsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PostExists(tx, post.ID)
	if err != nil {
		t.Errorf("Unable to check if Post exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PostExistsG to return true, but got false.")
	}
}
func testPostsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	postFound, err := FindPost(tx, post.ID)
	if err != nil {
		t.Error(err)
	}

	if postFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPostsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Posts(tx).Bind(post); err != nil {
		t.Error(err)
	}
}

func testPostsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Posts(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPostsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = postOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Posts(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPostsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = postOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func postBeforeInsertHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterInsertHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterSelectHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeUpdateHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterUpdateHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeDeleteHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterDeleteHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postBeforeUpsertHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func postAfterUpsertHook(e boil.Executor, o *Post) error {
	*o = Post{}
	return nil
}

func testPostsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Post{}
	o := &Post{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, postDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Post object: %s", err)
	}

	AddPostHook(boil.BeforeInsertHook, postBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	postBeforeInsertHooks = []PostHook{}

	AddPostHook(boil.AfterInsertHook, postAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	postAfterInsertHooks = []PostHook{}

	AddPostHook(boil.AfterSelectHook, postAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	postAfterSelectHooks = []PostHook{}

	AddPostHook(boil.BeforeUpdateHook, postBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	postBeforeUpdateHooks = []PostHook{}

	AddPostHook(boil.AfterUpdateHook, postAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	postAfterUpdateHooks = []PostHook{}

	AddPostHook(boil.BeforeDeleteHook, postBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	postBeforeDeleteHooks = []PostHook{}

	AddPostHook(boil.AfterDeleteHook, postAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	postAfterDeleteHooks = []PostHook{}

	AddPostHook(boil.BeforeUpsertHook, postBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	postBeforeUpsertHooks = []PostHook{}

	AddPostHook(boil.AfterUpsertHook, postAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	postAfterUpsertHooks = []PostHook{}
}
func testPostsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx, postColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostToManyComments(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Post
	var b, c Comment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, commentDBTypes, false, commentColumnsWithDefault...)
	randomize.Struct(seed, &c, commentDBTypes, false, commentColumnsWithDefault...)

	b.PostID = a.ID
	c.PostID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	comment, err := a.Comments(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range comment {
		if v.PostID == b.PostID {
			bFound = true
		}
		if v.PostID == c.PostID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PostSlice{&a}
	if err = a.L.LoadComments(tx, false, (*[]*Post)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Comments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Comments = nil
	if err = a.L.LoadComments(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Comments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", comment)
	}
}

func testPostToManyAddOpComments(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Post
	var b, c, d, e Comment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Comment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Comment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddComments(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PostID {
			t.Error("foreign key was wrong value", a.ID, first.PostID)
		}
		if a.ID != second.PostID {
			t.Error("foreign key was wrong value", a.ID, second.PostID)
		}

		if first.R.Post != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Post != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Comments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Comments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Comments(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPostsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = post.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPostsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PostSlice{post}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPostsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Posts(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	postDBTypes = map[string]string{`Author`: `text`, `Body`: `text`, `CreatedAt`: `timestamp with time zone`, `ID`: `integer`, `Notes`: `text`, `Title`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testPostsUpdate(t *testing.T) {
	t.Parallel()

	if len(postColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err = post.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPostsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(postColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	post := &Post{}
	if err = randomize.Struct(seed, post, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, post, postDBTypes, true, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(postColumns, postPrimaryKeyColumns) {
		fields = postColumns
	} else {
		fields = strmangle.SetComplement(
			postColumns,
			postPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(post))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PostSlice{post}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPostsUpsert(t *testing.T) {
	t.Parallel()

	if len(postColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	post := Post{}
	if err = randomize.Struct(seed, &post, postDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = post.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err := Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &post, postDBTypes, false, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err = post.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err = Posts(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}