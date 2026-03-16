package fpscounter

import (
	"fmt"
	"log/slog"

	"kaijuengine.com/editor"
	"kaijuengine.com/editor/editor_plugin"
)

const (
	pluginKey      = "github.com/schmidt/kaiju-editor-fps-counter"
	sampleInterval = 0.5
	titlePrefix    = "Kaiju Editor"
)

type Plugin struct{}

func init() { editor.RegisterPlugin(pluginKey, &Plugin{}) }

func (p *Plugin) Launch(ed editor_plugin.EditorInterface) error {
	slog.Info("launching FPS Counter plugin", "key", pluginKey)
	slog.Warn("FPS Counter plugin loaded")
	host := ed.Host()
	host.Window.SetTitle(titlePrefix + " | FPS plugin loaded")

	elapsed := 0.0
	frames := 0

	host.Updater.AddUpdate(func(delta float64) {
		elapsed += delta
		frames++

		if elapsed < sampleInterval {
			return
		}

		fps := float64(frames) / elapsed
		host.Window.SetTitle(fmt.Sprintf("%s | FPS: %.1f", titlePrefix, fps))

		elapsed = 0
		frames = 0
	})

	return nil
}
