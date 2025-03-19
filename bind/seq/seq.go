// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package seq implements the machine-dependent seq serialization format.
//
// Implementations of Transact and FinalizeRef are provided by a
// specific foreign language binding package, e.g. go.mobile/bind/java.
//
// Designed only for use by the code generated by gobind. Don't try to
// use this directly.
package seq // import "github.com/konglong147/newgomomas/bind/seq"

import _ "github.com/konglong147/newgomomas/internal/mobileinit"

// FinalizeRef is the finalizer used on foreign objects.
var FinalizeRef func(ref *Ref)

// IncRef increments the foreign reference count for ref while it is in transit.
// The count is decremented after the ref is received and translated on the foreign side.
var IncForeignRef func(refnum int32)
