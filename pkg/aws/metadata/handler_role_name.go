// Copyright 2017 uSwitch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package metadata

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/uswitch/kiam/pkg/k8s"
)

type roleHandler struct {
	roleFinder k8s.RoleFinder
	clientIP   clientIPFunc
}

func (h *roleHandler) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request) (int, error) {
	startTime := time.Now()
	defer handlerTimer.WithLabelValues("roleName").Observe(float64(time.Since(startTime) / time.Millisecond))

	err := req.ParseForm()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	ip, err := h.clientIP(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	role, err := findRole(ctx, h.roleFinder, ip)
	if err != nil {
		findRoleError.WithLabelValues("roleName").Inc()
		return http.StatusInternalServerError, err
	}

	if role == "" {
		emptyRole.WithLabelValues("roleName").Inc()
		return http.StatusNotFound, EmptyRoleError
	}

	fmt.Fprint(w, role)
	success.WithLabelValues("roleName").Inc()

	return http.StatusOK, nil
}
