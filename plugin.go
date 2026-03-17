package fpscounter

import (
	"fmt"
	"kaijuengine.com/editor"
	"kaijuengine.com/editor/editor_plugin"
)

const (
	pluginKey      = "github.com/JrSchmidtt/kaiju-fps-counter/"
	sampleInterval = 1.0
	titlePrefix    = "Kaiju Editor"
)

type Plugin struct{}

func init() { editor.RegisterPlugin(pluginKey, &Plugin{}) }

func (p *Plugin) Launch(ed editor_plugin.EditorInterface) error {
	host := ed.Host()

	host.RunAfterFrames(5, func() {
		elapsed := 0.0

		host.Updater.AddUpdate(func(delta float64) {
			elapsed += delta

			if elapsed < sampleInterval {
				return
			}

			fps := 1.0 / delta
			host.RunOnMainThread(func() {
				host.Window.SetTitle(fmt.Sprintf("%s | FPS: %.1f", titlePrefix, fps))
			})

			elapsed = 0
		})
	})

	return nil
}
