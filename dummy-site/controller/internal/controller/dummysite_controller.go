/*
Copyright 2024.

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

package controller

import (
	"context"
	"reflect"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	stabledwkv1 "stable.dwk/api/v1"
)

// DummySiteReconciler reconciles a DummySite object
type DummySiteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func generateDeploymentName(dummysite *stabledwkv1.DummySite) string {
	return dummysite.Name + "-dep"
}

func deploymentsEqual(existing *apps.Deployment, desired *apps.Deployment) bool {
	if existing.Spec.Template.Spec.Containers[0].Image != desired.Spec.Template.Spec.Containers[0].Image {
		return false
	}
	if !reflect.DeepEqual(existing.Spec.Template.Spec.Containers[0].Env, desired.Spec.Template.Spec.Containers[0].Env) {
		return false
	}

	return true
}

func constructDeployment(dummysite *stabledwkv1.DummySite) (*apps.Deployment, error) {
	deploymentName := generateDeploymentName(dummysite)
	return &apps.Deployment{
		ObjectMeta: meta.ObjectMeta{
			Name:      deploymentName,
			Namespace: dummysite.Namespace,
			Labels: map[string]string{
				"app": dummysite.Name,
			},
		},
		Spec: apps.DeploymentSpec{
			Selector: &meta.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: core.PodTemplateSpec{
				ObjectMeta: meta.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},
				Spec: core.PodSpec{
					Containers: []core.Container{
						{
							Name:  "nginx",
							Image: "3nd3r1/website-copyer:latest",
							Env: []core.EnvVar{
								{
									Name:  "WEBSITE_URL",
									Value: dummysite.Spec.WebsiteUrl,
								},
							},
							Ports: []core.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}, nil
}

// +kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *DummySiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var dummysite stabledwkv1.DummySite
	if err := r.Get(ctx, req.NamespacedName, &dummysite); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var deployments apps.DeploymentList
	if err := r.List(ctx, &deployments, client.InNamespace(dummysite.Namespace), client.MatchingLabels{"app": dummysite.Name}); err != nil {
		log.Log.Error(err, "unable to list Deployments")
		return ctrl.Result{}, err
	}

	if len(deployments.Items) == 0 {
		newDeployment, err := constructDeployment(&dummysite)
		if err != nil {
			log.Log.Error(err, "unable to construct Deployment")
			return ctrl.Result{}, err
		}

		if err := controllerutil.SetControllerReference(&dummysite, newDeployment, r.Scheme); err != nil {
			log.Log.Error(err, "unable to set controller reference")
			return ctrl.Result{}, err
		}

		if err := r.Create(ctx, newDeployment); err != nil {
			log.Log.Error(err, "unable to create Deployment")
			return ctrl.Result{}, err
		}

		log.Log.Info("created Deployment", "deployment", newDeployment.Name)
		return ctrl.Result{}, nil
	}

	if len(deployments.Items) > 1 {
		for i := 1; i < len(deployments.Items); i++ {
			if err := r.Delete(ctx, &deployments.Items[i]); err != nil {
				log.Log.Error(err, "unable to delete Deployment")
				return ctrl.Result{}, err
			}
			log.Log.Info("deleted Deployment", "deployment", deployments.Items[i].Name)
		}
	}

	desiredDeployment, err := constructDeployment(&dummysite)
	if err != nil {
		log.Log.Error(err, "unable to construct Deployment")
		return ctrl.Result{}, err
	}
	if !deploymentsEqual(&deployments.Items[0], desiredDeployment) {
		if err := controllerutil.SetControllerReference(&dummysite, desiredDeployment, r.Scheme); err != nil {
			log.Log.Error(err, "unable to set controller reference")
			return ctrl.Result{}, err
		}

		if err := r.Update(ctx, desiredDeployment); err != nil {
			log.Log.Error(err, "unable to update Deployment")
			return ctrl.Result{}, err
		}
		log.Log.Info("updated Deployment", "deployment", desiredDeployment.Name)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DummySiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stabledwkv1.DummySite{}).
		Owns(&apps.Deployment{}).
		Complete(r)
}
