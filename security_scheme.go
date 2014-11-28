package swagger

import (
	"encoding/json"

	"github.com/casualjim/go-swagger/reflection"
)

const (
	basic       = "basic"
	apiKey      = "apiKey"
	oauth2      = "oauth2"
	implicit    = "implicit"
	password    = "password"
	application = "application"
	accessCode  = "accessCode"
)

// SecurityRequirement is a single security requirement
type SecurityRequirement map[string][]string

// SecurityRequirements contains all the supported security requirements
type SecurityRequirements []SecurityRequirement

// BasicAuth creates a basic auth security scheme
func BasicAuth() *SecurityScheme {
	return &SecurityScheme{Type: basic}
}

// ApiKeyAuth creates an api key auth security scheme
func ApiKeyAuth(fieldName, valueSource string) *SecurityScheme {
	return &SecurityScheme{Type: apiKey, Name: fieldName, In: valueSource}
}

// OAuth2Implicit creates an implicit flow oauth2 security scheme
func OAuth2Implicit(authorizationURL string) *SecurityScheme {
	return &SecurityScheme{
		Type:             oauth2,
		Flow:             implicit,
		AuthorizationURL: authorizationURL,
	}
}

// OAuth2Password creates a password flow oauth2 security scheme
func OAuth2Password(tokenURL string) *SecurityScheme {
	return &SecurityScheme{
		Type:     oauth2,
		Flow:     password,
		TokenURL: tokenURL,
	}
}

// OAuth2Application creates an application flow oauth2 security scheme
func OAuth2Application(tokenURL string) *SecurityScheme {
	return &SecurityScheme{
		Type:     oauth2,
		Flow:     application,
		TokenURL: tokenURL,
	}
}

// OAuth2AccessToken creates an access token flow oauth2 security scheme
func OAuth2AccessToken(authorizationURL, tokenURL string) *SecurityScheme {
	return &SecurityScheme{
		Type:             oauth2,
		Flow:             accessCode,
		AuthorizationURL: authorizationURL,
		TokenURL:         tokenURL,
	}
}

// SecurityScheme allows the definition of a security scheme that can be used by the operations.
// Supported schemes are basic authentication, an API key (either as a header or as a query parameter)
// and OAuth2's common flows (implicit, password, application and access code).
//
// For more information: http://goo.gl/8us55a#securitySchemeObject
type SecurityScheme struct {
	Description      string                 `swagger:"description,omitempty"`
	Extensions       map[string]interface{} `swagger:"-"` // custom extensions, omitted when empty
	Type             string                 `swagger:"type"`
	Name             string                 `swagger:"name,omitempty"`             // api key
	In               string                 `swagger:"in,omitempty"`               // api key
	Flow             string                 `swagger:"flow,omitempty"`             // oauth2
	AuthorizationURL string                 `swagger:"authorizationUrl,omitempty"` // oauth2
	TokenURL         string                 `swagger:"tokenUrl,omitempty"`         // oauth2
	Scopes           map[string]string      `swagger:"scopes,omitempty"`           // oauth2
}

func (s *SecurityScheme) AddScope(scope, description string) {
	if s.Scopes == nil {
		s.Scopes = make(map[string]string)
	}
	s.Scopes[scope] = description
}

func (s SecurityScheme) MarshalMap() map[string]interface{} {
	res := reflection.MarshalMapRecursed(s)
	addExtensions(res, s.Extensions)
	return res
}

func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.MarshalMap())
}

func (s SecurityScheme) MarshalYAML() (interface{}, error) {
	return s.MarshalMap(), nil
}
