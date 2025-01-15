package test

import (
	"testing"
	"fmt"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

	"example.com/TRYSE8/entity"
)

func TestStudentID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("StudentID is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "",
			FirstName: "juju",
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run("StudentID is invalid", func(t *testing.T){
		users := entity.User{
			StudentID: "B65113266",
			FirstName: "juju",
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[MDB]\\d{7}$)", users.StudentID)))
	})
}

func TestFirstName(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("FirstName is required", func(t *testing.T){
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "",
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})
}
