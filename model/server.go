package model

// https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L456
type Server struct {
	MaxCapacities int // Maximum capacities the server.
	InUse         int // The number of servers currently in use.
	Idle          int // The number of idle servers.
}

func NewServer() *Server {
	return &Server{
		MaxCapacities: 1,
		InUse:         0,
		Idle:          1,
	}
}

func (s *Server) IsBusy() bool {
	return s.MaxCapacities == s.InUse
}
