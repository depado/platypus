package mocker

// Header is a simple header struct used to parse the configuration file
type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Response is a simple response struct used to parse the configuration file
type Response struct {
	Code    int      `yaml:"code"`
	Body    string   `yaml:"body"`
	Headers []Header `yaml:"headers"`
	Preset  string   `yaml:"preset"`
	Ratio   int      `yaml:"ratio"`
}

// Endpoint is a simple endpoint struct used to parse the configuration file
type Endpoint struct {
	Path    string           `yaml:"path"`
	Get     *GetEndpoint     `yaml:"get"`
	Post    *PostEndpoint    `yaml:"post"`
	Put     *PutEndpoint     `yaml:"put"`
	Patch   *PatchEndpoint   `yaml:"patch"`
	Delete  *DeleteEndpoint  `yaml:"delete"`
	Head    *HeadEndpoint    `yaml:"head"`
	Options *OptionsEndpoint `yaml:"options"`

	All []EndpointGenerator `yaml:"-"`
}

// Compute checks every possible endpoint to generate a slice of
// EndpointGenerator
func (e *Endpoint) Compute() {
	e.All = []EndpointGenerator{}
	if e.Get != nil {
		e.All = append(e.All, e.Get)
	}
	if e.Post != nil {
		e.All = append(e.All, e.Post)
	}
	if e.Put != nil {
		e.All = append(e.All, e.Put)
	}
	if e.Patch != nil {
		e.All = append(e.All, e.Patch)
	}
	if e.Delete != nil {
		e.All = append(e.All, e.Delete)
	}
	if e.Head != nil {
		e.All = append(e.All, e.Head)
	}
	if e.Options != nil {
		e.All = append(e.All, e.Options)
	}
}

// All stores all the information of the configuration file
type All struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}
