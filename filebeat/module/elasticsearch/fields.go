// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
