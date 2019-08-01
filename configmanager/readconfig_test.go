package configmanager

import (
	"reflect"
	"testing"
)

func typeIsExpected(actualType interface{}, expectedString string) bool {
	actualString := reflect.TypeOf(actualType).String()

	if actualString == expectedString {
		return true
	}

	return false
}

func TestConfigFileExists(t *testing.T) {
	exists := ConfigFileExists()

	if !exists {
		t.Error("config file doesn't exists")
	}
}

func TestLoadConfigFile(t *testing.T) {
	result := LoadConfigFile()
	expected := "[]uint8"

	typeEx := typeIsExpected(result, expected)

	if !typeEx {
		message := "file configuration loading error"
		t.Error(message)
	}
}

func TestOpenConfigFile(t *testing.T) {
	result := OpenConfigFile()
	expected := "*os.File"

	typeEx := typeIsExpected(result, expected)

	if !typeEx {
		message := "open file config error"
		t.Error(message)
	}
}

func TestReadConfigFile(t *testing.T) {
	jsonboject := ReadConfig()

	expectedVal := jsonboject.Name

	if expectedVal != "config" {
		message := "Configformat error"
		t.Error(message)
	}
}
