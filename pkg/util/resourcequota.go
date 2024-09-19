/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	corev1 "k8s.io/api/core/v1"
)

// SumResourceList sums the resource list.
func SumResourceList(lists []corev1.ResourceList) corev1.ResourceList {
	total := corev1.ResourceList{}
	for _, list := range lists {
		for name, quantity := range list {
			if value, ok := total[name]; !ok {
				total[name] = quantity.DeepCopy()
			} else {
				value.Add(quantity)
				total[name] = value
			}
		}
	}
	return total
}