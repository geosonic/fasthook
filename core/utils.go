/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import "strconv"

func toArray(slice []int) string {
	var s string

	for i := 0; i < len(slice); i++ {
		if i > 0 {
			s += ","
		}
		s += strconv.Itoa(slice[i])
	}
	return s
}
