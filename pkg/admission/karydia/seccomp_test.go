// Copyright (C) 2019 SAP SE or an SAP affiliate company. All rights reserved.
// This file is licensed under the Apache Software License, v. 2 except as
// noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package karydia

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

/* Mutating and Validating Webhook
 * Removes token mounts of the default service account when automountServiceToken is undefined.
 * kubectl annotate ns default karydia.gardener.cloud/seccompProfile=runtime/default
 */
func TestPodSeccompDefaultProfileWithAnnotation(t *testing.T) {
	pod := corev1.Pod{}
	var patches patchOperations
	var validationErrors []string

	pod.Annotations = make(map[string]string)
	pod.Annotations["seccomp.security.alpha.kubernetes.io/pod"] = "runtime/default"

	patches = mutatePodSeccompProfile(pod, "runtime/default", patches)
	if len(patches) != 0 {
		t.Errorf("expected 0 patches but got: %+v", patches)
	}
	mutatedPod, err := patchPod(pod, patches)
	if err != nil {
		t.Errorf("failed to apply patches: %+v", err)
	}
	// Zero validation errors expected for mutated pod
	validationErrors = validatePodSeccompProfile(mutatedPod, "runtime/default", validationErrors)
	if len(validationErrors) != 0 {
		t.Errorf("expected 0 validationErrors but got: %+v", validationErrors)
	}
	validationErrors = []string{}
	// Zero validation error expected for initial pod
	validationErrors = validatePodSeccompProfile(pod, "runtime/default", validationErrors)
	if len(validationErrors) != 0 {
		t.Errorf("expected 0 validationErrors but got: %+v", validationErrors)
	}
}

func TestPodSeccompDefaultProfileNoAnnotation(t *testing.T) {
	pod := corev1.Pod{}
	var patches patchOperations
	var validationErrors []string

	patches = mutatePodSeccompProfile(pod, "runtime/default", patches)
	if len(patches) != 1 {
		t.Errorf("expected 1 patches but got: %+v", patches)
	}
	mutatedPod, err := patchPod(pod, patches)
	if err != nil {
		t.Errorf("failed to apply patches: %+v", err)
	}
	// Zero validation errors expected for mutated pod
	validationErrors = validatePodSeccompProfile(mutatedPod, "runtime/default", validationErrors)
	if len(validationErrors) != 0 {
		t.Errorf("expected 0 validationErrors but got: %+v", validationErrors)
	}
	validationErrors = []string{}
	// One validation error expected for initial pod
	validationErrors = validatePodSeccompProfile(pod, "runtime/default", validationErrors)
	if len(validationErrors) != 1 {
		t.Errorf("expected 1 validationErrors but got: %+v", validationErrors)
	}
}

func TestPodSeccompDefaultProfileOtherAnnotation(t *testing.T) {
	pod := corev1.Pod{}
	var patches patchOperations
	var validationErrors []string

	pod.Annotations = make(map[string]string)
	pod.Annotations["seccomp.security.alpha.kubernetes.io/pod"] = "runtime/other"

	patches = mutatePodSeccompProfile(pod, "runtime/default", patches)
	if len(patches) != 0 {
		t.Errorf("expected 0 patches but got: %+v", patches)
	}
	mutatedPod, err := patchPod(pod, patches)
	if err != nil {
		t.Errorf("failed to apply patches: %+v", err)
	}
	// One, validation error expected for mutated pod
	validationErrors = validatePodSeccompProfile(mutatedPod, "runtime/default", validationErrors)
	if len(validationErrors) != 1 {
		t.Errorf("expected 1 validationErrors but got: %+v", validationErrors)
	}
	validationErrors = []string{}
	// One validation error expected for initial pod
	validationErrors = validatePodSeccompProfile(pod, "runtime/default", validationErrors)
	if len(validationErrors) != 1 {
		t.Errorf("expected 1 validationErrors but got: %+v", validationErrors)
	}
}
