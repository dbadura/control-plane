package external_services_mock_integration

import (
	"testing"

	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
	"github.com/stretchr/testify/assert"
)

func assertSpecInPackageNotNil(t *testing.T, pkg graphql.PackageExt) {
	assert.True(t, len(pkg.APIDefinitions.Data) > 0)
	assert.NotNil(t, pkg.APIDefinitions.Data[0])
	assert.NotNil(t, pkg.APIDefinitions.Data[0].Spec)
	assert.NotNil(t, pkg.APIDefinitions.Data[0].Spec.Data)
}
