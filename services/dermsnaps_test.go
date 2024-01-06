package services_test

import (
	"dermsnap/models"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DermsnapService", func() {
	var user *models.User
	BeforeAll(func() {
		user, err := Services.UserService.CreateUser("dermsnapUser1", models.Admin, models.Google)
		Expect(err).To(BeNil())
		Expect(user.Identifier).To(Equal("dermsnapUser1"))
	})
	Describe("CreateDermsnap", func() {
		It("should create a dermsnap", func() {
			dermsnap, err := Services.DermsnapService.CreateDermsnap(user.ID, models.CreateDermsnap{
				StartTime:      time.Now().AddDate(0, 0, -10),
				Duration:       10,
				Locations:      []models.BodyLocation{"arms"},
				Changed:        false,
				NewMedications: []string{},
				Itchy:          true,
				Painful:        false,
				MoreInfo:       "",
			})

			Expect(err).To(BeNil())
			Expect(dermsnap.ID).ToNot(BeNil())
		})
	})

	Describe("Get User Dermsnaps", func() {
		BeforeAll(func() {
			_, err := Services.DermsnapService.CreateDermsnap(user.ID, models.CreateDermsnap{
				StartTime:      time.Now().AddDate(0, 0, -10),
				Duration:       10,
				Locations:      []models.BodyLocation{"arms"},
				Changed:        false,
				NewMedications: []string{},
				Itchy:          true,
				Painful:        false,
				MoreInfo:       "",
			})
			Expect(err).To(BeNil())
			_, err = Services.DermsnapService.CreateDermsnap(user.ID, models.CreateDermsnap{
				StartTime:      time.Now().AddDate(0, 0, -10),
				Duration:       10,
				Locations:      []models.BodyLocation{"arms"},
				Changed:        false,
				NewMedications: []string{},
				Itchy:          true,
				Painful:        false,
				MoreInfo:       "",
			})
			Expect(err).To(BeNil())
		})
		It("should get dermsnaps for a user by user id", func() {
			dermsnaps, err := Services.DermsnapService.GetUserDermsnaps(user.ID)
			Expect(err).To(BeNil())
			Expect(len(dermsnaps)).To(Equal(2))
			Expect(dermsnaps[len(dermsnaps)-1].Itchy).To(Equal(true))
		})
	})
})
