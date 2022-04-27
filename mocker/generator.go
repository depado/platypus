package mocker

import "github.com/gin-gonic/gin"

// EndpointGenerator is a simple interface.
type EndpointGenerator interface {
	Generate(string, *gin.Engine)
}

// GetEndpoint implements the EndpointGenerator interface.
type GetEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e GetEndpoint) Generate(path string, r *gin.Engine) { r.GET(path, e.ToHandler()) }

// PostEndpoint implements the EndpointGenerator interface.
type PostEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e PostEndpoint) Generate(path string, r *gin.Engine) { r.POST(path, e.ToHandler()) }

// PutEndpoint implements the EndpointGenerator interface.
type PutEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e PutEndpoint) Generate(path string, r *gin.Engine) { r.PUT(path, e.ToHandler()) }

// PatchEndpoint implements the EndpointGenerator interface.
type PatchEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e PatchEndpoint) Generate(path string, r *gin.Engine) { r.PATCH(path, e.ToHandler()) }

// DeleteEndpoint implements the EndpointGenerator interface.
type DeleteEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e DeleteEndpoint) Generate(path string, r *gin.Engine) { r.DELETE(path, e.ToHandler()) }

// HeadEndpoint implements the EndpointGenerator interface.
type HeadEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint
func (e HeadEndpoint) Generate(path string, r *gin.Engine) { r.HEAD(path, e.ToHandler()) }

// OptionsEndpoint implements the EndpointGenerator interface.
type OptionsEndpoint struct {
	EndpointMethod `yaml:",inline"`
}

// Generate generates the endpoint.
func (e OptionsEndpoint) Generate(path string, r *gin.Engine) { r.OPTIONS(path, e.ToHandler()) }
