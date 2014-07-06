package short

import "testing"

var (
	bufThousand   = Encode(1e3)
	bufMillion    = Encode(1e6)
	buf10Million  = Encode(1e7)
	buf100Million = Encode(1e8)
)

func TestEncodeDecode(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		got := Decode(Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
}

func BenchmarkEncodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1e3)
	}
}

func BenchmarkEncodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1e6)
	}
}

func BenchmarkEncode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1e7)
	}
}

func BenchmarkEncode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1e8)
	}
}

func BenchmarkDecodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(bufThousand)
	}
}

func BenchmarkDecodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(bufMillion)
	}
}

func BenchmarkDecode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(buf10Million)
	}
}

func BenchmarkDecode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(buf100Million)
	}
}
