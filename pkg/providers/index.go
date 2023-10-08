package providers

type Providers struct {
	HTTPProvider HTTPProviderInterface
}

func (p Providers) UseHTTPProvider() {
	p.HTTPProvider = NewHTTPProvider()
}
