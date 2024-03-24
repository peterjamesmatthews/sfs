package mem

import (
	"reflect"
	"testing"
)

func TestGetNames(t *testing.T) {
	tests := []struct {
		name string   // test name
		uri  URI      // uri to call GetNames on
		want []string // uri names
	}{
		{
			name: "root uri",
			uri:  RootURI,
			want: []string{RootName},
		},
		{
			name: "/foo/bar",
			uri:  URI{uri: "/foo/bar"},
			want: []string{RootName, "foo", "bar"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.uri.GetNames()

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestAddNames(t *testing.T) {
	tests := []struct {
		name  string
		uri   URI
		names []string
		want  URI
	}{
		{
			name:  "root uri adds nothing",
			uri:   RootURI,
			names: []string{},
			want:  RootURI,
		},
		{
			name:  "root uri add foo",
			uri:   RootURI,
			names: []string{"foo"},
			want:  URI{uri: "/foo"},
		},
		{
			name:  "root uri adds lots of names",
			uri:   RootURI,
			names: []string{"foo", "bar", "baz"},
			want:  URI{uri: "/foo/bar/baz"},
		},
		{
			name:  "absolute uri adds lots of names",
			uri:   URI{uri: "/foo"},
			names: []string{"bar", "baz", "oob"},
			want:  URI{uri: "/foo/bar/baz/oob"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.uri.AddNames(test.names...)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
