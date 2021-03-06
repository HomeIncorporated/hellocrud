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

func testComments(t *testing.T) {
	t.Parallel()

	query := Comments(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCommentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = comment.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Comments(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{comment}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCommentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CommentExists(tx, comment.ID)
	if err != nil {
		t.Errorf("Unable to check if Comment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CommentExistsG to return true, but got false.")
	}
}
func testCommentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	commentFound, err := FindComment(tx, comment.ID)
	if err != nil {
		t.Error(err)
	}

	if commentFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCommentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Comments(tx).Bind(comment); err != nil {
		t.Error(err)
	}
}

func testCommentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Comments(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCommentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = commentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Comments(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCommentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = commentOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func commentBeforeInsertHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterInsertHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterSelectHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpdateHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpdateHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeDeleteHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterDeleteHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpsertHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpsertHook(e boil.Executor, o *Comment) error {
	*o = Comment{}
	return nil
}

func testCommentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Comment{}
	o := &Comment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, commentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Comment object: %s", err)
	}

	AddCommentHook(boil.BeforeInsertHook, commentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterInsertHook, commentAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	commentAfterInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterSelectHook, commentAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	commentAfterSelectHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpdateHook, commentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpdateHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpdateHook, commentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	commentAfterUpdateHooks = []CommentHook{}

	AddCommentHook(boil.BeforeDeleteHook, commentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	commentBeforeDeleteHooks = []CommentHook{}

	AddCommentHook(boil.AfterDeleteHook, commentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	commentAfterDeleteHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpsertHook, commentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpsertHook, commentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	commentAfterUpsertHooks = []CommentHook{}
}
func testCommentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx, commentColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentToOnePostUsingPost(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Comment
	var foreign Post

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PostID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Post(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CommentSlice{&local}
	if err = local.L.LoadPost(tx, false, (*[]*Comment)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Post == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Post = nil
	if err = local.L.LoadPost(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Post == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCommentToOneSetOpPostUsingPost(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Comment
	var b, c Post

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Post{&b, &c} {
		err = a.SetPost(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Post != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Comments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PostID != x.ID {
			t.Error("foreign key was wrong value", a.PostID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PostID))
		reflect.Indirect(reflect.ValueOf(&a.PostID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PostID != x.ID {
			t.Error("foreign key was wrong value", a.PostID, x.ID)
		}
	}
}
func testCommentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = comment.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCommentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{comment}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCommentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Comments(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	commentDBTypes = map[string]string{`Author`: `text`, `Body`: `text`, `CreatedAt`: `timestamp with time zone`, `ID`: `integer`, `Notes`: `text`, `PostID`: `integer`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testCommentsUpdate(t *testing.T) {
	t.Parallel()

	if len(commentColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if err = comment.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCommentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(commentColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	comment := &Comment{}
	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, comment, commentDBTypes, true, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(commentColumns, commentPrimaryKeyColumns) {
		fields = commentColumns
	} else {
		fields = strmangle.SetComplement(
			commentColumns,
			commentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(comment))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CommentSlice{comment}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCommentsUpsert(t *testing.T) {
	t.Parallel()

	if len(commentColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	comment := Comment{}
	if err = randomize.Struct(seed, &comment, commentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = comment.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err := Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &comment, commentDBTypes, false, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if err = comment.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err = Comments(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
