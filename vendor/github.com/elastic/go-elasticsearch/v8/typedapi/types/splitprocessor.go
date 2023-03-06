// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// SplitProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/ingest/_types/Processors.ts#L354-L360
type SplitProcessor struct {
	Description      *string              `json:"description,omitempty"`
	Field            string               `json:"field"`
	If               *string              `json:"if,omitempty"`
	IgnoreFailure    *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing    *bool                `json:"ignore_missing,omitempty"`
	OnFailure        []ProcessorContainer `json:"on_failure,omitempty"`
	PreserveTrailing *bool                `json:"preserve_trailing,omitempty"`
	Separator        string               `json:"separator"`
	Tag              *string              `json:"tag,omitempty"`
	TargetField      *string              `json:"target_field,omitempty"`
}

// NewSplitProcessor returns a SplitProcessor.
func NewSplitProcessor() *SplitProcessor {
	r := &SplitProcessor{}

	return r
}
