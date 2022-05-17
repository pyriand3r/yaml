package yaml

import (
	"os"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Parse tries to parses a yaml file into the given struct. Given struct must be a pointer.
func Parse(path string, cfg interface{}) error {
	s := reflect.ValueOf(cfg)
	if s.Kind() != reflect.Ptr {
		return errors.New("given config value is not a pointer to a struct")
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return errors.New("given config value is not a pointer to a struct")
	}

	file, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "could not open %s", path)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return errors.Wrap(err, "could not decode yaml file to config struct")
	}

	return nil
}
