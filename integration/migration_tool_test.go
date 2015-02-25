package integration_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Migration Tool", func() {
	It("can be built", func() {
		buildExecutable("github.com/pivotal-cf/rabbitmq-cluster-migration-tool")
	})
})
