package stores

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	GetPatientByPhoneAndEmail(ctx *goofy.Context, phone, email string) (string, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}

type Doctor interface {
	Create(ctx *goofy.Context, patient *models.DoctorDetails) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.DoctorDetails, error)
	Update(ctx *goofy.Context, patientDetails *models.DoctorRequest, id string) (*models.DoctorDetails, error)
	Delete(ctx *goofy.Context, id string) error
	CheckDoctorExist(ctx *goofy.Context, phone, email, licenceNumber string) (string, error)
}
