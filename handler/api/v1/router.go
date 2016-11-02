package v1

import (
	"github.com/aws/amazon-ecs-event-stream-handler/handler/regex"
	"github.com/gorilla/mux"
)

// TODO: add a map of path and query keys and use the map in task apis instead of hardcoding strings
var (
	// Stripping off '^' and '$' from the beginning and end of regexes respectively for the router
	clusterNameRegex = string(regex.ClusterNameRegex[1 : len(regex.ClusterNameRegex)-1])
	clusterARNRegex  = string(regex.ClusterARNRegex[1 : len(regex.ClusterARNRegex)-1])
	taskARNRegex     = string(regex.TaskARNRegex[1 : len(regex.TaskARNRegex)-1])
	instanceARNRegex = string(regex.InstanceARNRegex[1 : len(regex.InstanceARNRegex)-1])

	getTaskPath     = "/tasks/{cluster:" + clusterNameRegex + "}/{arn:" + taskARNRegex + "}"
	listTasksPath   = "/tasks"
	filterTasksPath = "/tasks/filter"
	streamTasksPath = "/tasks/stream"

	getInstancePath     = "/instances/{cluster:" + clusterNameRegex + "}/{arn:" + instanceARNRegex + "}"
	listInstancesPath   = "/instances"
	filterInstancesPath = "/instances/filter"
	streamInstancesPath = "/instances/stream"

	clusterKey     = "cluster"
	clusterNameVal = "{" + clusterKey + ":" + clusterNameRegex + "}"
	clusterARNVal  = "{" + clusterKey + ":" + clusterARNRegex + "}"

	taskKey    = "task"
	taskARNVal = "{" + taskKey + ":" + taskARNRegex + "}"

	instanceKey    = "instance"
	instanceARNVal = "{" + instanceKey + ":" + instanceARNRegex + "}"

	statusKey         = "status"
	taskStatusVal     = "{" + statusKey + ":pending|running|stopped}"
	instanceStatusVal = "{" + statusKey + ":active|inactive}"
)

func NewRouter(apis APIs) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	s := r.Path("/v1").Subrouter()

	// Tasks

	// Get task using cluster name and task ARN
	s.Path(getTaskPath).
		Methods("GET").
		HandlerFunc(apis.TaskApis.GetTask)

	// List tasks
	s.Path(listTasksPath).
		Methods("GET").
		HandlerFunc(apis.TaskApis.ListTasks)

	// Filter tasks by status
	s.Path(filterTasksPath).
		Queries(statusKey, taskStatusVal).
		Methods("GET").
		HandlerFunc(apis.TaskApis.FilterTasks)

	// Stream tasks
	s.Path(streamTasksPath).
		Methods("GET").
		HandlerFunc(apis.TaskApis.StreamTasks)

	// Instances

	// Get instance using cluster name and instance ARN
	s.Path(getInstancePath).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.GetInstance)

	// List instances
	s.Path(listInstancesPath).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.ListInstances)

	// Filter instances by status
	s.Path(filterInstancesPath).
		Queries(statusKey, instanceStatusVal).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.FilterInstances)

	// Filter instances by cluser name
	s.Path(filterInstancesPath).
		Queries(clusterKey, clusterNameVal).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.FilterInstances)

	// Filter instances by cluster ARN
	s.Path(filterInstancesPath).
		Queries(clusterKey, clusterARNVal).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.FilterInstances)

	// Stream instances
	s.Path(streamInstancesPath).
		Methods("GET").
		HandlerFunc(apis.ContainerInstanceApis.StreamInstances)

	return s
}