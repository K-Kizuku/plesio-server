package controller

import "github.com/K-Kizuku/plesio-server/app/usecase"

type MeetingController struct {
	MeetingUsecase usecase.IMeetingUsecase
}

type IMeetingController interface {
	Do()
}

func NewMeetingContrallor(meetingUsecase usecase.IMeetingUsecase) IMeetingController {
	return &MeetingController{
		MeetingUsecase: meetingUsecase,
	}
}

func (m *MeetingController) Do() {
}

// func (c *Controller) MeetingController() {

// }
