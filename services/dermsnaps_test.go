package services_test

import (
	"dermsnap/models"
	"time"

	"github.com/lib/pq"
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
			Expect(dermsnap.UserID).To(Equal(dermsnapUser.ID))
			Expect(dermsnap.NewMedications).To(Equal(pq.StringArray{"finasteride"}))
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

	Describe("Update dermsnap", Ordered, func() {
		var dermsnap *models.Dermsnap
		BeforeAll(func() {
			dermsnap, _ = DermsnapService.CreateDermsnap(dermsnapUser.ID, models.CreateDermsnap{
				StartTime: time.Now(),
				Duration:  10,
				Locations: []models.BodyLocation{
					models.Arms,
					models.Abdomen,
				},
				Changed: false,
				NewMedications: []string{
					"finasteride",
				},
				Itchy:    true,
				Painful:  false,
				MoreInfo: "",
			})
		})
		It("should update a dermsnap", func() {
			dermsnap.Itchy = false
			dermsnap.Duration = 20
			dermsnap.NewMedications = []string{"finasteride", "minoxidil"}
			dermsnap.Locations = []string{"buttocks"}
			updatedDermsnap, err := DermsnapService.UpdateDermsnap(dermsnap.ID, dermsnap)
			Expect(err).To(BeNil())
			Expect(updatedDermsnap.Itchy).To(Equal(false))
			Expect(updatedDermsnap.NewMedications).To(Equal(pq.StringArray{"finasteride", "minoxidil"}))
			Expect(updatedDermsnap.Locations).To(Equal(pq.StringArray{"buttocks"}))
			Expect(updatedDermsnap.Duration).To(Equal(20))
		})
	})

	Describe("Delete dermsnap", func() {
		var dermsnap *models.Dermsnap
		BeforeAll(func() {
			dermsnap, _ = DermsnapService.CreateDermsnap(dermsnapUser.ID, models.CreateDermsnap{
				StartTime: time.Now(),
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
		})
		It("should delete a dermsnap", func() {
			deletedDermsnap, err := DermsnapService.DeleteDermsnap(dermsnap)
			Expect(err).To(BeNil())
			Expect(deletedDermsnap.ID).To(Equal(dermsnap.ID))

			_, err = DermsnapService.GetDermsnapById(dermsnap.ID)
			Expect(err).ToNot(BeNil())
		})
	})
})
