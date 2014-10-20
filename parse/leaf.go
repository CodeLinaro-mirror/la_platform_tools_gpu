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

package parse

import "io"

// Leaf nodes are part of the cst that cannot have child nodes, they represent
// a single token from the input.
type Leaf struct {
	fragment
	node
}

func (n *Leaf) WriteTo(w io.Writer) error {
	if err := n.prefix.WriteTo(w); err != nil {
		return err
	}
	if err := n.fragment.WriteTo(w); err != nil {
		return err
	}
	return n.suffix.WriteTo(w)
}
