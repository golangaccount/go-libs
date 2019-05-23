package config

import (
	"errors"
	"os"
	"sync"

	lclog "git.lichengsoft.com/lichengsoft/go-libs/log"
	"github.com/Unknwon/goconfig"
)

var config *goconfig.ConfigFile                     //config文件
var setError = errors.New("goconfig setting error") //配置文件存储失败
var saveConfigLock sync.RWMutex                     //save config 锁

const (
	Config_Path = "./conf/app.conf" //默认的配置文件地址
)

func init() {
	var err error
	config, err = goconfig.LoadConfigFile(Config_Path) //默认由于配置 自动加载默认地址的配置文件
	if err != nil {
		lclog.Fatal("加载配置文件失败,错误信息:%s", err.Error())
		os.Exit(0)
	}
}

/*返回默认配置节点*/
func DEFAULTSECTION() string {
	return goconfig.DEFAULT_SECTION
}

// GetValue returns the value of key available in the given section.
// If the value needs to be unfolded
// (see e.g. %(google)s example in the GoConfig_test.go),
// then String does this unfolding automatically, up to
// _DEPTH_VALUES number of iterations.
// It returns an error and empty string value if the section does not exist,
// or key does not exist in DEFAULT and current sections.
func GetValue(section, key string) (string, error) {
	return config.GetValue(section, key)
}

// Bool returns bool type value.
func Bool(section, key string) (bool, error) {
	return config.Bool(section, key)
}

// Float64 returns float64 type value.
func Float64(section, key string) (float64, error) {
	return config.Float64(section, key)
}

// Int returns int type value.
func Int(section, key string) (int, error) {
	return config.Int(section, key)
}

// Int64 returns int64 type value.
func Int64(section, key string) (int64, error) {
	return config.Int64(section, key)
}

// MustValue always returns value without error.
// It returns empty string if error occurs, or the default value if given.
func MustValue(section, key string, defaultVal ...string) string {
	return config.MustValue(section, key, defaultVal...)
}

// MustValueRange always returns value without error,
// it returns default value if error occurs or doesn't fit into range.
func MustValueRange(section, key, defaultVal string, candidates []string) string {
	return config.MustValueRange(section, key, defaultVal, candidates)
}

// MustValueArray always returns value array without error,
// it returns empty array if error occurs, split by delimiter otherwise.
func MustValueArray(section, key, delim string) []string {
	return config.MustValueArray(section, key, delim)
}

// MustBool always returns value without error,
// it returns false if error occurs.
func MustBool(section, key string, defaultVal ...bool) bool {
	return config.MustBool(section, key, defaultVal...)
}

// MustFloat64 always returns value without error,
// it returns 0.0 if error occurs.
func MustFloat64(section, key string, defaultVal ...float64) float64 {
	return config.MustFloat64(section, key, defaultVal...)
}

// MustInt always returns value without error,
// it returns 0 if error occurs.
func MustInt(section, key string, defaultVal ...int) int {
	return config.MustInt(section, key, defaultVal...)
}

// MustInt64 always returns value without error,
// it returns 0 if error occurs.
func MustInt64(section, key string, defaultVal ...int64) int64 {
	return config.MustInt64(section, key, defaultVal...)
}

// GetSectionList returns the list of all sections
// in the same order in the file.
func GetSectionList() []string {
	return config.GetSectionList()
}

// GetKeyList returns the list of all keys in give section
// in the same order in the file.
// It returns nil if given section does not exist.
func GetKeyList(section string) []string {
	return config.GetKeyList(section)
}

// GetSection returns key-value pairs in given section.
// It section does not exist, returns nil and error.
func GetSection(section string) (map[string]string, error) {
	return config.GetSection(section)
}

// GetSectionComments returns the comments in the given section.
// It returns an empty string(0 length) if the comments do not exist.
func GetSectionComments(section string) (comments string) {
	return config.GetSectionComments(section)
}

// GetKeyComments returns the comments of key in the given section.
// It returns an empty string(0 length) if the comments do not exist.
func GetKeyComments(section, key string) (comments string) {
	return config.GetKeyComments(section, key)
}

func SetValue(section, key, value string) (bool, error) {
	result := config.SetValue(section, key, value)
	if !result {
		return result, setError
	}
	return result, saveConfig()
}

func DeleteKey(section, key string) (bool, error) {
	result := config.DeleteKey(section, key)
	if !result {
		return result, setError
	}
	return result, saveConfig()
}

// MustValue always returns value without error,
// It returns empty string if error occurs, or the default value if given,
// and a bool value indicates whether default value is returned.
func MustValueSet(section, key string, defaultVal ...string) (string, bool, error) {
	str, result := config.MustValueSet(section, key, defaultVal...)
	if !result {
		return str, result, setError
	}
	return str, result, saveConfig()
}

// DeleteSection deletes the entire section by given name.
// It returns true if the section was deleted, and false if the section didn't exist.
func DeleteSection(section string) (bool, error) {
	result := config.DeleteSection(section)
	if !result {
		return result, setError
	}
	return result, saveConfig()
}

// SetSectionComments adds new section comments to the configuration.
// If comments are empty(0 length), it will remove its section comments!
// It returns true if the comments were inserted or removed,
// or returns false if the comments were overwritten.
func SetSectionComments(section, comments string) (bool, error) {
	result := config.SetSectionComments(section, comments)
	if !result {
		return result, setError
	}
	return result, saveConfig()
}

// SetKeyComments adds new section-key comments to the configuration.
// If comments are empty(0 length), it will remove its section-key comments!
// It returns true if the comments were inserted or removed,
// or returns false if the comments were overwritten.
// If the section does not exist in advance, it is created.
func SetKeyComments(section, key, comments string) (bool, error) {
	result := config.SetKeyComments(section, key, comments)
	if !result {
		return result, setError
	}
	return result, saveConfig()
}

func saveConfig() error {
	saveConfigLock.Lock()
	defer saveConfigLock.Unlock()
	return goconfig.SaveConfigFile(config, Config_Path)
}
