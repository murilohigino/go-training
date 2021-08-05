package main

import "testing"

func Test_achepercentual(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := achepercentual(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("achepercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addpercentual(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addpercentual(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("addpercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mult(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mult(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_potencia(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := potencia(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("potencia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtracao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "teste 4 - 1",
			args: args{4 , 1},
			want: 3,
		},
		{
			name: "teste 4 - 2",
			args: args{4 , 2},
			want: 2,
		},
		{
			name: "teste 4 - 3",
			args: args{4 , 3},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*{
	name: "teste 4 - 1",
	args: args{4 , 1},
	want: 3,
},
{
	name: "teste 4 - 2",
	args: args{4 , 2},
	want: 2,
},
{
	name: "teste 4 - 3",
	args: args{4 , 3},
	want: 1,
},

 */
