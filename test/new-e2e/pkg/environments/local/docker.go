package local

import (
	"fmt"

	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/e2e"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/environments"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/runner"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/utils/optional"
	"github.com/DataDog/test-infra-definitions/resources/local/docker"
	dclocal "github.com/DataDog/test-infra-definitions/scenarios/local/docker"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	provisionerBaseID = "local-docker-"
	defaultVMName     = "vm"
)

// ProvisionerParams is a set of parameters for the Provisioner.
type ProvisionerParams struct {
	name              string
	extraConfigParams runner.ConfigMap
}

func newProvisionerParams() *ProvisionerParams {
	// We use nil arrays to decide if we should create or not
	return &ProvisionerParams{
		name: defaultVMName,
	}
}

// GetProvisionerParams return ProvisionerParams from options opts setup
func GetProvisionerParams(opts ...ProvisionerOption) *ProvisionerParams {
	params := newProvisionerParams()
	err := optional.ApplyOptions(params, opts)
	if err != nil {
		panic(fmt.Errorf("unable to apply ProvisionerOption, err: %w", err))
	}
	return params
}

// ProvisionerOption is a provisioner option.
type ProvisionerOption func(*ProvisionerParams) error

func Provisioner(opts ...ProvisionerOption) e2e.TypedProvisioner[environments.DockerLocal] {
	// We need to build params here to be able to use params.name in the provisioner name
	params := GetProvisionerParams(opts...)

	provisioner := e2e.NewTypedPulumiProvisioner(provisionerBaseID+params.name, func(ctx *pulumi.Context, env *environments.DockerLocal) error {
		// We ALWAYS need to make a deep copy of `params`, as the provisioner can be called multiple times.
		// and it's easy to forget about it, leading to hard to debug issues.
		params := GetProvisionerParams(opts...)
		return Run(ctx, env, params)
	}, params.extraConfigParams)

	return provisioner
}

// Run deploys a environment given a pulumi.Context
func Run(ctx *pulumi.Context, env *environments.DockerLocal, params *ProvisionerParams) (err error) {
	var localEnv docker.Environment
	if env.Local != nil {
		localEnv = *env.Local
	} else {
		localEnv, err = docker.NewEnvironment(ctx)
		if err != nil {
			return err
		}
	}
	host, err := dclocal.NewVM(localEnv, "test")
	if err != nil {
		return err
	}
	err = host.Export(ctx, &env.RemoteHost.HostOutput)
	if err != nil {
		return err
	}
	_ = ctx.Log.Info(fmt.Sprintf("Running test on container '%v'", host.Name()), nil)
	return nil
}
