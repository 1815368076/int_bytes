package int_bytes

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

func ToUInt16(
	b [2]byte,
) uint16 {
	i := uint16(b[0]) << 8
	i += uint16(b[1])
	return i
}
func FromUInt16(
	i uint16,
) [2]byte {
	return [2]byte{
		byte(i >> 8),
		byte(i),
	}
}
func ToInt16(
	b [2]byte,
) int16 {
	i := int16(b[0]) << 8
	i += int16(b[1])
	return i
}
func FromInt16(
	i int16,
) [2]byte {
	return [2]byte{
		byte(i >> 8),
		byte(i),
	}
}
func ToUInt32(
	b [4]byte,
) uint32 {
	i := uint32(b[0]) << 24
	i += uint32(b[1]) << 16
	i += uint32(b[2]) << 8
	i += uint32(b[3])
	return i
}
func FromUInt32(
	i uint32,
) [4]byte {
	return [4]byte{
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}
func ToInt32(
	b [4]byte,
) int32 {
	i := int32(b[0]) << 24
	i += int32(b[1]) << 16
	i += int32(b[2]) << 8
	i += int32(b[3])
	return i
}
func FromInt32(
	i int32,
) [4]byte {
	return [4]byte{
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}
func ToUInt64(
	b [8]byte,
) uint64 {
	i := uint64(b[0]) << 56
	i += uint64(b[1]) << 48
	i += uint64(b[2]) << 40
	i += uint64(b[3]) << 32
	i += uint64(b[4]) << 24
	i += uint64(b[5]) << 16
	i += uint64(b[6]) << 8
	i += uint64(b[7])
	return i
}
func FromUInt64(
	i uint64,
) [8]byte {
	return [8]byte{
		byte(i >> 56),
		byte(i >> 48),
		byte(i >> 40),
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}
func ToInt64(
	b [8]byte,
) int64 {
	i := int64(b[0]) << 56
	i += int64(b[1]) << 48
	i += int64(b[2]) << 40
	i += int64(b[3]) << 32
	i += int64(b[4]) << 24
	i += int64(b[5]) << 16
	i += int64(b[6]) << 8
	i += int64(b[7])
	return i
}
func FromInt64(
	i int64,
) [8]byte {
	return [8]byte{
		byte(i >> 56),
		byte(i >> 48),
		byte(i >> 40),
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}
