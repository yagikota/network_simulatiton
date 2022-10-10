package model

// https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L456
// https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L1063
type Server struct {
	MaxCapacities int // Maximum capacities the server.
	InUse         int // The number of servers currently in use.
}

func NewServer(cap int) *Server {
	return &Server{
		MaxCapacities: cap,
		InUse:         0,
	}
}

// Use uses n servers.
func (s *Server) Use(n int) {
	s.InUse += n
}

// Free frees n servers.
func (s *Server) Free(n int) {
	s.InUse -= n
}

// IsBusy returns true if the server has no capable capacities.
func (s *Server) IsBusy() bool {
	return s.MaxCapacities <= s.InUse
}
