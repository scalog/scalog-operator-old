package scalogservice

import (
	"context"
	"fmt"
	"strconv"

	scalogv1alpha1 "github.com/scalog/scalog-operator/pkg/apis/scalog/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_scalogservice")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ScalogService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileScalogService{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("scalogservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ScalogService
	err = c.Watch(&source.Kind{Type: &scalogv1alpha1.ScalogService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ScalogService
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &scalogv1alpha1.ScalogService{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileScalogService{}

// ReconcileScalogService reconciles a ScalogService object
type ReconcileScalogService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ScalogService object and makes changes based on the state read
// and what is in the ScalogService.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileScalogService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ScalogService")

	// Fetch the ScalogService instance
	instance := &scalogv1alpha1.ScalogService{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("instance not found")
			return reconcile.Result{}, nil
		}
		reqLogger.Info("error getting the instance")
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Attach a service account if it does not yet exist
	serviceAccount := corev1.ServiceAccount{}
	err2 := r.client.Get(context.TODO(), types.NamespacedName{Namespace: "scalog", Name: "scalog-data-service-account"}, &serviceAccount)
	if err2 != nil {
		if errors.IsNotFound(err2) {
			reqLogger.Info("Service Account resource not found. Creating...")
			sa := newServiceAccount()
			if err := r.client.Create(context.TODO(), sa); err != nil {
				reqLogger.Info("Something went wrong while creating the service account")
				return reconcile.Result{}, err
			}
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong with reading service account")
	}

	dataService := corev1.Service{}
	err = r.client.Get(context.Background(), types.NamespacedName{Namespace: "scalog", Name: "scalog-data-headless-service"}, &dataService)
	if err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Data Service not found. Creating...")
			service := newDataService()
			if err := r.client.Create(context.Background(), service); err != nil {
				reqLogger.Info("Something went wrong while creating the data service")
				return reconcile.Result{}, err
			}
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong with reading data service")
	}

	// Reconcile the number of data shards
	existingDataShards := &appsv1.StatefulSetList{}
	dataShardSelector := client.ListOptions{}
	dataShardSelector.SetLabelSelector(fmt.Sprintf("app=%s", "scalog-data"))
	err = r.client.List(context.TODO(), &dataShardSelector, existingDataShards)
	if err == nil {
		currSize := len(existingDataShards.Items)
		if instance.Spec.NumShards == currSize {
			reqLogger.Info(fmt.Sprintf("god saw that there were %d shards running and it was good", currSize))
		} else if instance.Spec.NumShards > currSize {
			reqLogger.Info(fmt.Sprintf("Not enough shards. Current: %d. Desired: %d", currSize, instance.Spec.NumShards))

			// Updating the latest shardID
			instance.Status.LatestShardID++
			if err := r.client.Update(context.TODO(), instance); err != nil {
				reqLogger.Error(err, "Failed to update ScalogService status")
				return reconcile.Result{}, err
			}

			// With the service now properly created, we can attempt to create a statefulset to live under that service
			replicas := newDataStatefulSet(strconv.Itoa(instance.Status.LatestShardID))
			if err := r.client.Create(context.TODO(), replicas); err != nil {
				reqLogger.Error(err, fmt.Sprintf("Failed to create statefulset for shard %d", instance.Status.LatestShardID))
				return reconcile.Result{}, err
			}
		} else { // We have too many shards
			// TODO: Randomly finalize one and then kill
			reqLogger.Info(fmt.Sprintf("Too many shards. Current: %d. Desired: %d", currSize, instance.Spec.NumShards))
		}
	}

	return reconcile.Result{}, nil
}
