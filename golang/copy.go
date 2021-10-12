package main

import "fmt"

func main() {
	a := []byte{'a'}
	var b, c []byte

	fmt.Printf("a: %s, b: %s, c: %s \n", a, b, c)

	b = a
	copy(c, a)
	// a = append(a,'1')
	fmt.Printf("a: %s, b: %s, c %s\n", a, b, c)

	a[0] = '1'
	fmt.Printf("a: %s, b: %s, c %s", a, b, c)

	d := []byte{'c', 'c', 'c', 'c', 'c', 'c'}
	i := copy(d, a)
	j := make([]byte, len(d))
	copy(j, d)
	fmt.Printf("a: %s, d: %s,j: %s\n", a, d[:i], j)

	h := d
	fmt.Printf("before %p\n", h)
	fmt.Println(cap(h), len(h))
	h = h[:4]
	fmt.Printf("after %p\n", h)
	fmt.Println(cap(h), len(h))

	copy(h, d)
	fmt.Printf("after %p\n", h)
	fmt.Println(cap(h), len(h), string(h))
	h = h[:cap(h)]
	fmt.Println(cap(h), len(h), string(h))

	a1 := []byte{'a', 'a', 'a'}
	b1 := []byte{'b', 'b', 'b', 'b'}

	b1 = a1
	fmt.Printf("%s %s\n", a1, b1)

	b1[0] = '1'
	fmt.Printf("%s %s\n", a1, b1)

	s := "ssssss"
	s1 := s

	fmt.Println(*s, *s1)
}
