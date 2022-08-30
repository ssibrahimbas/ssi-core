package config

import "testing"

func TestConfigNewPositive(t *testing.T) {
	_, err := New("./test.env")
	if err != nil {
		t.Errorf("New() error = %v", err)
	}
}

func TestConfigNewNegative(t *testing.T) {
	_, err := New("./testttttt.env")
	if err == nil {
		t.Error("New() error = nil")
	}
}

func TestConfig_DecodePositive(t *testing.T) {
	type TestConfig struct {
		Label string `env:"LABEL"`
	}
	cnf := &TestConfig{}
	c, err := New("./test.env")
	if err != nil {
		t.Errorf("New() error = %v", err)
	}
	err = c.Decode(cnf)
	if err != nil {
		t.Errorf("New() error = %v", err)
	}
}

func TestConfig_DecodeNegative(t *testing.T) {
	c, err := New("./test.env")
	if err != nil {
		t.Errorf("New() error = %v", err)
	}
	err = c.Decode(0)
	if err == nil {
		t.Error("Decode() error = nil")
	}
}
