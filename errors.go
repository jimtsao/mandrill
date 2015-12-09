package mandrill

import (
	"errors"
	"fmt"
)

type apiError struct {
	Status  string
	Code    int
	Name    string
	Message string
}

var (
	ErrInvalidKey            = errors.New("The provided API key is not a valid Mandrill API key")
	ErrValidation            = errors.New("The parameters passed to the API call are invalid or not provided when required")
	ErrGeneral               = errors.New("An unexpected error occurred processing the request. Mandrill developers will be notified.")
	ErrPayment               = errors.New("The requested feature requires payment")
	ErrSubaccount            = errors.New("The provided subaccount id does not exist")
	ErrInvalidTagName        = errors.New("The requested tag does not exist or contains invalid characters")
	ErrServiceUnavailable    = errors.New("The subsystem providing this API call is down for maintenance")
	ErrUnknownSender         = errors.New("The requested sender does not exist")
	ErrUnknownURL            = errors.New("The requested URL has not been seen in a tracked link")
	ErrUnknownTrackingDomain = errors.New("The provided tracking domain does not exist")
	ErrUnknownTemplate       = errors.New("The given template name already exists or contains invalid characters")
)

func (a *apiError) Error() error {
	switch a.Name {
	case "Invalid_Key":
		return ErrInvalidKey
	case "ValidationError":
		return ErrValidation
	case "GeneralError":
		return ErrGeneral
	case "PaymentRequired":
		return ErrPayment
	case "Unknown_Subaccount":
		return ErrSubaccount
	case "Invalid_Tag_Name":
		return ErrInvalidTagName
	case "ServiceUnavailable":
		return ErrServiceUnavailable
	case "Unknown_Sender":
		return ErrUnknownSender
	case "Unknown_Url":
		return ErrUnknownURL
	case "Unknown_TrackingDomain":
		return ErrUnknownTrackingDomain
	case "Unknown_Template":
		return ErrUnknownTemplate
	default:
		return fmt.Errorf("An unknown error response was received from API. %+v", a)
	}
}
