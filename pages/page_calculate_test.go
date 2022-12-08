package pages

import "testing"

func TestToDbQuery(t *testing.T) {
	type args struct {
		currentPage int
		pageSize    int
	}
	tests := []struct {
		name       string
		args       args
		wantLimit  int
		wantOffset int
	}{
		{
			name: "正常-第一页",
			args: args{
				currentPage: 1,
				pageSize:    10,
			},
			wantLimit:  10,
			wantOffset: 0,
		}, {
			name: "正常-第十页",
			args: args{
				currentPage: 10,
				pageSize:    10,
			},
			wantLimit:  10,
			wantOffset: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit, gotOffset := ToDbQuery(tt.args.currentPage, tt.args.pageSize)
			if gotLimit != tt.wantLimit {
				t.Errorf("ToDbQuery() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("ToDbQuery() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}
