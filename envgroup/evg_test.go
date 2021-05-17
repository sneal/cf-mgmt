package envgroup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmwarepivotallabs/cf-mgmt/config"
	cfgfakes "github.com/vmwarepivotallabs/cf-mgmt/config/fakes"
	"github.com/vmwarepivotallabs/cf-mgmt/envgroup"
	"github.com/vmwarepivotallabs/cf-mgmt/envgroup/fakes"
)

// These are directly from go-cfclient and don't really need to be tested

var _ = Describe("EnvironmentVariableGroups", func() {
	Context("running", func() {
		var cf *fakes.FakeCFClient

		BeforeEach(func() {
			cf = new(fakes.FakeCFClient)
		})

		It("Gets the group", func() {
			expected := map[string]interface{}{
				"foo": "bar",
			}
			cf.GetRunningEnvironmentVariableGroupReturns(expected, nil)

			manager := envgroup.NewManager(cf, nil)

			actual, err := manager.GetEnvironmentVariableGroup(true)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(actual).Should(Equal(expected))
		})

		It("Sets the Group", func() {
			expected := map[string]interface{}{
				"bar": "baz",
			}
			cfg := new(cfgfakes.FakeManager)
			cfg.GetGlobalConfigReturns(&config.GlobalConfig{
				RunningEnvironmentVariableGroup: expected,
			}, nil)

			cf.SetRunningEnvironmentVariableGroupReturns(nil)

			manager := envgroup.NewManager(cf, cfg)
			err := manager.SetEnvironmentVariableGroup(true)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
