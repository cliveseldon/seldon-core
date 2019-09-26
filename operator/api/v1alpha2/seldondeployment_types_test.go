/*
Copyright 2019 The Seldon Authors.

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

package v1alpha2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetSeldonDeploymentName(t *testing.T) {
	scenarios := map[string]struct {
		mlDep SeldonDeployment
		want  string
	}{
		"empty spec.Name and objectMeta.name": {
			mlDep: SeldonDeployment{},
			want:  "-",
		},
		"empty spec.Name": {
			mlDep: SeldonDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "example",
				},
			},
			want: "-example",
		},
	}

	for name, scenario := range scenarios {
		t.Run(name, func(t *testing.T) {
			got := GetSeldonDeploymentName(&scenario.mlDep)
			if diff := cmp.Diff(scenario.want, got); diff != "" {
				t.Errorf("GetSeldonDeploymentName() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
