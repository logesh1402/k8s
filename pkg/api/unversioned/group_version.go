/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package unversioned

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying
// concepts during lookup stages without having partially valid types
type GroupKind struct {
	Group string
	Kind  string
}

func (gk GroupKind) WithVersion(version string) GroupVersionKind {
	return GroupVersionKind{Group: gk.Group, Version: version, Kind: gk.Kind}
}

func (gk *GroupKind) String() string {
	return gk.Group + ", Kind=" + gk.Kind
}

// GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion
// to avoid automatic coersion.  It doesn't use a GroupVersion to avoid custom marshalling
type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

// TODO remove this
func NewGroupVersionKind(gv GroupVersion, kind string) GroupVersionKind {
	return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}

func (gvk GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}

func (gvk GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvk.Group, Version: gvk.Version}
}

func (gvk *GroupVersionKind) String() string {
	return gvk.Group + "/" + gvk.Version + ", Kind=" + gvk.Kind
}

// GroupVersion contains the "group" and the "version", which uniquely identifies the API.
type GroupVersion struct {
	Group   string
	Version string
}

// IsEmpty returns true if group and version are empty
func (gv GroupVersion) IsEmpty() bool {
	return len(gv.Group) == 0 && len(gv.Version) == 0
}

// String puts "group" and "version" into a single "group/version" string. For the legacy v1
// it returns "v1".
func (gv GroupVersion) String() string {
	// special case the internal apiVersion for the legacy kube types
	if gv.IsEmpty() {
		return ""
	}

	// special case of "v1" for backward compatibility
	if len(gv.Group) == 0 && gv.Version == "v1" {
		return gv.Version
	}
	if len(gv.Group) > 0 {
		return gv.Group + "/" + gv.Version
	}
	return gv.Version
}

// ParseGroupVersion turns "group/version" string into a GroupVersion struct. It reports error
// if it cannot parse the string.
func ParseGroupVersion(gv string) (GroupVersion, error) {
	// this can be the internal version for the legacy kube types
	// TODO once we've cleared the last uses as strings, this special case should be removed.
	if (len(gv) == 0) || (gv == "/") {
		return GroupVersion{}, nil
	}

	s := strings.Split(gv, "/")
	// "v1" is the only special case. Otherwise GroupVersion is expected to contain
	// one "/" dividing the string into two parts.
	switch {
	case len(s) == 1 && gv == "v1":
		return GroupVersion{"", "v1"}, nil
	case len(s) == 1:
		return GroupVersion{"", s[0]}, nil
	case len(s) == 2:
		return GroupVersion{s[0], s[1]}, nil
	default:
		return GroupVersion{}, fmt.Errorf("Unexpected GroupVersion string: %v", gv)
	}
}

func ParseGroupVersionOrDie(gv string) GroupVersion {
	ret, err := ParseGroupVersion(gv)
	if err != nil {
		panic(err)
	}

	return ret
}

// WithKind creates a GroupVersionKind based on the method receiver's GroupVersion and the passed Kind.
func (gv GroupVersion) WithKind(kind string) GroupVersionKind {
	return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}

// MarshalJSON implements the json.Marshaller interface.
func (gv GroupVersion) MarshalJSON() ([]byte, error) {
	s := gv.String()
	if strings.Count(s, "/") > 1 {
		return []byte{}, fmt.Errorf("illegal GroupVersion %v: contains more than one /", s)
	}
	return json.Marshal(s)
}

func (gv *GroupVersion) unmarshal(value []byte) error {
	var s string
	if err := json.Unmarshal(value, &s); err != nil {
		return err
	}
	parsed, err := ParseGroupVersion(s)
	if err != nil {
		return err
	}
	*gv = parsed
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (gv *GroupVersion) UnmarshalJSON(value []byte) error {
	return gv.unmarshal(value)
}

// UnmarshalTEXT implements the Ugorji's encoding.TextUnmarshaler interface.
func (gv *GroupVersion) UnmarshalText(value []byte) error {
	return gv.unmarshal(value)
}

// UpdateTypeMeta is a helper for setting a TypeMeta based on a provided gvk.
func UpdateTypeMeta(meta *TypeMeta, gvk *GroupVersionKind) {
	if gvk == nil {
		meta.APIVersion, meta.Kind = "", ""
		return
	}
	meta.APIVersion, meta.Kind = gvk.GroupVersion().String(), gvk.Kind
}

// TypeMetaToGroupVersionKind attempts to return a GVK for the provided TypeMeta.
// If the APIVersion cannot be parsed under the rules for ParseGroupVersion, neither
// group or version are set in the returned object.
func TypeMetaToGroupVersionKind(meta TypeMeta) *GroupVersionKind {
	if gv, err := ParseGroupVersion(meta.APIVersion); err == nil {
		return &GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: meta.Kind}
	}
	return &GroupVersionKind{Kind: meta.Kind}
}

// ToAPIVersionAndKind is a convenience method for satisfying runtime.Object on types that
// do not use TypeMeta.
func (gvk *GroupVersionKind) ToAPIVersionAndKind() (string, string) {
	if gvk == nil {
		return "", ""
	}
	return gvk.GroupVersion().String(), gvk.Kind
}

// FromAPIVersionAndKind returns a GVK representing the provided fields for types that
// do not use TypeMeta.
func FromAPIVersionAndKind(apiVersion, kind string) *GroupVersionKind {
	if gv, err := ParseGroupVersion(apiVersion); err == nil {
		return &GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
	}
	return &GroupVersionKind{Kind: kind}
}
