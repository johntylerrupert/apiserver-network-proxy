/*
Copyright 2022 The Kubernetes Authors.

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

package tests

import (
	"testing"
)

func TestReadiness(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)

	proxy, server, cleanup, err := runGRPCProxyServerWithServerCount(1)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ready, _ := server.Readiness.Ready()
	if ready {
		t.Fatalf("expected not ready")
	}

	clientset := runAgent(proxy.agent, stopCh)
	waitForConnectedServerCount(t, 1, clientset)

	ready, _ = server.Readiness.Ready()
	if !ready {
		t.Fatalf("expected ready")
	}
}
