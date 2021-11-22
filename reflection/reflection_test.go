package reflection

import (
	"reflect"
	"testing"
)

//https://go.dev/blog/laws-of-reflection

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Country string
	Age     int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Stuart"},
			[]string{"Stuart"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Daniel", "Manila"},
			[]string{"Daniel", "Manila"},
		},
		{
			"Struct with one string and one int field",
			struct {
				Name string
				Age  int
			}{"Lucy", 5},
			[]string{"Lucy"},
		},
		{
			"Struct with nested struct",
			Person{
				Name: "Juan",
				Profile: Profile{
					Country: "Lesotho",
					Age:     1,
				},
			},
			[]string{"Juan", "Lesotho"},
		},
		{
			"Pointers to things",
			&Person{
				Name: "Lisa",
				Profile: Profile{
					Country: "Japan",
					Age:     1,
				},
			},
			[]string{"Lisa", "Japan"},
		},
		{
			"Slices",
			[]Profile{
				{"Canada", 12312},
				{"Russia", 123},
			},
			[]string{"Canada", "Russia"},
		},
		{
			"Arrays",
			[2]Profile{
				{"Jamaica", 12312},
				{"Denmark", 123},
			},
			[]string{"Jamaica", "Denmark"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		myMap := map[string]string{
			"Yes": "No",
			"Foo": "Bar",
		}

		var got []string
		walk(myMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "No")
		assertContains(t, got, "Bar")
	})

	t.Run("channesl", func(t *testing.T) {
		myChann := make(chan Profile)

		go func() {
			myChann <- Profile{"Japan", 123}
			myChann <- Profile{"Denmark", 123}
			close(myChann)
		}()

		var got []string
		want := []string{"Japan", "Denmark"}

		walk(myChann, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		myFunction := func() (Profile, Profile) {
			return Profile{"Japan", 123}, Profile{"Denmark", 123}
		}

		var got []string
		want := []string{"Japan", "Denmark"}

		walk(myFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t *testing.T, got []string, s string) {
	t.Helper()
	contains := false
	for _, v := range got {
		if v == s {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it did not", got, s)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
