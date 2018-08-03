package allitebooks

type Builder interface {
	WithStartPage(int) Builder
	WithStartURL(string) Builder
	Build() Allitebooks
}

type builder struct {
	startPage int
	startURL  string
}

func (b *builder) WithStartPage(startPage int) Builder {
	b.startPage = startPage
	return b
}

func (b *builder) WithStartURL(startURL string) Builder {
	b.startURL = startURL
	return b
}

func (b *builder) Build() Allitebooks {
	return &allitebooks{
		startPage: b.startPage,
		startURL:  b.startURL,
	}
}

func New() *builder {
	return &builder{}
}
