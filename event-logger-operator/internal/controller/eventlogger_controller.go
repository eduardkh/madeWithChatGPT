package controller

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"

	monitorv1 "github.com/you/event-logger-operator/api/v1"
)

// EventLoggerReconciler reconciles a EventLogger object
type EventLoggerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *EventLoggerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Fetch the EventLogger CR
	var el monitorv1.EventLogger
	if err := r.Get(ctx, req.NamespacedName, &el); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// List Deployments in the same namespace as the CR
	var deployments appsv1.DeploymentList
	if err := r.List(ctx, &deployments, client.InNamespace(req.Namespace)); err != nil {
		return ctrl.Result{}, err
	}

	for _, d := range deployments.Items {
		replicas := int32(0)
		if d.Spec.Replicas != nil {
			replicas = *d.Spec.Replicas
		}
		fmt.Printf("[Deployment Watch] Namespace=%s Name=%s Replicas=%d\n", d.Namespace, d.Name, replicas)
	}

	return ctrl.Result{}, nil
}

func (r *EventLoggerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitorv1.EventLogger{}).
		Watches(
			&appsv1.Deployment{},
			&handler.EnqueueRequestForObject{},
		).
		Complete(r)
}
