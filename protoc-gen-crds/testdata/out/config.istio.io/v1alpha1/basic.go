//
// GENERATED CODE -- DO NOT EDIT
//
package v1alpha1

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

type Basic struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec BasicSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BasicList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of Ingress.
	Items []Basic `json:"items" protobuf:"bytes,2,rep,name=items"`
}


// TODO: where do we get the name of the original type, e.g. testdata.BasicMessage?
type BasicSpec testdata.BasicMessage

// TODO: everything that follows is really only needed for the top level type

func (m *BasicSpec) Reset()         { *m = BasicSpec{} }
func (m *BasicSpec) String() string { return proto.CompactTextString(m) }
func (*BasicSpec) ProtoMessage()    {}

// Implement proto.Message
func (spec BasicSpec) MarshalJSON() ([]byte, error) {
	return marshalJSONProtobufHelper(&spec)
}

// TODO: do we need to actually implement this? Guessing it won't get called
func (spec BasicSpec) UnmarshalJSON(b []byte) error {
	return nil
}



// TODO: where do we get the name of the original type, e.g. testdata.BasicMessage?
type InnerMessage testdata.BasicMessage

// TODO: everything that follows is really only needed for the top level type

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}

// Implement proto.Message
func (spec InnerMessage) MarshalJSON() ([]byte, error) {
	return marshalJSONProtobufHelper(&spec)
}

// TODO: do we need to actually implement this? Guessing it won't get called
func (spec InnerMessage) UnmarshalJSON(b []byte) error {
	return nil
}



func marshalJSONProtobufHelper(pb proto.Message) ([]byte, error) {
	j := &bytes.Buffer{}
	marshaler := &jsonpb.Marshaler{}
	err := marshaler.Marshal(j, pb)
	if err != nil {
		return nil, err
	}
	return j.Bytes(), nil
}
