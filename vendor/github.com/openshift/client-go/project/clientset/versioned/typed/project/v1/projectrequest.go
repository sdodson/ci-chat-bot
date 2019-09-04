package v1

import (
	v1 "github.com/openshift/api/project/v1"
	rest "k8s.io/client-go/rest"
)

// ProjectRequestsGetter has a method to return a ProjectRequestInterface.
// A group's client should implement this interface.
type ProjectRequestsGetter interface {
	ProjectRequests() ProjectRequestInterface
}

// ProjectRequestInterface has methods to work with ProjectRequest resources.
type ProjectRequestInterface interface {
	Create(*v1.ProjectRequest) (*v1.Project, error)

	ProjectRequestExpansion
}

// projectRequests implements ProjectRequestInterface
type projectRequests struct {
	client rest.Interface
}

// newProjectRequests returns a ProjectRequests
func newProjectRequests(c *ProjectV1Client) *projectRequests {
	return &projectRequests{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a projectRequest and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *projectRequests) Create(projectRequest *v1.ProjectRequest) (result *v1.Project, err error) {
	result = &v1.Project{}
	err = c.client.Post().
		Resource("projectrequests").
		Body(projectRequest).
		Do().
		Into(result)
	return
}