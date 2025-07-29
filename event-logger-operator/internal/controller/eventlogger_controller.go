package controller

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

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
		for _, c := range d.Spec.Template.Spec.Containers {
			replicas := int32(0)
			if d.Spec.Replicas != nil {
				replicas = *d.Spec.Replicas
			}
			fmt.Printf("[Deployment Watch] Namespace=%s Name=%s Replicas=%d Container=%s Image=%s\n", d.Namespace, d.Name, replicas, c.Name, c.Image)
		}
	}

	return ctrl.Result{}, nil
}

func (r *EventLoggerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitorv1.EventLogger{}).
		Watches(
			&appsv1.Deployment{},
			handler.EnqueueRequestsFromMapFunc(r.mapDeploymentToEventLogger),
		).
		Complete(r)
}

func (r *EventLoggerReconciler) mapDeploymentToEventLogger(ctx context.Context, obj client.Object) []reconcile.Request {
	var eventLoggers monitorv1.EventLoggerList
	if err := r.List(ctx, &eventLoggers, client.InNamespace(obj.GetNamespace())); err != nil {
		// Handle error, maybe log it
		return nil
	}

	if len(eventLoggers.Items) == 0 {
		// No EventLogger CRs found, no need to reconcile
		return nil
	}

	// Assuming you want to reconcile the first EventLogger found in the namespace
	// You might want to adjust this logic based on your specific requirements
	return []reconcile.Request{
		{
			NamespacedName: types.NamespacedName{
				Name:      eventLoggers.Items[0].Name,
				Namespace: eventLoggers.Items[0].Namespace,
			},
		},
	}
}
