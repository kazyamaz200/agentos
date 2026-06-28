// Copyright 2026 AgentOS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sandbox

// Policy defines sandbox write restrictions and file size limits.
type Policy struct {
	DenyWritePatterns []string
	MaxFileSize       int64
}

// NewPolicy returns a Policy with default deny patterns for secret files and
// a 1 MB max file size.
func NewPolicy() *Policy {
	return &Policy{
		DenyWritePatterns: []string{
			".env", ".env.*",
			"*.pem", "id_rsa", "id_ed25519",
		},
		MaxFileSize: 1024 * 1024,
	}
}
