package dom

type SVGTests interface {
	HasExtension(extension string) (b bool)

	RequiredExtensions() (requiredExtensions *SVGStringList)

	RequiredFeatures() (requiredFeatures *SVGStringList)

	SystemLanguage() (systemLanguage *SVGStringList)
}
