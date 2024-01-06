package services_test

import (
	"dermsnap/models"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
	Describe("CreateUser", func() {
		It("should create a user", func() {
			user, err := Services.UserService.CreateUser("user1", models.Admin, models.Google)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(BeNil())
			Expect(user.Identifier).To(Equal("user1"))
		})

		It("should not create a user with the same identifier", func() {
			user, err := Services.UserService.CreateUser("user1", models.Admin, models.Google)
			Expect(err).ToNot(BeNil())
			Expect(user).To(BeNil())
		})
	})

	Describe("GetUserByIdentifier", func() {
		BeforeEach(func() {
			user, err := Services.UserService.CreateUser("user2", models.Admin, models.Google)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(BeNil())
			Expect(user.Identifier).To(Equal("user2"))
		})
		It("should get a user by identifier", func() {
			found, err := Services.UserService.GetUserByIdentifier("user2", models.Google)
			Expect(err).To(BeNil())
			Expect(found).ToNot(BeNil())
			Expect(found.ID).ToNot(BeNil())
			Expect(found.Identifier).To(Equal("user2"))
		})
	})

	Describe("GetUserById", func() {
		var userID string
		BeforeEach(func() {
			user, err := Services.UserService.CreateUser("user3", models.Admin, models.Google)
			userID = user.ID.String()
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(BeNil())
			Expect(user.Identifier).To(Equal("user3"))
		})
		It("should get a user by id", func() {
			found, err := Services.UserService.GetUserByID(userID)
			Expect(err).To(BeNil())
			Expect(found).ToNot(BeNil())
			Expect(found.ID).ToNot(BeNil())
			Expect(found.Identifier).To(Equal("user3"))
		})
	})

	Describe("CreateDoctorInfo", func() {
		var doctorId uuid.UUID
		BeforeEach(func() {
			user, err := Services.UserService.CreateUser("doctor1", models.Doctor, models.Doximity)
			doctorId = user.ID
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(BeNil())
			Expect(user.Identifier).To(Equal("doctor1"))
		})

		It("should create a doctor info", func() {
			doctorInfo, err := Services.UserService.CreateDoctorInfo(doctorId, models.CreateDoctorInfo{
				Specialty:   "Dermatology",
				Credentials: "MD",
			})
			Expect(err).To(BeNil())
			Expect(doctorInfo).ToNot(BeNil())
			Expect(doctorInfo.ID).ToNot(BeNil())
			Expect(doctorInfo.Specialty).To(Equal("Dermatology"))
			Expect(doctorInfo.Credentials).To(Equal("MD"))
		})
	})

	Describe("CreateUserInfo", Ordered, func() {
		var userId uuid.UUID
		BeforeAll(func() {
			user, err := Services.UserService.CreateUser("user4", models.Doctor, models.Doximity)
			userId = user.ID
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(BeNil())
			Expect(user.Identifier).To(Equal("user4"))
		})

		It("should create a user info", func() {
			userInfo, err := Services.UserService.CreateUserInfo(userId, models.CreateUserInfo{
				Height: 180,
				Weight: 80,
				Age:    30,
				Gender: "male",
			})
			Expect(err).To(BeNil())
			Expect(userInfo).ToNot(BeNil())
			Expect(userInfo.ID).ToNot(BeNil())
			Expect(userInfo.Weight).To(Equal(80))
			Expect(userInfo.Gender).To(Equal("male"))
		})

		It("should get user info", func() {
			userInfo, err := Services.UserService.GetUserInfo(userId)
			Expect(err).To(BeNil())
			Expect(userInfo).ToNot(BeNil())
			Expect(userInfo.ID).ToNot(BeNil())
			Expect(userInfo.UserID).To(Equal(userId))
			Expect(userInfo.Weight).To(Equal(80))
			Expect(userInfo.Gender).To(Equal("male"))
		})
	})
})
