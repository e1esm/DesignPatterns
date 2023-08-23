package main

type Nginx struct {
	Application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		Application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiter(url)
	if allowed {
		return n.Application.handleRequest(url, method)
	}
	return 403, "Not allowed"
}

func (n *Nginx) checkRateLimiter(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}
