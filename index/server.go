// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package index

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/blast/protobuf/raft"
)

type Server struct {
	node      *raft.Node
	bootstrap bool
	joinAddr  string

	raftServer *RaftServer

	grpcService *GRPCService
	grpcServer  *GRPCServer
	grpcClient  *GRPCClient

	httpServer *HTTPServer

	logger     *log.Logger
	httpLogger *log.Logger
}

func NewServer(nodeId string, bindAddr string, grpcAddr string, httpAddr string, dataDir string, joinAddr string, indexMappingPath string, logger *log.Logger, httpLogger *log.Logger) *Server {
	var err error

	server := &Server{
		bootstrap:  joinAddr == "",
		joinAddr:   joinAddr,
		logger:     logger,
		httpLogger: httpLogger,
	}

	// set default index mapping
	indexMapping := mapping.NewIndexMapping()

	// check index mapping file
	if indexMappingPath != "" {
		_, err = os.Stat(indexMappingPath)
		if err == nil {
			// read index mapping file
			indexMappingFile, err := os.Open(indexMappingPath)
			if err != nil {
				server.logger.Printf("[ERR] %v", err)
				return nil
			}
			defer func() {
				err = indexMappingFile.Close()
				if err != nil {
					server.logger.Printf("[ERR] %v", err)
				}
			}()

			b, err := ioutil.ReadAll(indexMappingFile)
			if err != nil {
				server.logger.Printf("[ERR] %v", err)
				return nil
			}

			err = json.Unmarshal(b, indexMapping)
			if err != nil {
				server.logger.Printf("[ERR] %v", err)
				return nil
			}
		} else if os.IsNotExist(err) {
			server.logger.Printf("[ERR] %v", err)
			return nil
		}
	}

	// create node information
	server.node = &raft.Node{
		Id:       nodeId,
		BindAddr: bindAddr,
		GrpcAddr: grpcAddr,
		HttpAddr: httpAddr,
		DataDir:  dataDir,
	}

	// create raft server
	server.raftServer, err = NewRaftServer(server.node, server.bootstrap, indexMapping, server.logger)
	if err != nil {
		server.logger.Printf("[ERR] %v", err)
		return nil
	}

	// create gRPC service
	server.grpcService, err = NewGRPCService(server.raftServer, server.logger)
	if err != nil {
		server.logger.Printf("[ERR] %v", err)
		return nil
	}

	// create gRPC server
	server.grpcServer, err = NewGRPCServer(grpcAddr, server.grpcService, server.logger)
	if err != nil {
		server.logger.Printf("[ERR] %v", err)
		return nil
	}

	// create gRPC client for HTTP server
	server.grpcClient, err = NewGRPCClient(grpcAddr)
	if err != nil {
		server.logger.Printf("[ERR] %v", err)
		return nil
	}

	// create HTTP server
	server.httpServer, err = NewHTTPServer(httpAddr, server.grpcClient, server.logger, server.httpLogger)
	if err != nil {
		server.logger.Printf("[ERR] %v", err)
		return nil
	}

	return server
}

func (s *Server) Start() {
	// start Raft server
	go func() {
		err := s.raftServer.Start()
		if err != nil {
			s.logger.Printf("[ERR] %v", err)
			return
		}
	}()
	s.logger.Print("[INFO] Raft server started")

	// start gRPC server
	go func() {
		err := s.grpcServer.Start()
		if err != nil {
			s.logger.Printf("[ERR] %v", err)
			return
		}
	}()
	s.logger.Print("[INFO] gRPC server started")

	// start HTTP server
	go func() {
		err := s.httpServer.Start()
		if err != nil {
			s.logger.Printf("[ERR] %v", err)
			return
		}
	}()
	s.logger.Print("[INFO] HTTP server started")

	if !s.bootstrap {
		// create gRPC client
		client, err := NewGRPCClient(s.joinAddr)
		defer func() {
			err := client.Close()
			if err != nil {
				s.logger.Printf("[ERR] %v", err)
			}
		}()
		if err != nil {
			s.logger.Printf("[ERR] %v", err)
			return
		}

		// join to the existing cluster
		err = client.Join(s.node)
		if err != nil {
			s.logger.Printf("[ERR] %v", err)
			return
		}
	}
}

func (s *Server) Stop() {
	// stop HTTP server
	err := s.httpServer.Stop()
	if err != nil {
		s.logger.Printf("[ERR] %v", err)
	}

	// close gRPC client
	err = s.grpcClient.Close()
	if err != nil {
		s.logger.Printf("[ERR] %v", err)
	}

	// stop gRPC server
	err = s.grpcServer.Stop()
	if err != nil {
		s.logger.Printf("[ERR] %v", err)
	}

	// stop Raft server
	err = s.raftServer.Stop()
	if err != nil {
		s.logger.Printf("[ERR] %v", err)
	}
}
