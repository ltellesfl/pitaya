// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cluster

import (
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/topfreegames/pitaya/internal/message"
	"github.com/topfreegames/pitaya/protos"
	"github.com/topfreegames/pitaya/route"
	"github.com/topfreegames/pitaya/session"
)

// NatsRPCClient struct
type NatsRPCClient struct {
	connString string
	server     *Server
	conn       *nats.Conn
	reqTimeout time.Duration
	running    bool
}

// NewNatsRPCClient ctor
func NewNatsRPCClient(connectString string, server *Server) *NatsRPCClient {
	ns := &NatsRPCClient{
		connString: connectString,
		reqTimeout: time.Duration(5) * time.Second,
		server:     server,
		running:    false,
	}
	return ns
}

// Send publishes a message in a given topic
// TODO handle errors (to the client)
func (ns *NatsRPCClient) Send(topic string, data []byte) error {
	return ns.conn.Publish(topic, data)
}

// Call calls a method remotally
// TODO use channel and create async Go method
// TODO oh my, this is hacky! will it perform good?
// even if it performs we need better concurrency control!!!
// TODO should we permit bigger mailbox size? current design only allows 1 message per user to be processes at a time
func (ns *NatsRPCClient) Call(
	rpcType protos.RPCType,
	route *route.Route,
	session *session.Session,
	msg *message.Message,
	server *Server,
) (*protos.Response, error) {

	req := protos.Request{
		Type: rpcType,
		Msg: &protos.Msg{
			Route: route.String(),
			Data:  msg.Data,
		},
	}

	if ns.server.Frontend {
		req.FrontendID = ns.server.ID
	}

	switch msg.Type {
	case message.Request:
		req.Msg.Type = protos.MsgType_MsgRequest
	case message.Notify:
		req.Msg.Type = protos.MsgType_MsgNotify
	}

	if rpcType == protos.RPCType_Sys {
		mid := uint(0)
		if msg.Type == message.Request {
			mid = msg.ID
		}
		req.Msg.ID = uint64(mid)
		req.Session = &protos.Session{
			ID:   session.ID(),
			Uid:  session.UID(),
			Data: session.GetDataEncoded(),
		}
	}

	var marshalledData []byte
	var err error
	marshalledData, err = proto.Marshal(&req)
	if err != nil {
		return nil, err
	}

	m, err := ns.conn.Request(getChannel(server.Type, server.ID), marshalledData, ns.reqTimeout)
	if err != nil {
		return nil, err
	}

	res := &protos.Response{}
	err = proto.Unmarshal(m.Data, res)

	if err != nil {
		return nil, err
	}

	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}

// Init inits nats rpc server
func (ns *NatsRPCClient) Init() error {
	ns.running = true
	conn, err := setupNatsConn(ns.connString)
	if err != nil {
		return err
	}
	ns.conn = conn
	return nil
}

// AfterInit runs after initialization
func (ns *NatsRPCClient) AfterInit() {}

// BeforeShutdown runs before shutdown
func (ns *NatsRPCClient) BeforeShutdown() {}

// Shutdown stops nats rpc server
func (ns *NatsRPCClient) Shutdown() error {
	return nil
}

func (ns *NatsRPCClient) stop() {
	ns.running = false
}

func (ns *NatsRPCClient) getSubscribeChannel() string {
	return fmt.Sprintf("pitaya/servers/%s/%s", ns.server.Type, ns.server.ID)
}