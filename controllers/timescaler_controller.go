package controllers

import (
    "context"
    "time"

    "github.com/robfig/cron/v3"
    appsv1 "k8s.io/api/apps/v1"
    "k8s.io/apimachinery/pkg/types"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"    
    "sigs.k8s.io/controller-runtime/pkg/log"

    autoscalingv1 "github.com/anurag-2911/go-operator-sample/api/v1"
)
func (r *TimeScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    _log := log.FromContext(ctx)

    // Fetch the TimeScaler instance
    var timeScaler autoscalingv1.TimeScaler
    if err := r.Get(ctx, req.NamespacedName, &timeScaler); err != nil {
        _log.Error(err, "Unable to fetch TimeScaler")
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Parse the schedule
    scheduler, err := cron.ParseStandard(timeScaler.Spec.Schedule)
    if err != nil {
        _log.Error(err, "Unable to parse schedule")
        return ctrl.Result{}, err
    }

    // Calculate next scheduled time and requeue if it's not yet time
    now := time.Now()
    nextRun := scheduler.Next(now)
    if nextRun.After(now) {
        return ctrl.Result{RequeueAfter: time.Until(nextRun)}, nil
    }

    // Assuming the name of the Deployment is stored in the TimeScaler's annotations for simplicity
    deploymentName, exists := timeScaler.Annotations["deploymentName"]
    if !exists {
        _log.Error(err, "Deployment name not specified in TimeScaler annotations")
        return ctrl.Result{}, err
    }

    var deployment appsv1.Deployment
    if err := r.Get(ctx, types.NamespacedName{Name: deploymentName, Namespace: req.Namespace}, &deployment); err != nil {
        _log.Error(err, "Unable to fetch Deployment")
        return ctrl.Result{}, err
    }

    // Scale the Deployment
    deployment.Spec.Replicas = &timeScaler.Spec.Replicas
    if err := r.Update(ctx, &deployment); err != nil {
        _log.Error(err, "Failed to scale Deployment")
        return ctrl.Result{}, err
    }

    // Optionally update the TimeScaler status or emit an event here

    return ctrl.Result{RequeueAfter: time.Until(nextRun)}, nil
}
