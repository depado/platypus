package mocker

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

func codeToColor(code int) aurora.Value {
	switch {
	case code >= 100 && code < 200:
		return aurora.Cyan(code)
	case code >= 200 && code < 300:
		return aurora.Green(code)
	case code >= 300 && code < 400:
		return aurora.Yellow(code)
	case code >= 400 && code < 500:
		return aurora.Red(code)
	case code >= 500 && code < 600:
		return aurora.BrightRed(code)
	}
	return aurora.White(code)
}

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

// Info returns a string to print out the information of a response
func (r Response) Info(prefix string, last bool) string {
	var sb strings.Builder
	s := "├─"
	if last {
		s = "└─"
	}
	sb.WriteString(fmt.Sprintf("%s %s %s", prefix, s, codeToColor(r.Code).String()))
	if r.Preset != "" {
		switch r.Preset {
		case "json":
			sb.WriteString(" JSON")
		case "text":
			sb.WriteString(" Text")
		default:
			sb.WriteString(" " + r.Preset)
		}
	}
	if r.Ratio != 0 && r.Ratio != 100 {
		sb.WriteString(fmt.Sprintf(" [%d%%]", r.Ratio))
	}
	return sb.String()
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
	hasNext := true
	withNext := "├─ %s%s\n│\n"
	withoutNext := "└─ %s%s\n"

	fmt.Printf("\n%s\n", aurora.Underline(e.Path))
	if e.Get != nil {
		e.Get.CalcRatios()
		e.All = append(e.All, e.Get)
		hasNext = e.Post != nil || e.Put != nil || e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Blue("GET"), e.Get.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Blue("GET"), e.Get.Info(!hasNext))
		}
	}
	if e.Post != nil {
		e.Post.CalcRatios()
		e.All = append(e.All, e.Post)
		hasNext = e.Put != nil || e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Green("POST"), e.Post.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Green("POST"), e.Post.Info(!hasNext))
		}
	}
	if e.Put != nil {
		e.Put.CalcRatios()
		e.All = append(e.All, e.Put)
		hasNext = e.Patch != nil || e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Yellow("PUT"), e.Put.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Yellow("PUT"), e.Put.Info(!hasNext))
		}
	}
	if e.Patch != nil {
		e.Patch.CalcRatios()
		e.All = append(e.All, e.Patch)
		hasNext = e.Delete != nil || e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.BrightYellow("PATCH"), e.Patch.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.BrightYellow("PATCH"), e.Patch.Info(!hasNext))
		}
	}
	if e.Delete != nil {
		e.Delete.CalcRatios()
		e.All = append(e.All, e.Delete)
		hasNext = e.Head != nil || e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Red("DELETE"), e.Delete.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Red("DELETE"), e.Delete.Info(!hasNext))
		}
	}
	if e.Head != nil {
		e.Head.CalcRatios()
		e.All = append(e.All, e.Head)
		hasNext = e.Options != nil
		if hasNext {
			fmt.Printf(withNext, aurora.Cyan("HEAD"), e.Head.Info(!hasNext))
		} else {
			fmt.Printf(withoutNext, aurora.Cyan("HEAD"), e.Head.Info(!hasNext))
		}
	}
	if e.Options != nil {
		e.Options.CalcRatios()
		e.All = append(e.All, e.Options)
		fmt.Printf(withoutNext, aurora.BrightCyan("OPTIONS"), e.Options.Info(true))
	}
}

// All stores all the information of the configuration file
type All struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}
