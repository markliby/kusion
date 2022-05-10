package compile

import (
	"reflect"
	"strings"
	"testing"

	kcl "kusionstack.io/kclvm-go"
)

func TestCompileResult_YAMLString(t *testing.T) {
	type fields struct {
		Documents []kcl.KCLResult
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "t1",
			fields: fields{
				Documents: []kcl.KCLResult{{"a": "b"}},
			},
			want: "a: b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompileResult{
				Documents: tt.fields.Documents,
			}
			if got := strings.TrimSpace(c.YAMLString()); got != tt.want {
				t.Errorf("CompileResult.YAMLString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCompileResultByMapList(t *testing.T) {
	type args struct {
		mapList []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *CompileResult
	}{
		{
			name: "t1",
			args: args{
				mapList: []map[string]interface{}{{"replicas": 1}},
			},
			want: &CompileResult{
				Documents: []kcl.KCLResult{map[string]interface{}{"replicas": 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompileResultByMapList(tt.args.mapList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompileResultByMapList() = %v, want %v", got, tt.want)
			}
		})
	}
}
