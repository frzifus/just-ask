// BSD 2-Clause License
//
// Copyright (c) 2018, frzifus
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"image"
	"os"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/style"
)

const (
	scaling = 1.0
)

var (
	theme  = style.DarkTheme
	window nucular.MasterWindow
	size   = image.Point{300, 200}
	input  = make([]string, 0)
)

func id(fn func(*nucular.Window)) func() func(*nucular.Window) {
	return func() func(*nucular.Window) {
		return fn
	}
}

type popup struct {
	Name     string
	Title    string
	Flags    nucular.WindowFlags
	UpdateFn func() func(*nucular.Window)
}

func askPass(w *nucular.Window) {
	scale := 8
	for i := 0; i < scale; i++ {
		w.Row((120) / scale).Dynamic(1)
		if i < len(input) {
			w.Label(input[i], "CT")
		} else {
			w.Label("", "CT")
		}
	}

	w.Row(20).Static(size.X/2-80, 60, 60)
	w.Label("", "CL")
	if w.Button(label.TA("OK", "CB"), false) {
		os.Exit(0)
	}
	if w.Button(label.TA("Cancel", "CB"), false) {
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	var argStr string
	if len(args) > 1 {
		argStr = args[len(args)-1]
	} else {
		argStr = "Text"
	}

	var res string

	for i, c := range argStr {
		res = res + string(c)
		if i > 0 && (i+1)%40 == 0 {
			input = append(input, res)
			res = ""
		}
	}
	if len(res) > 0 {
		input = append(input, res)
	}

	p := popup{"AskPass", "AskPass", 0, id(askPass)}

	window = nucular.NewMasterWindowSize(p.Flags, p.Title, size, p.UpdateFn())
	window.SetStyle(style.FromTheme(theme, scaling))
	window.Main()
}
