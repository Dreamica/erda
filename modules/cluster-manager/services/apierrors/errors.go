// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package apierrors

import (
	"github.com/erda-project/erda/pkg/http/httpserver/errorresp"
)

var (
	ErrPreCheckCluster = err("ErrPreCheckCluster", "auth failed")
	ErrCreateCluster   = err("ErrCreateCluster", "failed to create cluster")
	ErrUpdateCluster   = err("ErrUpdateCluster", "failed to update cluster")
	ErrPatchCluster    = err("ErrPatchCluster", "failed to patch cluster")
	ErrGetCluster      = err("ErrGetCluster", "failed to get cluster")
	ErrListCluster     = err("ErrListCluster", "failed to list cluster")
	ErrDeleteCluster   = err("ErrDeleteCluster", "failed to delete cluster")
)

func err(template, defaultValue string) *errorresp.APIError {
	return errorresp.New(errorresp.WithTemplateMessage(template, defaultValue))
}
