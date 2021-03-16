package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	configYaml := []byte(`
a: a
b:
  ba:
    baa:
    - foo
    - bar
  bb: [1, 2]
`)
	got := mockConfig{}
	want := mockConfig{
		A: "a",
		B: mockConfigB{
			BA: mockConfigBA{
				BAA: []string{"foo", "bar"},
			},
			BB: []int{1, 2},
		},
	}

	err := Unmarshal(&got, RawConfigOption(configYaml))
	if err != nil {
		t.Errorf("failed to parse config: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v", want, got)
	}
}

func TestBindEnvs(t *testing.T) {
	gots := []interface{}{
		&mockConfig{A: "foobar"},
		mockConfig{B: mockConfigB{
			BB: []int{1, 2, 3}},
		},
		&mockConfig{},
	}

	want := []string{
		"a",
		"b.ba.baa",
		"b.bb",
	}

	for _, got := range gots {
		ifp := &mockIfaceProccess{}

		if err := traverseIface(ifp, got); err != nil {
			t.Errorf("failed to traverse: %v", err)
		}

		if !reflect.DeepEqual(want, ifp.binds) {
			t.Errorf("wanted binds %v got %v", want, ifp.binds)
		}
	}
}

type mockConfig struct {
	A string      `mapstructure:"a"`
	B mockConfigB `mapstructure:"b"`
}

type mockConfigBA struct {
	BAA []string `mapstructure:"baa"`
}

type mockConfigB struct {
	BA mockConfigBA `mapstructure:"ba"`
	BB []int        `mapstructure:"bb"`
}

type mockIfaceProccess struct {
	binds []string
}

func (i *mockIfaceProccess) process(mapKey string) error {
	i.binds = append(i.binds, mapKey)

	return nil
}
