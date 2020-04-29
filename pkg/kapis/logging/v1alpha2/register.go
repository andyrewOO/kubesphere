/*

 Copyright 2019 The KubeSphere Authors.

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
package v1alpha2

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/api/logging/v1alpha2"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/simple/client/k8s"
	"kubesphere.io/kubesphere/pkg/simple/client/logging"
	"net/http"
)

const (
	GroupName = "logging.kubesphere.io"
	RespOK    = "ok"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

func AddToContainer(c *restful.Container, k8sClient k8s.Client, loggingClient logging.Interface) error {
	ws := runtime.NewWebService(GroupVersion)

	h := newHandler(k8sClient, loggingClient)

	ws.Route(ws.GET("/cluster").
		To(h.handleClusterQuery).
		Doc("Query logs against the cluster.").
		Param(ws.QueryParameter("operation", "Operation type. This can be one of four types: query (for querying logs), statistics (for retrieving statistical data), histogram (for displaying log count by time interval) and export (for exporting logs). Defaults to query.").DefaultValue("query").DataType("string").Required(false)).
		Param(ws.QueryParameter("workspaces", "A comma-separated list of workspaces. This field restricts the query to specified workspaces. For example, the following filter matches the workspace my-ws and demo-ws: `my-ws,demo-ws`").DataType("string").Required(false)).
		Param(ws.QueryParameter("workspace_query", "A comma-separated list of keywords. Differing from **workspaces**, this field performs fuzzy matching on workspaces. For example, the following value limits the query to workspaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("namespaces", "A comma-separated list of namespaces. This field restricts the query to specified namespaces. For example, the following filter matches the namespace my-ns and demo-ns: `my-ns,demo-ns`").DataType("string").Required(false)).
		Param(ws.QueryParameter("namespace_query", "A comma-separated list of keywords. Differing from **namespaces**, this field performs fuzzy matching on namespaces. For example, the following value limits the query to namespaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("workloads", "A comma-separated list of workloads. This field restricts the query to specified workloads. For example, the following filter matches the workload my-wl and demo-wl: `my-wl,demo-wl`").DataType("string").Required(false)).
		Param(ws.QueryParameter("workload_query", "A comma-separated list of keywords. Differing from **workloads**, this field performs fuzzy matching on workloads. For example, the following value limits the query to workloads whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("pods", "A comma-separated list of pods. This field restricts the query to specified pods. For example, the following filter matches the pod my-po and demo-po: `my-po,demo-po`").DataType("string").Required(false)).
		Param(ws.QueryParameter("pod_query", "A comma-separated list of keywords. Differing from **pods**, this field performs fuzzy matching on pods. For example, the following value limits the query to pods whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("containers", "A comma-separated list of containers. This field restricts the query to specified containers. For example, the following filter matches the container my-cont and demo-cont: `my-cont,demo-cont`").DataType("string").Required(false)).
		Param(ws.QueryParameter("container_query", "A comma-separated list of keywords. Differing from **containers**, this field performs fuzzy matching on containers. For example, the following value limits the query to containers whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("log_query", "A comma-separated list of keywords. The query returns logs which contain at least one keyword. Case-insensitive matching. For example, if the field is set to `err,INFO`, the query returns any log containing err(ERR,Err,...) *OR* INFO(info,InFo,...).").DataType("string").Required(false)).
		Param(ws.QueryParameter("interval", "Time interval. It requires **operation** is set to histogram. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min).").DefaultValue("15m").DataType("string").Required(false)).
		Param(ws.QueryParameter("start_time", "Start time of query. Default to 0. The format is a string representing milliseconds since the epoch, eg. 1559664000000.").DataType("string").Required(false)).
		Param(ws.QueryParameter("end_time", "End time of query. Default to now. The format is a string representing milliseconds since the epoch, eg. 1559664000000.").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "Sort order. One of acs, desc. This field sorts logs by timestamp.").DataType("string").DefaultValue("desc").Required(false)).
		Param(ws.QueryParameter("from", "The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to query. Defaults to 0 (i.e. from the beginning of the result set).").DataType("integer").DefaultValue("0").Required(false)).
		Param(ws.QueryParameter("size", "Size of result to return. It requires **operation** is set to query. Defaults to 10 (i.e. 10 log records).").DataType("integer").DefaultValue("10").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.LogQueryTag}).
		Writes(v1alpha2.APIResponse{}).
		Returns(http.StatusOK, RespOK, v1alpha2.APIResponse{})).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, "text/plain")

	ws.Route(ws.GET("/namespaces/{namespace}/pods/{pod}/containers/{container}").
		To(h.handleContainerQuery).
		Doc("Query logs against the specific container.").
		Param(ws.PathParameter("namespace", "The name of the namespace.").DataType("string").Required(true)).
		Param(ws.PathParameter("pod", "Pod name.").DataType("string").Required(true)).
		Param(ws.PathParameter("container", "Container name.").DataType("string").Required(true)).
		Param(ws.QueryParameter("operation", "Operation type. This can be one of four types: query (for querying logs), statistics (for retrieving statistical data), histogram (for displaying log count by time interval) and export (for exporting logs). Defaults to query.").DefaultValue("query").DataType("string").Required(false)).
		Param(ws.QueryParameter("log_query", "A comma-separated list of keywords. The query returns logs which contain at least one keyword. Case-insensitive matching. For example, if the field is set to `err,INFO`, the query returns any log containing err(ERR,Err,...) *OR* INFO(info,InFo,...).").DataType("string").Required(false)).
		Param(ws.QueryParameter("interval", "Time interval. It requires **operation** is set to histogram. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min).").DefaultValue("15m").DataType("string").Required(false)).
		Param(ws.QueryParameter("start_time", "Start time of query. Default to 0. The format is a string representing milliseconds since the epoch, eg. 1559664000000.").DataType("string").Required(false)).
		Param(ws.QueryParameter("end_time", "End time of query. Default to now. The format is a string representing milliseconds since the epoch, eg. 1559664000000.").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "Sort order. One of acs, desc. This field sorts logs by timestamp.").DataType("string").DefaultValue("desc").Required(false)).
		Param(ws.QueryParameter("from", "The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to query. Defaults to 0 (i.e. from the beginning of the result set).").DataType("integer").DefaultValue("0").Required(false)).
		Param(ws.QueryParameter("size", "Size of result to return. It requires **operation** is set to query. Defaults to 10 (i.e. 10 log records).").DataType("integer").DefaultValue("10").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.LogQueryTag}).
		Writes(v1alpha2.APIResponse{}).
		Returns(http.StatusOK, RespOK, v1alpha2.APIResponse{})).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_OCTET)

	c.Add(ws)
	return nil
}
