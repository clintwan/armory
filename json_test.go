package armory

import (
	"reflect"
	"testing"
)

func Test_j_MarshalWithoutEscapeHTML(t *testing.T) {
	type args struct {
		v      interface{}
		pretty bool
	}
	tests := []struct {
		name    string
		s       *j
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &j{}
			got, err := s.MarshalWithoutEscapeHTML(tt.args.v, tt.args.pretty)
			if (err != nil) != tt.wantErr {
				t.Errorf("j.MarshalWithoutEscapeHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("j.MarshalWithoutEscapeHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMarshalWithoutEscapeHTML(t *testing.T) {
	v := map[string]string{
		"a1": "<div></div>",
		"a2": ">1",
		"a3": "!=2",
		"a4": "<3",
	}
	buf, err := Json.MarshalWithoutEscapeHTML(&v, true)
	t.Error(string(buf))
	t.Error(err)
}

func TestIndent(t *testing.T) {
	s := `{"a1":1,"a2":"<div></div>","a3":3,"a4":4}`
	buf, err := Json.Indent([]byte(s))
	t.Error(string(buf))
	t.Error(err)
}
