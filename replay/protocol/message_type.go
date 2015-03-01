/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package protocol

// MessageType defines the packet type sent from the replay system to the server.
type MessageType uint8

const (
	// MessageTypeGet is sent for a packet that requests a resource.
	MessageTypeGet = MessageType(0)
	// MessageTypePost is sent for a packet containing postback data.
	MessageTypePost = MessageType(1)
)
