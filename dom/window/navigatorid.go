package window

type NavigatorID interface {
	AppCodeName() (appCodeName string)

	AppName() (appName string)

	AppVersion() (appVersion string)

	Platform() (platform string)

	Product() (product string)

	ProductSub() (productSub string)

	UserAgent() (userAgent string)

	Vendor() (vendor string)

	VendorSub() (vendorSub string)
}
