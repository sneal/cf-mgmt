package ldap_test

import (
	l "github.com/go-ldap/ldap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmwarepivotallabs/cf-mgmt/ldap"
	"github.com/vmwarepivotallabs/cf-mgmt/ldap/fakes"
)

var _ = Describe("ConnectionAdapter", func() {
	fakeConnection := &fakes.FakeConnection{}
	connectionAdapter := ldap.ConnectionAdapter{
		Connection: fakeConnection,
	}
	Describe("Search", func() {
		Context("Connection is closing", func() {
			fakeConnection.IsClosingReturns(true)
			It("Creates a new connection", func() {
				connectionAdapter.Search(&l.SearchRequest{})
				Expect(connectionAdapter.Connection).ShouldNot(Equal(fakeConnection))
			})
		})
	})
})
