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

// Package profile provides types and functions for managing agent profiles,
// including LLM configuration, tool permissions, and resource limits.
package profile

// ProfileService provides access to a Profile and convenience methods for
// checking tool and command permissions.
type ProfileService struct {
	profile *Profile
}

// NewProfileService returns a ProfileService wrapping the given Profile.
func NewProfileService(profile *Profile) *ProfileService {
	return &ProfileService{profile: profile}
}

// Profile returns the underlying Profile.
func (s *ProfileService) Profile() *Profile {
	return s.profile
}

// IsToolAllowed returns true if the tool named toolName is in the profile's
// allowed list, or if no allow list is configured (all tools allowed).
func (s *ProfileService) IsToolAllowed(toolName string) bool {
	if len(s.profile.Tools.Allow) == 0 {
		return true
	}
	for _, t := range s.profile.Tools.Allow {
		if t == toolName {
			return true
		}
	}
	return false
}

// DenyCommands returns the list of denied command patterns from the profile.
func (s *ProfileService) DenyCommands() []string {
	return s.profile.Tools.DenyCommands
}
