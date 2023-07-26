// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import "github.com/thediveo/go-mntinfo"

func (si *SysInfo) getMountInfo() {
	si.Mount = mntinfo.Mounts()
}
