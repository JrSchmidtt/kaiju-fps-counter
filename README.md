# FPS Counter Plugin

A simple plugin for the Kaiju Engine that displays the current FPS in the editor window title.

## About

This plugin shows the current frames per second (FPS) in the editor window title, making it easy to monitor performance while working in the Kaiju Engine editor.

**Version:** 0.002  
**Author:** JrSchmidtt

## Kaiju Engine

This plugin is designed for the [Kaiju Engine](https://github.com/KaijuEngine/kaiju).

## Installation

### Option 1: Clone directly to plugins folder

Navigate to your Kaiju Engine plugins directory and clone this repository:

```bash
cd KaijuEngine/Editor/plugins/
git clone https://github.com/JrSchmidtt/kaiju-fps-counter.git fpscounter
```

### Option 2: Manual installation

Place this plugin folder in the `plugins` directory of your Kaiju Engine installation:

```
KaijuEngine/Editor/plugins/fpscounter/
```

The plugin folder should contain:
- `plugin.json` - Plugin configuration file
- `plugin.go` - Plugin source code

After placing the files, the plugin will be automatically recognized by the engine. You can enable/disable it through the plugin manager.

For more information about Kaiju Engine plugins, visit: https://kaijuengine.com/editor/plugins/