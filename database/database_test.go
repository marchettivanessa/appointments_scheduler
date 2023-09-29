package database

import (
	"appointments_scheduler/config"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Suite")
}

var _ = Describe("Conversions Tracker Database", func() {
	Context("build connection string", func() {
		It("builds the data source name according to config", func() {
			c := config.DatabaseConfig{
				Name:           "any_database_name",
				Host:           "any_host",
				Password:       "any_password",
				Username:       "any_username",
				Port:           5432,
				ConnectTimeout: 20,
				Schema:         "conversions_tracker",
			}

			exercise := buildConnectionString(c)
			expectedString := "host=any_host port=5432 user=any_username password=any_password dbname=any_database_name connect_timeout=20 search_path=public,conversions_tracker sslmode=disable"
			Expect(exercise).To(Equal(expectedString))
		})
	})
})
