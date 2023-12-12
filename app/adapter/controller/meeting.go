package controller

import "github.com/K-Kizuku/plesio-server/app/usecase"

type MeetingController struct {
	MeetingUsecase *usecase.MeetingUsecase
}

type IMeetingController interface {
	Do()
}

func NewMeetingContrallor(meetingUsecase *usecase.MeetingUsecase) IMeetingController {
	return &MeetingController{
		MeetingUsecase: meetingUsecase,
	}
}

func (m *MeetingController) Do() {

}

// func (c *Controller) MeetingController() {

// }
