/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"encoding/json"
	"net/http"

	go_restful "github.com/emicklei/go-restful"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/tkeel-io/core-broker/pkg/core"
	"github.com/tkeel-io/core-broker/pkg/types"
	"github.com/tkeel-io/kit/log"
)

type EntityService struct {
	msgChanMap map[string]map[string]chan []byte // entityID  clientID msgChan
	coreClient core.Client
}

func NewEntityService() *EntityService {
	msgChanMap := make(map[string]map[string]chan []byte)
	client, err := core.NewCoreClient()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &EntityService{msgChanMap: msgChanMap, coreClient: *client}
}

func (s *EntityService) Run() {
	var entityID string
	var msgData []byte
	for {
		msg := <-types.MsgChan
		log.Debugf("event msg data: %+v", msg.Data.AsInterface())
		log.Debugf("event msg data: %#v", msg.Data.AsInterface())
		log.Debugf("event msg data type: %T", msg.Data.AsInterface())
		switch kv := msg.Data.AsInterface().(type) {
		case map[string]interface{}:
			subID := types.Interface2string(kv["id"])
			entityID = types.GetEntityID(subID)
			msgData, _ = json.Marshal(kv["properties"])
		}

		if clientMsgChan, ok := s.msgChanMap[entityID]; ok {
			for _, msgChan := range clientMsgChan {
				msgChan <- msgData
			}
		}
	}
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func (s *EntityService) handleRequest(c *websocket.Conn, stopChan chan struct{}, msgChan chan []byte) {
	clientID := uuid.New().String()
	var entityID string
	for {
		_, p, err := c.ReadMessage()
		if err != nil {
			if _, ok := s.msgChanMap[entityID]; ok {
				delete(s.msgChanMap[entityID], clientID)
				if len(s.msgChanMap[entityID]) == 0 {
					subID := types.SubscriptionIDByJoin(entityID, types.Topic)
					if err := s.coreClient.Unsubscribe(subID); err != nil {
						log.Error("call unsubscribe entity error:", err)
					}
					delete(s.msgChanMap, entityID)
				}
			}
			close(stopChan)
			return
		}

		wsReq := types.WsRequest{}
		err = json.Unmarshal(p, &wsReq)
		if err != nil || wsReq.ID == "" {
			log.Error(err)
			continue
		}

		entityIDTemp := wsReq.ID
		if entityID == "" && entityIDTemp != "" {
			entityID = entityIDTemp
		} else if entityID != entityIDTemp {
			delete(s.msgChanMap[entityID], clientID)
			entityID = entityIDTemp
		}
		if _, ok := s.msgChanMap[entityID]; !ok {
			s.msgChanMap[entityID] = make(map[string]chan []byte)

			subID := types.SubscriptionIDByJoin(entityID, types.Topic)
			if err := s.coreClient.Subscribe(subID, entityID, types.Topic); err != nil {
				log.Error("call subscribing to core err:", err)
			}
		}
		s.msgChanMap[entityID][clientID] = msgChan
	}
}

func (s *EntityService) GetEntity(req *go_restful.Request, resp *go_restful.Response) {
	c, err := upgrader.Upgrade(resp, req.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	var stopChan = make(chan struct{})
	var msgChan = make(chan []byte)

	defer close(msgChan)

	go s.handleRequest(c, stopChan, msgChan)

	for {
		select {
		case msg := <-msgChan:
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		case <-stopChan:
			log.Info("ws stop")
			return
		}
	}
}
