// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

// SyncType represents type of k8s received message
type SyncType string

// SyncType values
//nolint
const (
	COMMAND       SyncType = "COMMAND"
	CONFIGMAP     SyncType = "CONFIGMAP"
	ENDPOINTS     SyncType = "ENDPOINTS"
	INGRESS       SyncType = "INGRESS"
	INGRESS_CLASS SyncType = "INGRESS_CLASS"
	NAMESPACE     SyncType = "NAMESPACE"
	SERVICE       SyncType = "SERVICE"
	SECRET        SyncType = "SECRET"
)

// SyncDataEvent represents converted k8s received message
type SyncDataEvent struct {
	_ [0]int
	SyncType
	Namespace string
	Data      interface{}
}
