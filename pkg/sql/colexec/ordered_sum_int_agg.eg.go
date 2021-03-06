// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"strings"
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func newSumIntOrderedAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &sumIntInt16OrderedAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &sumIntInt32OrderedAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &sumIntInt64OrderedAggAlloc{aggAllocBase: allocBase}, nil
		}
	default:
		return nil, errors.Errorf("unsupported sum %s agg type %s", strings.ToLower("Int"), t.Name())
	}
}

type sumIntInt16OrderedAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumIntInt16OrderedAgg{}

const sizeOfSumIntInt16OrderedAgg = int64(unsafe.Sizeof(sumIntInt16OrderedAgg{}))

func (a *sumIntInt16OrderedAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int64()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumIntInt16OrderedAgg) Reset() {
	a.scratch.curIdx = 0
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumIntInt16OrderedAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumIntInt16OrderedAgg) SetOutputIndex(idx int) {
	a.scratch.curIdx = idx
}

func (a *sumIntInt16OrderedAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int16(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt16OrderedAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumIntInt16OrderedAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumIntInt16OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt16OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt16OrderedAggAlloc{}

func (a *sumIntInt16OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumIntInt16OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumIntInt16OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt32OrderedAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumIntInt32OrderedAgg{}

const sizeOfSumIntInt32OrderedAgg = int64(unsafe.Sizeof(sumIntInt32OrderedAgg{}))

func (a *sumIntInt32OrderedAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int64()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumIntInt32OrderedAgg) Reset() {
	a.scratch.curIdx = 0
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumIntInt32OrderedAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumIntInt32OrderedAgg) SetOutputIndex(idx int) {
	a.scratch.curIdx = idx
}

func (a *sumIntInt32OrderedAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int32(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt32OrderedAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumIntInt32OrderedAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumIntInt32OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt32OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt32OrderedAggAlloc{}

func (a *sumIntInt32OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumIntInt32OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumIntInt32OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt64OrderedAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumIntInt64OrderedAgg{}

const sizeOfSumIntInt64OrderedAgg = int64(unsafe.Sizeof(sumIntInt64OrderedAgg{}))

func (a *sumIntInt64OrderedAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int64()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumIntInt64OrderedAgg) Reset() {
	a.scratch.curIdx = 0
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumIntInt64OrderedAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumIntInt64OrderedAgg) SetOutputIndex(idx int) {
	a.scratch.curIdx = idx
}

func (a *sumIntInt64OrderedAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int64(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.scratch.nulls.SetNull(a.scratch.curIdx)
					} else {
						a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroInt64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt64OrderedAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumIntInt64OrderedAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumIntInt64OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt64OrderedAgg
}

var _ aggregateFuncAlloc = &sumIntInt64OrderedAggAlloc{}

func (a *sumIntInt64OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumIntInt64OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumIntInt64OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
