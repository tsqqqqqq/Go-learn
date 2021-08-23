package infar

type Retrieve interface {
	Get(url string) string
}
