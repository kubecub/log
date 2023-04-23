// Copyright © 2023 KubeCub. All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package cuslog

type Formatter interface {
	// Maybe in async goroutine
	// Please write the result to buffer
	Format(entry *Entry) error
}