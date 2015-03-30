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

// Package flat implements binary.Encoder and binary.Decoder. It does not
// keep track of objects it has seen before, so encoding an object twice will
// result in two copies both in the stream and when decoded. This also means it
// will hang if it is given cyclic data structures.

// Objects are encoded as:
//   type [20]byte // A unique identifier for the type of the object
//   ...data...    // The object's data (length dependent on the object type)
//
// If the object is nil, then the object is encoded as a null type.
//
package flat
