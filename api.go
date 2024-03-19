package main

type APIServer struct {
	listenAddr string
}

func newAPIServer(listenAddr sting) * APIServer{
	return &APIServer{
		listenAddr: listenAddr,
	}
}

