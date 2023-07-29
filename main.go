package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	patientsHandler "main/internal/http/patients"
	patientsSvc "main/internal/services/patients"
	patientsStore "main/internal/stores/patients"

	doctorOPDScheduleHandler "main/internal/http/doctors/doctorOPDSchedule"
	doctorOPDScheduleService "main/internal/services/doctors/doctorOPDSchedule"
	doctorOPDScheduleStore "main/internal/stores/doctors/doctorOPDSchedule"
)

func main() {

	app := goofy.New()

	// Store Layer
	patientStore := patientsStore.New(app.Database)
	doctorOPDScheduleStore := doctorOPDScheduleStore.New(app.Database)

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)
	doctorOPDScheduleSvc := doctorOPDScheduleService.New(doctorOPDScheduleStore)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)
	doctorOPDScheduleHandler := doctorOPDScheduleHandler.New(doctorOPDScheduleSvc)

	app.POST("/patient", patientHandler.Create)

	app.POST("/doctor-opd-schedule", doctorOPDScheduleHandler.Create)
	app.GET("/doctor-opd-schedule", doctorOPDScheduleHandler.Index)
	app.GET("/doctor-opd-schedule/{id}", doctorOPDScheduleHandler.Read)
	app.PUT("/doctor-opd-schedule/{id}", doctorOPDScheduleHandler.Update)

	app.Start()
}
