package proxy

//Proxy

type Ngingx struct {
	Application     Server
	MaxAllowRequest int
	RateLimiter     map[string]int
}

func NewNginxServer() *Ngingx {
	return &Ngingx{
		Application:     &Application{},
		MaxAllowRequest: 2,
		RateLimiter:     make(map[string]int),
	}
}

func (n *Ngingx) HandleRequest(url, method string) (int, string) {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.Application.HandleRequest(url, method)
}

func (n *Ngingx) CheckRateLimiting(url string) bool {
	if n.RateLimiter[url] == 0 {
		n.RateLimiter[url] = 1
	}

	if n.RateLimiter[url] > n.MaxAllowRequest {
		return false
	}

	n.RateLimiter[url] = n.RateLimiter[url] + 1
	return true
}
