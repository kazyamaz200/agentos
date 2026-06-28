package profile

type ProfileService struct {
	profile *Profile
}

func NewProfileService(profile *Profile) *ProfileService {
	return &ProfileService{profile: profile}
}

func (s *ProfileService) Profile() *Profile {
	return s.profile
}

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

func (s *ProfileService) DenyCommands() []string {
	return s.profile.Tools.DenyCommands
}
