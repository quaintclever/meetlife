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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

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
	// make install 之后, 获取 cr
	cr := &paasv1.ReplicasDemo{}
	if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
		logger.Error(err, fmt.Sprintf("cr get fail!, crName: %s, namespace: %s", cr.Name, cr.Namespace))
		return ctrl.Result{}, err
	}
	logger.Info("======= Perform Reconcile =======", "name", cr.Name, "namespace", cr.Namespace)

	// 查询 deployment 是否存在, 不存在则创建
	foundDeployment := &corev1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: cr.Name, Namespace: cr.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// 新建 deployment
		deployment := &corev1.Deployment{}
		deployment.APIVersion = cr.APIVersion
		deployment.Kind = "Deployment"
		deployment.ObjectMeta.Name = cr.ObjectMeta.Name
		deployment.ObjectMeta.Namespace = cr.ObjectMeta.Namespace
		deployment.Spec = cr.Spec.DeploymentSpec
		var initReplicas int32 = 0
		deployment.Spec.Replicas = &initReplicas
		// 如果一批次小于总 replicas, 启动一批次的 pod
		// 设置 ownerReferences
		if cr.Spec.BatchSize < *cr.Spec.DeploymentSpec.Replicas {
			*deployment.Spec.Replicas = cr.Spec.BatchSize
		}
		trueVal := true
		falseVal := false
		deployment.OwnerReferences = append(deployment.OwnerReferences, v1.OwnerReference{
			APIVersion:         cr.APIVersion,
			Kind:               cr.Kind,
			Name:               cr.Name,
			UID:                cr.UID,
			Controller:         &trueVal,
			BlockOwnerDeletion: &falseVal,
		})
		logger.Info("======= Creating Deployment =======", "deployment", cr.Name)
		if err := r.Create(ctx, deployment); err != nil {
			return ctrl.Result{}, err
		}

		cr.Status.CurrentBatch = 1
		cr.Status.Ready = fmt.Sprintf("%d/%d", 0, foundDeployment.Spec.Replicas)
		logger.Info("======= Update CR Status 1 =======", "crName:", cr.Name)
		err = r.Update(ctx, cr)
	} else if err == nil {
		// 如果deploy 里 ready 的数量 等于总数量, 根据 cr, 修改 deployment 状态
		if *foundDeployment.Spec.Replicas != *cr.Spec.DeploymentSpec.Replicas &&
			foundDeployment.Status.ReadyReplicas == *foundDeployment.Spec.Replicas {
			// 下一批次 到顶了.
			if *foundDeployment.Spec.Replicas+cr.Spec.BatchSize >= *cr.Spec.DeploymentSpec.Replicas {
				*foundDeployment.Spec.Replicas = *cr.Spec.DeploymentSpec.Replicas
			} else {
				*foundDeployment.Spec.Replicas = *foundDeployment.Spec.Replicas + cr.Spec.BatchSize
			}
			logger.Info("======= Updating Deployment =======", "deployment", cr.Name)
			err = r.Update(ctx, foundDeployment)
			// 修改cr 状态
			cr.Status.CurrentBatch = cr.Status.CurrentBatch + 1
			cr.Status.Ready = fmt.Sprintf("%d/%d", foundDeployment.Status.ReadyReplicas, foundDeployment.Status.Replicas)
			logger.Info("======= Update CR Status 2 =======", "crName:", cr.Name)
			err = r.Update(ctx, cr)
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ReplicasDemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&paasv1.ReplicasDemo{}).
		Complete(r)
}
