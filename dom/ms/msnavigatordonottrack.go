package ms

import (
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/utils"
)

type MSNavigatorDoNotTrack interface {
	ConfirmSiteSpecificTrackingException(args *navigation.ConfirmSiteSpecificExceptionsInformation) (b bool)

	ConfirmWebWideTrackingException(args *utils.ExceptionInformation) (b bool)

	RemoveSiteSpecificTrackingException(args *utils.ExceptionInformation)

	RemoveWebWideTrackingException(args *utils.ExceptionInformation)

	StoreSiteSpecificTrackingException(args *utils.StoreSiteSpecificExceptionsInformation)

	StoreWebWideTrackingException(args *utils.StoreExceptionsInformation)
}
