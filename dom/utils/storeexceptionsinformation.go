package utils

type StoreExceptionsInformation struct {
	*ExceptionInformation

	detailURI         *string
	explanationString *string
	siteName          *string
}
