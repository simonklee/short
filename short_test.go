package short

import "testing"

var (
	bufThousand   = Encode(1000)
	bufMillion    = Encode(1000000)
	buf10Million  = Encode(10000000)
	buf100Million = Encode(100000000)
)

func TestEncodeDecode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := Decode(Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
}

func BenchmarkEncodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1000)
	}
}

func BenchmarkEncodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1000000)
	}
}

func BenchmarkEncode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(10000000)
	}
}

func BenchmarkEncode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(100000000)
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
