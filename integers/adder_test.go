package integers

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}
