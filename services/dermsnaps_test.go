package services_test

import (
	"dermsnap/models"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DermsnapService", Ordered, func() {
	var dermsnapUser *models.User
	BeforeAll(func() {
		dermsnapUser, _ = UserService.CreateUser("dermsnapUser1", models.Admin, models.Google)
		Expect(dermsnapUser).ToNot(BeNil())
		Expect(dermsnapUser.Identifier).To(Equal("dermsnapUser1"))
	})
	Describe("CreateDermsnap", func() {
		It("should create a dermsnap", func() {
			dermsnap, err := DermsnapService.CreateDermsnap(dermsnapUser.ID, models.CreateDermsnap{
				StartTime: time.Now().AddDate(0, 0, -10),
				Duration:  10,
				Locations: []models.BodyLocation{
					models.Arms,
					models.Abdomen,
				},
				Changed:        false,
				NewMedications: []string{"finasteride"},
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
			_, err := DermsnapService.CreateDermsnap(dermsnapUser.ID, models.CreateDermsnap{
				StartTime: time.Now().AddDate(0, 0, -10),
				Duration:  10,
				Locations: []models.BodyLocation{
					models.Arms,
				},
				Changed:        false,
				NewMedications: []string{},
				Itchy:          true,
				Painful:        false,
				MoreInfo:       "",
			})
			Expect(err).To(BeNil())
			_, err = DermsnapService.CreateDermsnap(dermsnapUser.ID, models.CreateDermsnap{
				StartTime: time.Now().AddDate(0, 0, -10),
				Duration:  10,
				Locations: []models.BodyLocation{
					models.Arms,
				},
				Changed:        false,
				NewMedications: []string{},
				Itchy:          true,
				Painful:        false,
				MoreInfo:       "",
			})
			Expect(err).To(BeNil())
		})
		It("should get dermsnaps for a user by user id", func() {
			dermsnaps, err := DermsnapService.GetUserDermsnaps(dermsnapUser.ID)
			Expect(err).To(BeNil())
			Expect(len(dermsnaps)).To(Equal(3))
			Expect(dermsnaps[len(dermsnaps)-1].Itchy).To(Equal(true))
		})
	})
})
