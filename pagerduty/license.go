package pagerduty

import (
	"context"
)

// LicenseService handles the communication with field related methods of the PagerDuty API.
type LicenseService service

// License represents a license.
type License struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	CurrentValue         int      `json:"current_value,omitempty"`
	AllocationsAvailable int      `json:"allocations_available,omitempty"`
	ValidRoles           []string `json:"valid_roles"`
	RoleGroup            string   `json:"role_group,omitempty"`
	Summary              string   `json:"summary,omitempty"`
	Type                 string   `json:"type,omitempty"`
	Self                 *string  `json:"self"`
	HTMLURL              *string  `json:"html_url"`
}

// ListLicenseResponse represents a list response of licenses
type ListLicenseResponse struct {
	Licenses []*License `json:"licenses,omitempty"`
}

// List lists existing licenses. If a non-zero Limit is passed as an option, only a single page of results will be
// returned. Otherwise, the entire list of licenses will be returned.
func (s *LicenseService) List() (*ListLicenseResponse, *Response, error) {
	return s.ListContext(context.Background())
}

// ListContext lists existing licenses. If a non-zero Limit is passed as an option, only a single page of results will be
// returned. Otherwise, the entire list of licenses will be returned.
func (s *LicenseService) ListContext(ctx context.Context) (*ListLicenseResponse, *Response, error) {
	u := "/licenses"
	v := new(ListLicenseResponse)

	resp, err := s.client.newRequestDoContext(ctx, "GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
