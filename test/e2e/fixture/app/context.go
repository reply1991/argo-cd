package app

import (
	"testing"
	"time"

	. "github.com/argoproj/argo-cd/common"
	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/test/e2e/fixture"
	"github.com/argoproj/argo-cd/test/e2e/fixture/certs"
	"github.com/argoproj/argo-cd/test/e2e/fixture/repos"
	"github.com/argoproj/argo-cd/util/settings"
)

// this implements the "given" part of given/when/then
type Context struct {
	t           *testing.T
	path        string
	chart       string
	repoURLType fixture.RepoURLType
	// seconds
	timeout                int
	name                   string
	destServer             string
	env                    string
	parameters             []string
	jsonnetTLAStr          []string
	jsonnetTLACode         []string
	namePrefix             string
	nameSuffix             string
	resource               string
	prune                  bool
	configManagementPlugin string
	async                  bool
	localPath              string
	project                string
	revision               string
	force                  bool
}

func Given(t *testing.T) *Context {
	fixture.EnsureCleanState(t)
	return &Context{t: t, destServer: KubernetesInternalAPIServerAddr, repoURLType: fixture.RepoURLTypeFile, name: fixture.Name(), timeout: 10, project: "default", prune: true}
}

func (c *Context) CustomCACertAdded() *Context {
	certs.AddCustomCACert()
	return c
}

func (c *Context) CustomSSHKnownHostsAdded() *Context {
	certs.AddCustomSSHKnownHostsKeys()
	return c
}

func (c *Context) HTTPSRepoURLAdded(withCreds bool) *Context {
	repos.AddHTTPSRepo(false, withCreds)
	return c
}

func (c *Context) HTTPSInsecureRepoURLAdded(withCreds bool) *Context {
	repos.AddHTTPSRepo(true, withCreds)
	return c
}

func (c *Context) HTTPSInsecureRepoURLWithClientCertAdded() *Context {
	repos.AddHTTPSRepoClientCert(true)
	return c
}

func (c *Context) HTTPSRepoURLWithClientCertAdded() *Context {
	repos.AddHTTPSRepoClientCert(false)
	return c
}

func (c *Context) SSHRepoURLAdded(withCreds bool) *Context {
	repos.AddSSHRepo(false, withCreds)
	return c
}

func (c *Context) SSHInsecureRepoURLAdded(withCreds bool) *Context {
	repos.AddSSHRepo(true, withCreds)
	return c
}

func (c *Context) HelmRepoAdded(name string) *Context {
	repos.AddHelmRepo(name)
	return c
}

func (c *Context) HTTPSCredentialsUserPassAdded() *Context {
	repos.AddHTTPSCredentialsUserPass()
	return c
}

func (c *Context) HTTPSCredentialsTLSClientCertAdded() *Context {
	repos.AddHTTPSCredentialsTLSClientCert()
	return c
}

func (c *Context) SSHCredentialsAdded() *Context {
	repos.AddSSHCredentials()
	return c
}

func (c *Context) ProjectSpec(spec v1alpha1.AppProjectSpec) *Context {
	fixture.SetProjectSpec(c.project, spec)
	return c
}

func (c *Context) RepoURLType(urlType fixture.RepoURLType) *Context {
	c.repoURLType = urlType
	return c
}

func (c *Context) Name(name string) *Context {
	c.name = name
	return c
}

func (c *Context) Path(path string) *Context {
	c.path = path
	return c
}

func (c *Context) Chart(chart string) *Context {
	c.chart = chart
	return c
}

func (c *Context) Revision(revision string) *Context {
	c.revision = revision
	return c
}

func (c *Context) Timeout(timeout int) *Context {
	c.timeout = timeout
	return c
}

func (c *Context) DestServer(destServer string) *Context {
	c.destServer = destServer
	return c
}

func (c *Context) Env(env string) *Context {
	c.env = env
	return c
}

func (c *Context) Parameter(parameter string) *Context {
	c.parameters = append(c.parameters, parameter)
	return c
}

func (c *Context) JsonnetTLAStrParameter(parameter string) *Context {
	c.jsonnetTLAStr = append(c.jsonnetTLAStr, parameter)
	return c
}

func (c *Context) JsonnetTLACodeParameter(parameter string) *Context {
	c.jsonnetTLACode = append(c.jsonnetTLACode, parameter)
	return c
}

// group:kind:name
func (c *Context) SelectedResource(resource string) *Context {
	c.resource = resource
	return c
}

func (c *Context) NamePrefix(namePrefix string) *Context {
	c.namePrefix = namePrefix
	return c
}

func (c *Context) NameSuffix(nameSuffix string) *Context {
	c.nameSuffix = nameSuffix
	return c
}

func (c *Context) ResourceOverrides(overrides map[string]v1alpha1.ResourceOverride) *Context {
	fixture.SetResourceOverrides(overrides)
	return c
}

func (c *Context) ResourceFilter(filter settings.ResourcesFilter) *Context {
	fixture.SetResourceFilter(filter)
	return c
}

// this both configures the plugin, but forces use of it
func (c *Context) ConfigManagementPlugin(plugin v1alpha1.ConfigManagementPlugin) *Context {
	fixture.SetConfigManagementPlugins(plugin)
	c.configManagementPlugin = plugin.Name
	return c
}

func (c *Context) And(block func()) *Context {
	block()
	return c
}

func (c *Context) When() *Actions {
	// in case any settings have changed, pause for 1s, not great, but fine
	time.Sleep(1 * time.Second)
	return &Actions{context: c}
}

func (c *Context) Prune(prune bool) *Context {
	c.prune = prune
	return c
}

func (c *Context) Async(async bool) *Context {
	c.async = async
	return c
}

func (c *Context) LocalPath(localPath string) *Context {
	c.localPath = localPath
	return c
}

func (c *Context) Project(project string) *Context {
	c.project = project
	return c
}

func (c *Context) Force() *Context {
	c.force = true
	return c
}