package golang

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"istio.io/tools/protoc-gen-crds/pkg/crd"
	"istio.io/tools/protoc-gen-crds/pkg/crd/openapi"
	"istio.io/tools/protoc-gen-crds/pkg/naming"
)

const definitionTmpl = `//
// GENERATED CODE -- DO NOT EDIT
//
package {{.PackageName}}

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"istio.io/tools/protoc-gen-crds/testdata"	// TODO: generate this properly
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type {{.StructName}} struct {
	metav1.TypeMeta ` + "`json:\",inline\"`" + `

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta ` + "`json:\"metadata,omitempty\" protobuf:\"bytes,1,opt,name=metadata\"`" + `

	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec {{.SpecName}} ` + "`json:\"spec,omitempty\" protobuf:\"bytes,2,opt,name=spec\"`" + `
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type {{.ListName}} struct {
	metav1.TypeMeta ` + "`json:\",inline\"`" + `

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta ` + "`json:\"metadata,omitempty\" protobuf:\"bytes,1,opt,name=metadata\"`" + `

	// Items is the list of Ingress.
	Items []{{.StructName}} ` + "`json:\"items\" protobuf:\"bytes,2,rep,name=items\"`" + `
}
{{ range $oid, $type := .ObjectTypes }}

// TODO: where do we get the name of the original type, e.g. testdata.BasicMessage?
type {{$type.Name}} testdata.BasicMessage

// TODO: everything that follows is really only needed for the top level type

func (m *{{$type.Name}}) Reset()         { *m = {{$type.Name}}{} }
func (m *{{$type.Name}}) String() string { return proto.CompactTextString(m) }
func (*{{$type.Name}}) ProtoMessage()    {}

// Implement proto.Message
func (spec {{$type.Name}}) MarshalJSON() ([]byte, error) {
	return marshalJSONProtobufHelper(&spec)
}

// TODO: do we need to actually implement this? Guessing it won't get called
func (spec {{$type.Name}}) UnmarshalJSON(b []byte) error {
	return nil
}

{{end}}

func marshalJSONProtobufHelper(pb proto.Message) ([]byte, error) {
	j := &bytes.Buffer{}
	marshaler := &jsonpb.Marshaler{}
	err := marshaler.Marshal(j, pb)
	if err != nil {
		return nil, err
	}
	return j.Bytes(), nil
}
`

var definitionTemplate = template.
	Must(template.New("template").Parse(definitionTmpl))

type definitionContext struct {
	PackageName string
	StructName  string
	ListName    string
	SpecName    string

	ObjectTypes []*openapi.Object
}

// Emit go code for the given Custom Resource Definition.
func Emit(m *crd.ResourceDefinition) (string, error) {
	contents := ""

	var c definitionContext
	c.PackageName = m.Spec.Version
	c.StructName = naming.PascalCase(m.Spec.Names.Kind)
	c.SpecName = naming.PascalCase(m.Spec.Names.Kind) + "Spec"
	c.ListName = naming.PascalCase(m.Spec.Names.Kind + "List")

	// Get object types in order
	types := make(map[*openapi.Object]struct{})
	addTypes(types, m.Spec.Validation)

	var ordered []*openapi.Object
	for t := range types {
		ordered = append(ordered, t)
	}
	sort.Slice(ordered, func(i, j int) bool {
		ti := ordered[i]
		tj := ordered[j]
		if ti == m.Spec.Validation {
			return true
		}
		if tj == m.Spec.Validation {
			return false
		}
		return strings.Compare(ti.Name, tj.Name) < 0
	})
	c.ObjectTypes = ordered

	var b bytes.Buffer
	if err := definitionTemplate.Execute(&b, c); err != nil {
		return "", err
	}
	contents += string(b.Bytes())

	return contents, nil
}

func addTypes(types map[*openapi.Object]struct{}, t openapi.Type) {
	switch ty := t.(type) {
	case *openapi.Object:
		types[ty] = struct{}{}

		for _, f := range ty.Fields {
			addTypes(types, f.Type)
		}

	case *openapi.Array:
		addTypes(types, ty.ElementType)
	}
}
