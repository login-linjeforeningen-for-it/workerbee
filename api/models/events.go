package models

import (
	"encoding/json"
	"workerbee/internal"
)

type EventBase struct {
	ID                 int                 `db:"id" json:"id,omitempty"`
	Visible            bool                `db:"visible" json:"visible,omitempty"`
	NameNo             string              `db:"name_no" json:"name_no" validate:"required"`
	NameEn             string              `db:"name_en" json:"name_en" validate:"required"`
	DescriptionNo      string              `db:"description_no" json:"description_no" validate:"required"`
	DescriptionEn      string              `db:"description_en" json:"description_en" validate:"required"`
	InformationalNo    *string             `db:"informational_no" json:"informational_no,omitempty"`
	InformationalEn    *string             `db:"informational_en" json:"informational_en,omitempty"`
	TimeType           string              `db:"time_type" json:"time_type,omitempty"`
	TimeStart          internal.LocalTime  `db:"time_start" json:"time_start" validate:"required,beforeField=TimeEnd"`
	TimeEnd            internal.LocalTime  `db:"time_end" json:"time_end" validate:"required,afterField=TimeStart"`
	TimePublish        internal.LocalTime  `db:"time_publish" json:"time_publish" validate:"required,beforeField=TimeStart"`
	TimeSignupRelease  *internal.LocalTime `db:"time_signup_release" json:"time_signup_release,omitempty" validate:"omitempty,beforeField=TimeSignupDeadline"`
	TimeSignupDeadline *internal.LocalTime `db:"time_signup_deadline" json:"time_signup_deadline,omitempty" validate:"omitempty,afterField=TimeSignupRelease"`
	Canceled           bool                `db:"canceled" json:"canceled"`
	Digital            bool                `db:"digital" json:"digital"`
	Highlight          bool                `db:"highlight" json:"highlight"`
	ImageSmall         *string             `db:"image_small" json:"image_small,omitempty"`
	ImageBanner        *string             `db:"image_banner" json:"image_banner,omitempty"`
	LinkFacebook       *string             `db:"link_facebook" json:"link_facebook,omitempty"`
	LinkDiscord        *string             `db:"link_discord" json:"link_discord,omitempty"`
	LinkSignup         *string             `db:"link_signup" json:"link_signup,omitempty"`
	LinkStream         *string             `db:"link_stream" json:"link_stream,omitempty"`
	Capacity           *int                `db:"capacity" json:"capacity,omitempty"`
	IsFull             bool                `db:"is_full" json:"is_full"`
	ParentID           *int                `db:"parent_id" json:"parent_id,omitempty"`
	UpdatedAt          internal.LocalTime  `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt          internal.LocalTime  `db:"created_at" json:"created_at,omitempty"`
}

type Event struct {
	EventBase
	Category     Category      `db:"category" json:"category"`
	Location     *Location     `db:"location" json:"location"`
	Rule         *Rule         `db:"rules" json:"rule"`
	Audience     *Audience     `db:"audience" json:"audience"`
	Organization *Organization `db:"organization" json:"organization"`
}

type NewEvent struct {
	EventBase
	CategoryID     int   `db:"category_id" json:"category_id,omitempty" validate:"required"`
	LocationID     *int  `db:"location_id" json:"location_id"`
	RuleID         *int  `db:"rule_id" json:"rule_id,omitempty"`
	AudienceID     *int  `db:"audience_id" json:"audience_id,omitempty"`
	OrganizationID *int  `db:"organization_id" json:"organization_id"`
	CreatedIDs     []int `db:"-" json:"created_ids,omitempty"`
}

type EventWithTotalCount struct {
	Event
	TotalCount int `db:"total_count" json:"-"`
}

type EventCategory struct {
	Category
}

type EventName struct {
	ID        int                `db:"id" json:"id"`
	NameEn    string             `db:"name_en" json:"name_en"`
	TimeStart internal.LocalTime `db:"time_start" json:"time_start"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event

	aux := &struct {
		*Alias
		Location     *Location     `json:"location"`
		Rule         *Rule         `json:"rule"`
		Audience     *Audience     `json:"audience"`
		Organization *Organization `json:"organization"`
	}{
		Alias: (*Alias)(&e),
	}

	if e.Location != nil && e.Location.ID == nil {
		aux.Location = nil
	} else {
		aux.Location = e.Location
	}

	if e.Rule != nil && e.Rule.ID == nil {
		aux.Rule = nil
	} else {
		aux.Rule = e.Rule
	}

	if e.Audience != nil && e.Audience.ID == nil {
		aux.Audience = nil
	} else {
		aux.Audience = e.Audience
	}

	if e.Organization != nil && e.Organization.ID == nil {
		aux.Organization = nil
	} else {
		aux.Organization = e.Organization
	}

	return json.Marshal(aux)
}
