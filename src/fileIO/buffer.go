package main

func main() {
	// buffer

	// 파일 읽기, 버퍼를 사용하여 파일 쓰기. bufio Package
	// ioutil, bufio 등은 Io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, writer, Read 메소드를 사용 가능

	// bufio package
	// https://golang.org/pkg/bufio

	// 두번째 매개변수
	// https://golang.org/pkg/os/#pkg-constants

	/*
		a -----> a
		b -----> ab
		c -----> abc
		d -----> abcd
		e -----> e ----->abcd
		f -----> ef -----> abcd
		g -----> efg -----> abcd
		h -----> efgh -----> abcd
		i -----> i -----> abcdefgh
		...
	*/

	// 예제 1

}
