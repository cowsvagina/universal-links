package uni_links

import (
	"net/http"
	"strings"
)

type Domain string

type HTTPServer struct {
	addr          string
	associations  map[Domain]string
	afterResponse func(*http.Request)
}

func NewHTTPServer(addr string, afterResponse func(r *http.Request)) *HTTPServer {
	return &HTTPServer{
		addr:          addr,
		associations:  make(map[Domain]string),
		afterResponse: afterResponse,
	}
}

func (s *HTTPServer) AddAssociation(d Domain, a AssociationMarshaler) error {
	v, err := a.Marshal()
	if err != nil {
		return err
	}

	s.associations[d] = v
	return nil
}

func (s *HTTPServer) Serve() error {
	return http.ListenAndServe(s.addr, s)
}

func (s *HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")[0]
	association, ok := s.associations[Domain(host)]

	if req.URL.Path != "/.well-known/apple-app-site-association" {
		resp.WriteHeader(http.StatusForbidden)
		_, _ = resp.Write([]byte("invalid path"))
		goto end
	}

	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusForbidden)
		_, _ = resp.Write([]byte("invalid method"))
		goto end
	}

	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		_, _ = resp.Write([]byte("unsupported domain"))
	} else {
		resp.Header().Set("Cache-Control", "no-cache")
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(http.StatusOK)
		_, _ = resp.Write([]byte(association))
	}

end:
	if s.afterResponse != nil {
		s.afterResponse(req)
	}
}
