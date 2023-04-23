package json

import jsoniter "github.com/json-iterator/go"

// RawMessage is exported by component-base/pkg/json package.
type RawMessage = jsoniter.RawMessage

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by component-base/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by component-base/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by component-base/pkgn/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by component-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by component-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
