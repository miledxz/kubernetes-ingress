// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e_sequential

package httpsruntime

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/haproxytech/kubernetes-ingress/deploy/tests/e2e"
)

type HTTPSSuite struct {
	suite.Suite
	test     e2e.Test
	client   *e2e.Client
	tmplData tmplData
}

type tmplData struct {
	Port string
}

func (suite *HTTPSSuite) SetupSuite() {
	var err error
	suite.test, err = e2e.NewTest()
	suite.Require().NoError(err)
}

func (suite *HTTPSSuite) TearDownSuite() {
	suite.test.TearDown()
}

func TestHTTPSSuite(t *testing.T) {
	suite.Run(t, new(HTTPSSuite))
}