/*
Copyright 2021.

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

package controllers

import (
	"context"
	"strings"

	gamev1 "github.com/fosuoy/tig-operator/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// TigReconciler reconciles a Tig object
type TigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=game.example.com,resources=tigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=game.example.com,resources=tigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=game.example.com,resources=tigs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=get;watch;list

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Tig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *TigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	serviceAccount := &v1.ServiceAccount{}

	err := r.Get(ctx, req.NamespacedName, serviceAccount)
	if err != nil {
		return ctrl.Result{}, nil
	}

	serviceAccountName := serviceAccount.ObjectMeta.Name

	if strings.ToLower(serviceAccountName) != "it" {
		return ctrl.Result{}, nil
	}

	tig := &gamev1.Tig{}
	tig.Name = "tig"
	tig.Namespace = serviceAccount.Namespace
	tig.Spec.It = serviceAccount.Namespace

	err = r.Get(
		ctx,
		client.ObjectKey{Namespace: serviceAccount.Namespace, Name: "tig"},
		tig)
	if err != nil {
		log.Info(serviceAccount.Namespace + " is it - tigging...")
		err = r.Create(ctx, tig)
		if err != nil {
			log.Error(err, "Failed to create tig")
		}
		return ctrl.Result{}, nil
	}

	err = r.Update(ctx, tig)
	if err != nil {
		log.Error(err, "Failed to update tig")
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.ServiceAccount{}).
		Complete(r)
}
