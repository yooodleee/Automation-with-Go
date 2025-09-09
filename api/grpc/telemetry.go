package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	xr "grpc/proto/ems"
)


// Provides the user/password for the connection
// It implements the PerRPCCredentials interface.
type loginCreds struct {
	Username, Password	string
	requireTLS			bool
}

// Method of the PerRPCCredentials interface.
func(c *loginCreds) GetRequestMetadata(
	context.Context,
	...string,
) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

// Method of the PerRPCCredentials interface.
func (c *loginCreds) RequiredTransportSecurity() bool {
	return c.requireTLS
}

// GetSubscription follows the channel Generator Pattern, it returns
// a []byte channel where the Streaming Telemetry data is sent/received.
// It also propagates error messages on an error channel.
func (x *xrgrpc) GetSubscription(
	sub, enc string,
) (chan []byte, chan error, error) {
	encodingMap := map[string]int64{
		"gpb": 2,
		"gpbkv": 3,
		"json": 4,
	}

	encoding, ok := encodingMap[enc]
	if !ok {
		return nil, nil, fmt.Errorf(
			"encoding value not supported: %s",
			enc,
		)
	}

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63()

	// 'c' is the gRPC stub.
	c := xr.NewGRPCConfigOperClient(x.conn)

	// 'b' is the bytes channel where Telemetry data is sent.
	b := make(chan []byte)

	// 'e' is the error channel where error messages are sent.
	e := make(chan error)

	// 'a' is the object, send to the router via the stub.
	a := xr.CreateSubsArgs{ReqId: id, Encode: encoding, Subidstr: sub}

	// 'r' is the result that comes back from the target.
	st, err := c.CreateSubs(x.ctx, &a)
	if err != nil {
		return b, e, fmt.Errorf("gRPC CreateSubs failed: %w", err)
	}

	// TODO: Review the logic. Make sure this goroutine ends and propagate
	// error messages
	go func() {
		r, err := st.Recv()
		if err != nil {
			close(b)
			e <- fmt.Errorf("error triggered by remote host for ReqId: %v; %s", id, r.GetErrors())
			return
		}
		if len(r.GetErrors()) != 0 {
			close(b)
			e <- fmt.Errorf("error triggered by remote host for ReqId: %v; %s", id, r.GetErrors())
			return
		}
		for {
			select {
			case <-x.ctx.Done():
				close(b)
				return
			case b <- r.GetData():
				r, err = st.Recv()
				if err == io.EOF {
					close(b)
					return
				}
				if err != nil {
					// don't report this error for now. 
					// sent and main does nto receive it, it would hang forever.
					// e <- fmt.Errorf("%s, ReqId: %s", err, si)
					close(b)
					return
				}
			}
		}
	}()
	return b, e, err
}


func (x *xrgrpc) SessionCancel(
	e chan error,
	c chan os.Signal,
	stop context.CancelFunc,
) {
	select {
	case <- c:
		fmt.Printf(
			"\nmanually cancelled the session to %v\n\n",
			x.conn.Target(),
		)
		stop()
		return
	case <-x.ctx.Done():
		// Timeout: "context deadline exceeded"
		err := x.ctx.Err()
		fmt.Printf(
			"\ngRPC session timed out after %s seconds: %v\n\n",
			"10",
			err.Error(),
		)
		return
	case err := <-e:
		// Session canceled: "context canceled"
		fmt.Printf(
			"\ngRPC session to %v failed: %v\n\n",
			x.conn.Target(),
			err.Error(),
		)
		return
	}
}