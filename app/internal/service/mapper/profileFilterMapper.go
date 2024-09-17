package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileFilterMapper struct {
}

func (pm *ProfileFilterMapper) MapToAddRequest(pr *request.ProfileAddRequestDto) *entity.ProfileFilterEntity {
	return &entity.ProfileFilterEntity{
		SessionID:    pr.SessionID,
		SearchGender: pr.SearchGender,
		LookingFor:   pr.LookingFor,
		AgeFrom:      pr.AgeFrom,
		AgeTo:        pr.AgeTo,
		Distance:     pr.Distance,
		Page:         pr.Page,
		Size:         pr.Size,
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
