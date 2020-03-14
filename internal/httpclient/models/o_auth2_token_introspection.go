// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OAuth2TokenIntrospection Introspection contains an access token's session data as specified by IETF RFC 7662, see:
//
// https://tools.ietf.org/html/rfc7662
// swagger:model oAuth2TokenIntrospection
type OAuth2TokenIntrospection struct {

	// Active is a boolean indicator of whether or not the presented token
	// is currently active.  The specifics of a token's "active" state
	// will vary depending on the implementation of the authorization
	// server and the information it keeps about its tokens, but a "true"
	// value return for the "active" property will generally indicate
	// that a given token has been issued by this authorization server,
	// has not been revoked by the resource owner, and is within its
	// given time window of validity (e.g., after its issuance time and
	// before its expiration time).
	// Required: true
	Active *bool `json:"active"`

	// Audience contains a list of the token's intended audiences.
	Aud []string `json:"aud"`

	// ClientID is aclient identifier for the OAuth 2.0 client that
	// requested this token.
	ClientID string `json:"client_id,omitempty"`

	// Expires at is an integer timestamp, measured in the number of seconds
	// since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`

	// Extra is arbitrary data set by the session.
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Issued at is an integer timestamp, measured in the number of seconds
	// since January 1 1970 UTC, indicating when this token was
	// originally issued.
	Iat int64 `json:"iat,omitempty"`

	// IssuerURL is a string representing the issuer of this token
	Iss string `json:"iss,omitempty"`

	// NotBefore is an integer timestamp, measured in the number of seconds
	// since January 1 1970 UTC, indicating when this token is not to be
	// used before.
	Nbf int64 `json:"nbf,omitempty"`

	// ObfuscatedSubject is set when the subject identifier algorithm was set to "pairwise" during authorization.
	// It is the `sub` value of the ID Token that was issued.
	ObfuscatedSubject string `json:"obfuscated_subject,omitempty"`

	// Scope is a JSON string containing a space-separated list of
	// scopes associated with this token.
	Scope string `json:"scope,omitempty"`

	// Subject of the token, as defined in JWT [RFC7519].
	// Usually a machine-readable identifier of the resource owner who
	// authorized this token.
	Sub string `json:"sub,omitempty"`

	// TokenType is the introspected token's type, for example `access_token` or `refresh_token`.
	TokenType string `json:"token_type,omitempty"`

	// Username is a human-readable identifier for the resource owner who
	// authorized this token.
	Username string `json:"username,omitempty"`
}

// Validate validates this o auth2 token introspection
func (m *OAuth2TokenIntrospection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuth2TokenIntrospection) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2TokenIntrospection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2TokenIntrospection) UnmarshalBinary(b []byte) error {
	var res OAuth2TokenIntrospection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
