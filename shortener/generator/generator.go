package generator

type Generator interface {
	GenerateShortUrl(string, string) string
}
