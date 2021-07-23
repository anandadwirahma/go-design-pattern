package proxy

type Server interface {
	HandleRequest(string2 string, string3 string) (int, string)
}
