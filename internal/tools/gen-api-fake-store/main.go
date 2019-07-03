package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/fake/zz_store.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-fake-store: ")
}

func main() {
	schema.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-fake-store'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

{{ range . }} 
func (s *store) get{{.TypeName}}(zone string) []*sacloud.{{.TypeName}} {
	values := s.get(Resource{{.TypeName}}, zone)
	var ret []*sacloud.{{.TypeName}}
	for _ , v := range values {
		if v, ok := v.(*sacloud.{{.TypeName}}); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) get{{.TypeName}}ByID(zone string, id types.ID) *sacloud.{{.TypeName}} {
	v := s.getByID(Resource{{.TypeName}}, zone, id)
	if v, ok := v.(*sacloud.{{.TypeName}}); ok {
		return v
	}
	return nil
}

func (s *store) set{{.TypeName}}(zone string, value *sacloud.{{.TypeName}}) {
	s.set(Resource{{.TypeName}}, zone, value)
}
{{ end }}
`
