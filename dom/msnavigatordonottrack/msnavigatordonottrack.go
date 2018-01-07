package msnavigatordonottrack

import (
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/joy/dom/utils"
)

// MSNavigatorDoNotTrack interface
// js:"MSNavigatorDoNotTrack"
type MSNavigatorDoNotTrack interface {

	// ConfirmSiteSpecificTrackingException
	// js:"confirmSiteSpecificTrackingException"
	// jsrewrite:"$_.confirmSiteSpecificTrackingException($1)"
	ConfirmSiteSpecificTrackingException(args *navigation.ConfirmSiteSpecificExceptionsInformation) (b bool)

	// ConfirmWebWideTrackingException
	// js:"confirmWebWideTrackingException"
	// jsrewrite:"$_.confirmWebWideTrackingException($1)"
	ConfirmWebWideTrackingException(args *utils.ExceptionInformation) (b bool)

	// RemoveSiteSpecificTrackingException
	// js:"removeSiteSpecificTrackingException"
	// jsrewrite:"$_.removeSiteSpecificTrackingException($1)"
	RemoveSiteSpecificTrackingException(args *utils.ExceptionInformation)

	// RemoveWebWideTrackingException
	// js:"removeWebWideTrackingException"
	// jsrewrite:"$_.removeWebWideTrackingException($1)"
	RemoveWebWideTrackingException(args *utils.ExceptionInformation)

	// StoreSiteSpecificTrackingException
	// js:"storeSiteSpecificTrackingException"
	// jsrewrite:"$_.storeSiteSpecificTrackingException($1)"
	StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation)

	// StoreWebWideTrackingException
	// js:"storeWebWideTrackingException"
	// jsrewrite:"$_.storeWebWideTrackingException($1)"
	StoreWebWideTrackingException(args *utils.StoreExceptionsInformation)
}
