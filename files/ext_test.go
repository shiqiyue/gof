package files

import "testing"

func TestExt(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a.txt test",
			args: args{fileName: "a.txt"},
			want: ".txt",
		},
		{
			name: "a.jpg test",
			args: args{fileName: "a.jpg"},
			want: ".jpg",
		},
		{
			name: "http url test",
			args: args{fileName: "https://codeup.aliyun.com/5fbdaa372f8cc15c287b5938/shiqiyue/jet_show_image_download/blob/master/main.go"},
			want: ".go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.args.fileName); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
