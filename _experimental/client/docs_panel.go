// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	// "io"
	// "net/http"
	// "regexp"

	//	gpuatom "android.googlesource.com/platform/tools/gpu/atom"
	"github.com/google/gxui"
	//	"golang.org/x/net/html"
	//	"golang.org/x/net/html/atom"
)

func CreateDocsPanel(appCtx *ApplicationContext) gxui.Control {
	theme := appCtx.Theme()

	ll := theme.CreateLinearLayout()
	appCtx.OnAtomSelected(func() {
		/*
			ll.RemoveAll()
			a := appCtx.Atoms()[appCtx.SelectedAtomID()]
			docs := gpuatom.MetadataOf(a).DocumentationUrl
			loadHtml(theme, docs, ll)
		*/
	})
	sl := theme.CreateScrollLayout()
	sl.SetChild(ll)
	return sl
}

/*
var duplicateWhitespaceRegExp = regexp.MustCompile("[ \t\r\n\f]+")

func parseHtml(r io.Reader, t gxui.Theme, c gxui.Container) {
	d := html.NewTokenizer(r)
	str := ""
	newline := func() {
		s := duplicateWhitespaceRegExp.ReplaceAllString(str, " ")
		str = ""
		t.Driver().Call(func() {
			l := t.CreateLabel()
			l.SetText(s)
			c.AddChild(l)
		})
	}
	commit := func() {
		if len(str) > 0 {
			newline()
		}
	}
	inBody := false
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			switch token.DataAtom {
			case atom.Body:
				inBody = true
			case atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6:
				commit()
			}
		case html.EndTagToken:
			switch token.DataAtom {
			case atom.Body:
				inBody = false
			case atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6,
				atom.P,
				atom.Tr,
				atom.Dl, atom.Dt,
				atom.Ol, atom.Li:
				commit()
			}
		case html.TextToken:
			if inBody {
				str += html.UnescapeString(token.String())
			}
		case html.SelfClosingTagToken:
			if token.DataAtom == atom.Br {
				newline()
			}
		}
	}
}

func loadHtml(t gxui.Theme, url string, c gxui.Container) {
	go func() {
		if resp, err := http.Get(url); err == nil {
			parseHtml(resp.Body, t, c)
		}
	}()
}
*/
