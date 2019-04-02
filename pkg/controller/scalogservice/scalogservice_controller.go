package scalogservice

import (
	"context"
	"fmt"
	"reflect"
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
	if err := r.client.Get(context.TODO(), request.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("scalog service instance not found")
			return reconcile.Result{}, nil
		}
		reqLogger.Info("error getting the instance")
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Attach a service account if it does not yet exist
	serviceAccount := corev1.ServiceAccount{}
	if err := r.client.Get(context.Background(), types.NamespacedName{Namespace: "scalog", Name: "scalog-service-account"}, &serviceAccount); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Service Account resource not found. Creating...")
			sa := newServiceAccount()
			if saErr := r.client.Create(context.Background(), sa); saErr != nil {
				reqLogger.Info("Something went wrong while creating the service account")
				return reconcile.Result{}, saErr
			}
			// Successfully created the service account. requeue to serve further requests
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong with reading service account")
		return reconcile.Result{}, err
	}

	// Create a order service if it does not exist
	orderService := corev1.Service{}
	if err := r.client.Get(context.Background(), types.NamespacedName{Namespace: "scalog", Name: "scalog-order-service"}, &orderService); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Order service not found. Creating...")
			service := newOrderService()
			if osErr := r.client.Create(context.Background(), service); osErr != nil {
				reqLogger.Info("Something went wrong while creating the order service")
				return reconcile.Result{}, osErr
			}
			// Successfully created the order service. requeue to serve further requests
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong while reading the order service")
		return reconcile.Result{}, err
	}

	// Create a order deployment if it doesn't exist
	orderDeploy := &appsv1.Deployment{}
	if err := r.client.Get(context.Background(), types.NamespacedName{Namespace: "scalog", Name: "scalog-order-deployment"}, orderDeploy); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Order deployment not found. Creating...")
			deploy := newOrderDeployment(int32(instance.Spec.NumMetadataReplica))
			if deployErr := r.client.Create(context.Background(), deploy); deployErr != nil {
				reqLogger.Info("Something went wrong while creating the order deployment")
				return reconcile.Result{}, deployErr
			}
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong while fetching the order deployment")
		return reconcile.Result{}, err
	}

	// Reconcile the number of ordering layer nodes
	orderReplicaSpecSize := int32(instance.Spec.NumMetadataReplica)
	if *orderDeploy.Spec.Replicas != orderReplicaSpecSize {
		orderDeploy.Spec.Replicas = &orderReplicaSpecSize
		if err := r.client.Update(context.Background(), orderDeploy); err != nil {
			reqLogger.Error(err, "Failed to update Deployment", "Deployment.Namespace", orderDeploy.Namespace, "Deployment.Name", orderDeploy.Name)
			return reconcile.Result{}, err
		}
		// Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
	}

	// Create a data service to contain all of the data layer stateful set
	dataService := corev1.Service{}
	if err := r.client.Get(context.Background(), types.NamespacedName{Namespace: "scalog", Name: "scalog-headless-data-service"}, &dataService); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("Data Service not found. Creating...")
			service := newDataService()
			if sErr := r.client.Create(context.Background(), service); sErr != nil {
				reqLogger.Info("Something went wrong while creating the data service")
				return reconcile.Result{}, sErr
			}
			return reconcile.Result{Requeue: true}, nil
		}
		reqLogger.Info("Something went wrong with reading data service")
		return reconcile.Result{}, err
	}

	// Reconcile the number of data shards
	existingDataShards := &appsv1.StatefulSetList{}
	dataShardSelector := client.ListOptions{}
	dataShardSelector.SetLabelSelector(fmt.Sprintf("role=scalog-data-shard"))
	dataShardSelector.InNamespace("scalog")
	if err := r.client.List(context.Background(), &dataShardSelector, existingDataShards); err == nil {
		currSize := len(existingDataShards.Items)
		if instance.Spec.NumShards == currSize {
			reqLogger.Info(fmt.Sprintf("god saw that there were %d shards running and it was good", currSize))
		} else if instance.Spec.NumShards > currSize {
			reqLogger.Info(fmt.Sprintf("Not enough shards. Current: %d. Desired: %d", currSize, instance.Spec.NumShards))

			// Updating the latest shardID
			instance.Status.LatestShardID++
			if instanceErr := r.client.Update(context.Background(), instance); instanceErr != nil {
				reqLogger.Error(err, "Failed to update ScalogService instance")
				return reconcile.Result{}, instanceErr
			}

			// With the service now properly created, we can attempt to create a statefulset to live under that service
			shard := newDataStatefulSet(strconv.Itoa(instance.Status.LatestShardID), int32(instance.Spec.NumDataReplica))
			if rErr := r.client.Create(context.Background(), shard); rErr != nil {
				reqLogger.Info(fmt.Sprintf("Failed to create statefulset for shard %d", instance.Status.LatestShardID))
				return reconcile.Result{}, rErr
			}
			// Successfully created shard
			return reconcile.Result{Requeue: true}, nil
		} else { // We have too many shards
			// TODO: Randomly finalize one and then kill
			reqLogger.Info(fmt.Sprintf("Too many shards. Current: %d. Desired: %d", currSize, instance.Spec.NumShards))
		}
	}

	// Ensure that each data replica maintains its own service
	existingDataReplicas := &corev1.PodList{}
	externalDataReplicaSelector := client.ListOptions{}
	externalDataReplicaSelector.SetLabelSelector("role=scalog-data-replica")
	externalDataReplicaSelector.InNamespace("scalog")
	if err := r.client.List(context.Background(), &externalDataReplicaSelector, existingDataReplicas); err == nil {
		// Ensure that each data replica maintains its own service
		externalDataService := &corev1.ServiceList{}
		externalDataServiceSelector := client.ListOptions{}
		externalDataServiceSelector.SetLabelSelector("role=scalog-exposed-data-service")
		externalDataServiceSelector.InNamespace("scalog")
		if dssErr := r.client.List(context.Background(), &externalDataServiceSelector, externalDataService); dssErr != nil {
			return reconcile.Result{}, dssErr
		}
		// Convert existing services into an easy to query map
		externalDataServiceMap := map[string]corev1.Service{}
		for _, service := range externalDataService.Items {
			externalDataServiceMap[service.Name] = service
		}

		for _, pod := range existingDataReplicas.Items {
			// Automatically written by k8. We need to search for a corresponding service
			serviceName := constructExternalDataServiceName(pod.Name)
			if _, ok := externalDataServiceMap[serviceName]; !ok {
				// We do not current have a service. We should create one
				dss := newDataServerService(pod.Name)
				if esErr := r.client.Create(context.Background(), dss); esErr != nil {
					reqLogger.Info("Failed to create external service")
					return reconcile.Result{}, esErr
				}
				return reconcile.Result{Requeue: true}, nil
			}
		}
	} else {
		// An error occured
		reqLogger.Info("Failed to retrieve existing replicas")
		return reconcile.Result{}, err
	}

	// Update Status
	potentialUpdate := instance.Status.DeepCopy()
	potentialUpdate.Phase = "Running"
	potentialUpdate.NumShards = len(existingDataShards.Items)
	potentialUpdate.NumMetadataReplica = int(*orderDeploy.Spec.Replicas)
	if !reflect.DeepEqual(potentialUpdate, instance.Status) {
		potentialUpdate.DeepCopyInto(&instance.Status)
		err := r.client.Update(context.Background(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update Scalog status")
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}
