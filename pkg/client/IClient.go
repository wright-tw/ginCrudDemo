package client

type IClient interface {
	CallAPI(method string, url string, params string, header map[string]string) (string, error)
}
