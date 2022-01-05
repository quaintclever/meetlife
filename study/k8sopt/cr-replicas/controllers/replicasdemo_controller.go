/*
Copyright 2022.

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
	"fmt"

	corev1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	paasv1 "github.com/quaintclever/meetlife/api/v1"
)

// ReplicasDemoReconciler reconciles a ReplicasDemo object
type ReplicasDemoReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=paas.github.com,resources=replicasdemoes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=paas.github.com,resources=replicasdemoes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=paas.github.com,resources=replicasdemoes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ReplicasDemo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *ReplicasDemoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	cr := &paasv1.ReplicasDemo{}
	if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
		logger.Error(err, fmt.Sprintf("cr get fail!, crName: %s, namespace: %s", cr.Name, cr.Namespace))
		return ctrl.Result{}, err
	}
	logger.Info("perform reconcile", "name", cr.Name, "namespace", cr.Namespace)

	// 新建 deployment
	deployment := &corev1.Deployment{}
	deployment.APIVersion = cr.APIVersion
	deployment.Kind = "Deployment"
	deployment.ObjectMeta.Name = cr.ObjectMeta.Name
	deployment.ObjectMeta.Namespace = cr.ObjectMeta.Namespace
	deployment.Spec = cr.Spec.DeploymentSpec
	if err := r.Create(ctx, deployment); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ReplicasDemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&paasv1.ReplicasDemo{}).
		Complete(r)
}
