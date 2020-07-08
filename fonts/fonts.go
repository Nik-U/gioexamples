// SPDX-License-Identifier: Unlicense OR MIT

package main

// A Gio program that demonstrates multiple fonts. See https://gioui.org for more information.

import (
	"fmt"
	"log"
	"os"

	"github.com/gonoto/notosans"

	"eliasnaur.com/font/roboto/robotoregular"

	"gioui.org/font/opentype"
	"gioui.org/unit"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		defer os.Exit(0)
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(700)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	fonts := gofont.Collection()
	// It is also possible to register other styles by setting other fields of text.Font.
	// This demo uses only the regular style for Roboto and Noto to minimize file size.
	fonts = appendTTF(fonts, text.Font{Typeface: "Roboto"}, robotoregular.TTF)
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())

	th := material.NewTheme(fonts)

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			render(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

var (
	helloList = &layout.List{Axis: layout.Vertical}
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func render(gtx C, th *material.Theme) D {
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				l := material.H2(th, "Fonts")
				l.Font = text.Font{Weight: text.Bold}
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(20)}.Layout(gtx,
					material.H6(th, "Go: a font family made specifically for the Go Project").Layout,
				)
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(20)}.Layout(gtx, func(gtx C) D {
					l := material.H6(th, "Roboto: the default font family on Android and Chrome OS")
					l.Font = text.Font{Typeface: "Roboto"}
					return l.Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(20)}.Layout(gtx, func(gtx C) D {
					l := material.H6(th, "Noto: a font family with broad unicode coverage:")
					l.Font = text.Font{Typeface: "Noto"}
					return l.Layout(gtx)
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return layout.UniformInset(unit.Sp(20)).Layout(gtx, func(gtx C) D {
					return helloList.Layout(gtx, len(helloWorlds), func(gtx C, i int) D {
						l := material.Body1(th, helloWorlds[i])
						l.Font = text.Font{Typeface: "Noto"}
						return l.Layout(gtx)
					})
				})
			}),
		)
	})
}

func appendTTF(collection []text.FontFace, fnt text.Font, ttf []byte) []text.FontFace {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

var helloWorlds = []string{
	"• 👋🗺 🎉",
	"• Hello, world!",
	"• 你好世界！",
	"• नमस्ते दुनिया!",
	"• ¡Hola Mundo!",
	"• Bonjour monde!",
	"• مرحبا بالعالم!",
	"• ওহে বিশ্ব!",
	"• Привет мир!",
	"• Olá Mundo!",
	"• Halo Dunia!",
	"• ہیلو ، دنیا!",
	"• Hallo Welt!",
	"• こんにちは世界！",
	"• Salamu Dunia!",
	"• नमस्कार, जग!",
	"• హలో, ప్రపంచం!",
	"• Selam Dünya!",
	"• வணக்கம், உலகமே!",
	"• ਸਤਿ ਸ੍ਰੀ ਅਕਾਲ ਦੁਨਿਆ!",
	"• 안녕, 세상!",
	"• Chào thế giới!",
	"• Sannu Duniya!",
	"• Halo, jagad!",
	"• Ciao mondo!",
	"• สวัสดีชาวโลก!",
	"• હેલો, વિશ્વ!",
	"• ಹಲೋ, ಜಗತ್ತು!",
	"• سلام دنیا!",
	"• Kumusta, mundo!",
}
