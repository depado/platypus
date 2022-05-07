package mocker

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// Endpoint represents a single endpoint.
type Endpoint struct {
	Path string `yaml:"path"`

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
// EndpointGenerator.
func (e *Endpoint) Compute() {
	e.All = []EndpointGenerator{}
	var hasNext bool
	withNext := "├─ %s%s\n│\n"
	withoutNext := "└─ %s%s\n"

	fmt.Printf("\n%s\n", aurora.Underline(e.Path))
	if e.Get != nil {
		e.All = append(e.All, e.Get)
		hasNext = e.Post != nil || e.Put != nil || e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Blue("GET"), e.Get.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Blue("GET"), e.Get.Info(!hasNext))
		}
	}
	if e.Post != nil {
		e.All = append(e.All, e.Post)
		hasNext = e.Put != nil || e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Green("POST"), e.Post.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Green("POST"), e.Post.Info(!hasNext))
		}
	}
	if e.Put != nil {
		e.All = append(e.All, e.Put)
		hasNext = e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Yellow("PUT"), e.Put.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Yellow("PUT"), e.Put.Info(!hasNext))
		}
	}
	if e.Patch != nil {
		e.All = append(e.All, e.Patch)
		hasNext = e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.BrightYellow("PATCH"), e.Patch.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.BrightYellow("PATCH"), e.Patch.Info(!hasNext))
		}
	}
	if e.Delete != nil {
		e.All = append(e.All, e.Delete)
		hasNext = e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Red("DELETE"), e.Delete.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Red("DELETE"), e.Delete.Info(!hasNext))
		}
	}
	if e.Head != nil {
		e.All = append(e.All, e.Head)
		hasNext = e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Cyan("HEAD"), e.Head.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Cyan("HEAD"), e.Head.Info(!hasNext))
		}
	}
	if e.Options != nil {
		e.All = append(e.All, e.Options)
		fmt.Printf(withoutNext, aurora.BrightCyan("OPTIONS"), e.Options.Info(true))
	}
}
