package idpool

import "testing"

func EXPECT(cond bool, t *testing.T, msg string, args ...any) {
	if !cond {
		t.Fatalf(msg, args...)
	}
}

func TestIdpoolSetStart(t *testing.T) {
	p := Idpool{}
	p.SetStart(10)

	n := p.Next()

	EXPECT(n == 10, t, "next: %d", n)
}

func TestIdpoolNext(t *testing.T) {
	p := Idpool{}

	n1 := p.Next()
	n2 := p.Next()

	EXPECT(n2 == (n1+1), t, "next increment: %d %d", n1, n2)
}

func TestIdpoolRelease(t *testing.T) {
	p := Idpool{}

	_ = p.Next()
	n2 := p.Next()
	_ = p.Next()

	p.Release(n2)
	n4 := p.Next()

	EXPECT(n4 == n2, t, "released: %d next: %d", n2, n4)
}

func TestIdpoolReleaseUnknown(t *testing.T) {
	p := Idpool{}

	p.Release(3)

	numReleased := len(p.released)
	EXPECT(numReleased == 0, t, "num released: %d", numReleased)
}
