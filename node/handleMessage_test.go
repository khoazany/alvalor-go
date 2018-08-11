// Copyright (c) 2017 The Alvalor Authors
//
// This file is part of Alvalor.
//
// Alvalor is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Alvalor is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Alvalor.  If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"io/ioutil"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/alvalor/alvalor-go/types"
)

func TestMessage(t *testing.T) {
	suite.Run(t, new(MessageSuite))
}

type MessageSuite struct {
	suite.Suite
	log zerolog.Logger
	wg  *sync.WaitGroup
}

func (suite *MessageSuite) SetupTest() {
	suite.log = zerolog.New(ioutil.Discard)
	suite.wg = &sync.WaitGroup{}
	suite.wg.Add(1)
}

func (suite *MessageSuite) TestMessageTransaction() {

	// arrange
	address := "192.0.2.100:1337"

	msg := &types.Transaction{}
	msg.Hash = msg.GetHash()

	net := &NetworkMock{}

	finder := &PathfinderMock{}

	handlers := &HandlersMock{}
	handlers.On("Entity", mock.Anything)

	// act
	// TODO: introduce new blockchain interface
	handleMessage(suite.log, suite.wg, net, finder, nil, handlers, address, msg)

	// assert
	t := suite.T()

	handlers.AssertCalled(t, "Entity", msg)
}
