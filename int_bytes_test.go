// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package int_bytes

import (
	"fmt"
	"github.com/sabey/unittest"
	"testing"
)

const (
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func TestInt16(t *testing.T) {
	fmt.Println("TestInt16()")
	unittest.EqualsExact(t, ToInt16([...]byte{0, 0}), int16(0))
	unittest.EqualsExact(t, ToInt16([...]byte{128, 0}), int16(MinInt16))
	unittest.EqualsExact(t, ToInt16([...]byte{127, 255}), int16(MaxInt16))
	unittest.EqualsExact(t, ToInt16([...]byte{255, 255}), int16(-1))
	unittest.EqualsExact(t, ToInt16(FromInt16(MinInt16)), int16(MinInt16))
	unittest.EqualsExact(t, ToInt16(FromInt16(MaxInt16)), int16(MaxInt16))
}
func TestInt32(t *testing.T) {
	fmt.Println("TestInt32()")
	unittest.EqualsExact(t, ToInt32([...]byte{0, 0, 0, 0}), int32(0))
	unittest.EqualsExact(t, ToInt32([...]byte{128, 0, 0, 0}), int32(MinInt32))
	unittest.EqualsExact(t, ToInt32([...]byte{127, 255, 255, 255}), int32(MaxInt32))
	unittest.EqualsExact(t, ToInt32([...]byte{255, 255, 255, 255}), int32(-1))
	unittest.EqualsExact(t, ToInt32(FromInt32(MinInt32)), int32(MinInt32))
	unittest.EqualsExact(t, ToInt32(FromInt32(MaxInt32)), int32(MaxInt32))
}
func TestInt64(t *testing.T) {
	fmt.Println("TestInt64()")
	unittest.EqualsExact(t, ToInt64([...]byte{0, 0, 0, 0, 0, 0, 0, 0}), int64(0))
	unittest.EqualsExact(t, ToInt64([...]byte{128, 0, 0, 0, 0, 0, 0, 0}), int64(MinInt64))
	unittest.EqualsExact(t, ToInt64([...]byte{127, 255, 255, 255, 255, 255, 255, 255}), int64(MaxInt64))
	unittest.EqualsExact(t, ToInt64([...]byte{255, 255, 255, 255, 255, 255, 255, 255}), int64(-1))
	unittest.EqualsExact(t, ToInt64(FromInt64(MinInt64)), int64(MinInt64))
	unittest.EqualsExact(t, ToInt64(FromInt64(MaxInt64)), int64(MaxInt64))
}
func TestUInt16(t *testing.T) {
	fmt.Println("TestUInt16()")
	unittest.EqualsExact(t, ToUInt16([...]byte{0, 0}), uint16(0))
	unittest.EqualsExact(t, ToUInt16([...]byte{255, 255}), uint16(MaxUint16))
	unittest.EqualsExact(t, ToUInt16(FromUInt16(MaxUint16)), uint16(MaxUint16))
}
func TestUInt32(t *testing.T) {
	fmt.Println("TestUInt32()")
	unittest.EqualsExact(t, ToUInt32([...]byte{0, 0, 0, 0}), uint32(0))
	unittest.EqualsExact(t, ToUInt32([...]byte{255, 255, 255, 255}), uint32(MaxUint32))
	unittest.EqualsExact(t, ToUInt32(FromUInt32(MaxUint32)), uint32(MaxUint32))
}
func TestUInt64(t *testing.T) {
	fmt.Println("TestUInt64()")
	unittest.EqualsExact(t, ToUInt64([...]byte{0, 0, 0, 0, 0, 0, 0, 0}), uint64(0))
	unittest.EqualsExact(t, ToUInt64([...]byte{255, 255, 255, 255, 255, 255, 255, 255}), uint64(MaxUint64))
	unittest.EqualsExact(t, ToUInt64(FromUInt64(MaxUint64)), uint64(MaxUint64))
}
