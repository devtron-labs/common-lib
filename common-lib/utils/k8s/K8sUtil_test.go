/*
 * Copyright (c) 2020-2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package k8s

import (
	"github.com/devtron-labs/common-lib/utils"
	"testing"
)

var k8sUtilClient *K8sServiceImpl
var clusterConfig *ClusterConfig

func init() {
	config := &RuntimeConfig{LocalDevMode: true}
	logger, _ := utils.NewSugardLogger()
	k8sUtilClient = NewK8sUtil(logger, config)
	clusterConfig = &ClusterConfig{
		Host:        "",
		BearerToken: "",
	}
}

func TestK8sUtil_checkIfNsExists(t *testing.T) {
	tests := []struct {
		name       string
		namespace  string
		wantExists bool
		wantErr    bool
	}{
		{
			name:       "test-kube-system",
			namespace:  "kube-system",
			wantErr:    false,
			wantExists: true,
		}, {
			name:       "test-randum",
			namespace:  "test-rand-laknd-kwejdwiu",
			wantExists: false,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.SkipNow()
		t.Run(tt.name, func(t *testing.T) {
			impl := k8sUtilClient
			k8s, _ := impl.GetCoreV1Client(clusterConfig)
			_, gotExists, err := impl.GetNsIfExists(tt.namespace, k8s)
			if (err != nil) != tt.wantErr {
				t.Errorf("K8sServiceImpl.checkIfNsExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExists != tt.wantExists {
				t.Errorf("K8sServiceImpl.checkIfNsExists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestK8sUtil_CreateNsIfNotExists(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		wantErr   bool
	}{
		{
			name:      "create test",
			namespace: "createtestns",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.SkipNow()
		t.Run(tt.name, func(t *testing.T) {
			impl := k8sUtilClient
			if _, _, err := impl.CreateNsIfNotExists(tt.namespace, clusterConfig); (err != nil) != tt.wantErr {
				t.Errorf("K8sServiceImpl.CreateNsIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
			k8s, _ := impl.GetCoreV1Client(clusterConfig)
			if err := impl.deleteNs(tt.namespace, k8s); (err != nil) != tt.wantErr {
				t.Errorf("K8sServiceImpl.deleteNs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
