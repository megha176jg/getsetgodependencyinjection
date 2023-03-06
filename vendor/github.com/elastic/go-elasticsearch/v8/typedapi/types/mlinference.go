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

// MlInference type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/xpack/usage/types.ts#L190-L195
type MlInference struct {
	Deployments      *MlInferenceDeployments               `json:"deployments,omitempty"`
	IngestProcessors map[string]MlInferenceIngestProcessor `json:"ingest_processors"`
	TrainedModels    MlInferenceTrainedModels              `json:"trained_models"`
}

// NewMlInference returns a MlInference.
func NewMlInference() *MlInference {
	r := &MlInference{
		IngestProcessors: make(map[string]MlInferenceIngestProcessor, 0),
	}

	return r
}
