package controllers

/*

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
// +kubebuilder:docs-gen:collapse=Apache License

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:docs-gen:collapse=Imports

var _ = Describe("Tig controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		ServiceAccountName = "test-cronjob"
		TigNamespace       = "testing-namespace"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating an 'it' service account", func() {
		It("Should create a Tig object", func() {
			By("By creating a new CronJob")
			ctx := context.Background()

			testingNamespace := &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: TigNamespace,
				},
			}
			Expect(k8sClient.Create(ctx, testingNamespace)).Should(Succeed())

			serviceAccount := &v1.ServiceAccount{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "v1",
					Kind:       "ServiceAccount",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ServiceAccountName,
					Namespace: TigNamespace,
				},
			}
			Expect(k8sClient.Create(ctx, serviceAccount)).Should(Succeed())
		})
	})

})
