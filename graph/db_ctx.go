// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package graph

import (
	"context"

	"github.com/volatiletech/sqlboiler/boil"
)

func (r *Resolver) db(ctx context.Context) boil.Executor {
	return r.sqldb
}
