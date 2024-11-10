package procedures

import (
	"log"
	"net"

	cart "cart-service/handlers/cart"
	cartProto "cart-service/proto/cart"

	"google.golang.org/grpc"
)

type Procedures struct {
	Listen net.Listener
	Grpc   *grpc.Server
	Cart   *cart.Handler
}

func (p *Procedures) setupRegister() {
	p.registerCart()
}

func (p *Procedures) registerCart() {
	cartProto.RegisterCartServiceServer(p.Grpc, p.Cart)
}

func (p *Procedures) RunRpcServer(port string) {
	p.setupRegister()

	log.Printf("[Running-Success] gRPC server on localhost on port : %s", port)
	if err := p.Grpc.Serve(p.Listen); err != nil {
		panic(err)
	}

	defer p.Listen.Close()
}
