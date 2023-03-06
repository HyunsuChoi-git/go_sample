package main

import "fmt"

type shoppingBasket struct {
	cnt   int
	price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본을 수정하는 포인터 형식
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본수정 없이 값 전달만
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {

	// reciever == 구조체 메소드.
	// 일반 메소드는 기본이 값 전달 방식이고 slice, map등의 객체의 경우에는 참조 전달 방식
	// 리시버도 포인터를 활요하여 참조전달 방식으로 메소드를 구현할 수 있다.
	//  - 값 : 일반 함수와 마찬자기로 값 복사 후 함수호출 시 파라미터로 전달 (Default)
	//  - 참조 : 포인터 형식으로 메소드를 정의하면 알아서 포인터형으로 파라미터가 전달됨

	// 예제 1
	bs1 := shoppingBasket{cnt: 3, price: 5000}
	fmt.Println("ex1 : ", bs1.purchase())
	fmt.Println(" -------------- ")

	bs1.rePurchaseP(2, 10000)
	fmt.Println("ex2 : ", bs1.cnt, bs1.price)
	fmt.Println("ex2 : ", bs1.purchase())
	fmt.Println(" -------------- ")

	bs1.rePurchaseD(2, 10000)
	fmt.Println("ex3 : ", bs1.cnt, bs1.price)
	fmt.Println("ex3 : ", bs1.purchase())

}
