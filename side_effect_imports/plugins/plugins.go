package plugins

import (
		"io/ioutil"
	"fmt"
	"github.com/prometheus/common/log"
	"regexp"
	"plugin"
	. "../plugin_definitions"
)

var plugins = map[string]PluginEntrypoint{}

func init() {
	log.Infoln("Initializing plugins...")

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Errorln("Could not load plugins!")
		return
	}

	fileRegexp := regexp.MustCompile(`\.(so|dll)`)

	pluginsLoaded := false

	for _, f := range files {
		if fileRegexp.MatchString(f.Name()) {
			ok := loadPlugin(f.Name())
			pluginsLoaded = pluginsLoaded || ok
		}
	}

	if !pluginsLoaded {
		log.Infoln("No plugins were found/loaded!")
	}

}

func loadPlugin(file string) bool {
	p, err := plugin.Open(file)
	if err != nil {
		log.Errorln("Could not load plugin " + file)
		return false
	}

	peSym, err := p.Lookup("Entrypoint")
	if err != nil {
		log.Errorln("Could not load plugin " + file)
		return false
	}
	var pe PluginEntrypoint
	pe, ok := peSym.(PluginEntrypoint)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		return false
	}

	infoRaw, err := pe("info", nil)
	if err != nil {
		fmt.Println("plugin does not implement info endpoint")
		return false
	}

	info, ok := infoRaw.(PluginInfo)
	if !ok {
		fmt.Println("returned plugin info is garbage")
		return false
	}

	plugins[info.Name] = pe
	return true
}

func Run(name string, payload interface{}) (interface{}, error) {
	for pluginName, plugin := range plugins {
		data, err := plugin("function_"+name, payload)
		if err == ENoFunction {
			log.Debugf("Plugin '%s' does not provide '%s'", pluginName, name)
			continue
		}
		return data, err
	}
	return nil, ENoFunction
}
