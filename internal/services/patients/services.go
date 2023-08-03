package patients

import (
	"main/internal/models"
	"main/internal/stores"
	"main/utils"
	"net/http"

	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
)

type PatientService struct {
	store stores.Patient
}

func New(store stores.Patient) *PatientService {
	return &PatientService{store}
}

func (s *PatientService) Create(ctx *goofy.Context, patient *models.PatientRequest) (interface{}, error) {
	existingId, err := s.store.GetPatientByPhoneAndEmail(ctx, patient.Phone, patient.Email)
	if _, ok := err.(errors.EntityNotFound); !ok {
		return nil, err
	}

	if existingId != "" {
		return nil, errors.Response{
			StatusCode: http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Reason:     "patient already exist with email/phone",
		}
	}

	patientDetails := models.PatientDetails{
		Name:    patient.Name,
		Gender:  patient.Gender,
		Phone:   patient.Phone,
		Email:   patient.Email,
		Age:     patient.Age,
		City:    patient.City,
		State:   patient.State,
		Pincode: patient.Pincode,
		Status:  "New",
	}

	password, err := utils.GenerateNumericUniqueId(8)
	if err != nil {
		ctx.Logger.Errorf("error while generating autogenerated password for new patient")
		return nil, err
	}

	patientDetails.Password = password

	id, err := utils.GenerateAlphaNumericUniqueId(8)
	if err != nil {
		ctx.Logger.Errorf("error while generating patient id")
		return nil, err
	}

	patientDetails.Id = id

	salt, err := utils.GenerateAlphaNumericUniqueId(12)
	if err != nil {
		ctx.Logger.Errorf("error while generating salt for new patient")
		return nil, err
	}

	patientDetails.Salt = salt

	res, err := s.store.Create(ctx, &patientDetails)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PatientService) Get(ctx *goofy.Context, id string) (*models.PatientDetails, error) {
	patientDetails, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return patientDetails, nil
}

func (s *PatientService) Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error) {
	existingPatient, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if (patientDetails.Phone != "" || patientDetails.Email != "") &&
		(existingPatient.Phone != patientDetails.Phone || existingPatient.Email != patientDetails.Email) {
		existingId, err := s.store.GetPatientByPhoneAndEmail(ctx, existingPatient.Phone, existingPatient.Email)
		if err != nil {
			return nil, err
		}

		if existingId != "" {
			return nil, errors.Response{
				StatusCode: http.StatusBadRequest,
				Status:     http.StatusText(http.StatusBadRequest),
				Reason:     "email or phone already registered with another patient",
			}
		}
	}

	updatedResponse, err := s.store.Update(ctx, patientDetails, existingPatient.Id)
	if err != nil {
		return nil, err
	}

	return updatedResponse, nil
}

func (s *PatientService) Delete(ctx *goofy.Context, id string) error {
	patient, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.store.Delete(ctx, patient.Id)
	if err != nil {
		return err
	}

	return nil
}
