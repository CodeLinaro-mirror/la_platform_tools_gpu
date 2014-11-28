// Copyright (C) 2014 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package binary implements encoding and decoding of various primitive
// data types to and from a binary stream. The package holds BitStream for packing and unpacking
// sequences of bits, Float16 for dealing with 16 bit floating-point values and Encoder/Decoder for
// encoding and decoding various value types to a binary stream.
//
// Encoding details
//
// binary.Encoder and binary.Decoder provide a symmetrical pair of methods for encoding and decoding
// various data types to a binary stream. For performance reasons, each data type has a separate
// method for encoding and decoding rather than having a single pair of methods encoding and
// decoding boxed values in an interface{}. The binary format has been intentionally kept simple,
// with a preference for simplicity over compactness.
//
// Boolean values are encoded as single bytes, where 0 represents false and non-zero represents
// true.
//
// Integer values are encoded as little-endian words, with no compression or alignment. For example
// a 16 bit unsigned integer would be encoded as two bytes, the first containing the
// least-significant portion of the word.
//
// 32 bit floating-point numbers are first converted to a 32 bit unsigned integer using
// math.Float32bits before being encoded as described above.
//
// 64 bit floating-point numbers are first converted to a 64 bit unsigned integer using
// math.Float64bits before being encoded as described above.
//
// Strings values can be encoded in two different formats, either with String or CString.
//
// String encodes a 32 bit unsigned integer representing the length of the string in bytes followed
// by the string data encoded in UTF-8:
//   count                  uint32 // in bytes
//   data0, data1, data2... byte   // string in UTF-8
//
// CString encodes the UTF-8 encoded string data terminated with a single 0 byte:
//   data0, data1, data2... byte // string in UTF-8
//   term                   byte // 0x00
//
// Binary blobs are encoded using the Data method as a 32 bit unsigned integers representing the
// number of bytes followed by the sequence of data bytes:
//   count                  uint32 // in bytes
//   data0, data1, data2... byte   // blob data
//
// Objects are encoded in three different forms, depending on whether the object was nil, or was
// already encoded to the stream. If the object is non-nil and is encoded for the first time for a
// given Encoder then the object is encoded as:
//   key  uint16 // A unique identifier for the object instance. Valid range: [0x0000, 0xfffe]
//   type uint16 // A unique identifier for the type of the object
//   ...data...  // The object's data (length dependent on the object type)
//
// All subsequent encodings of the object are encoded as the 16 bit key identifier without any
// additional data:
//   key uint16 // An identifier of a previously encoded object instance
//
// If the object is nil, then the object is encoded as the 16 bit unsigned integer 0xffff without
// any additional data.
//
package binary
