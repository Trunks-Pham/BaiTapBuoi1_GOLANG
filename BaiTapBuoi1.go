package main

import (
	"fmt"
	"strings"
)

// 1. Đảo ngược chuỗi
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// 2. Đếm số lần xuất hiện của một ký tự trong chuỗi
func countCharacter(s string, char rune) int {
	count := 0
	for _, c := range s {
		if c == char {
			count++
		}
	}
	return count
}

// 3. Kiểm tra chuỗi đối xứng (Palindrome)
func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// 4. Loại bỏ khoảng trắng
func removeWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// 5. Tìm chuỗi con (substring)
func containsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 1. Kiểm tra số nguyên tố
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 2. Tính giai thừa của một số
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 3. Tìm số Fibonacci thứ n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 4. Kiểm tra số hoàn hảo
func isPerfectNumber(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if i == n/i {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}
	return sum == n && n != 1
}

// 5. Tính tổng các chữ số của một số
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// 1. Tìm giá trị lớn nhất và nhỏ nhất trong mảng
func findMaxAndMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

// 2. Sắp xếp mảng
func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 3. Tìm số lần xuất hiện của một phần tử trong mảng
func countOccurrences(arr []int, x int) int {
	count := 0
	for _, v := range arr {
		if v == x {
			count++
		}
	}
	return count
}

// 4. Xóa phần tử khỏi mảng
func removeElement(arr []int, x int) []int {
	result := []int{}
	for _, v := range arr {
		if v != x {
			result = append(result, v)
		}
	}
	return result
}

// 5. Gộp hai mảng thành một
func mergeArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

func main() {
	fmt.Println("\nBÀI TẬP - NGÀY 17/08/2024\n\n")
	// Chuỗi
	fmt.Println("BÀI TẬP - CHUỖI")
	fmt.Println("Đảo ngược chuỗi 'hello':", reverseString("hello"))
	fmt.Println("Đếm số lần xuất hiện của ký tự 'l' trong 'hello':", countCharacter("hello", 'l'))
	fmt.Println("Kiểm tra 'madam' có phải là chuỗi đối xứng:", isPalindrome("madam"))
	fmt.Println("Loại bỏ khoảng trắng trong 'hello world':", removeWhitespace("hello world"))
	fmt.Println("Kiểm tra 'hello' có chứa 'ell':", containsSubstring("hello", "ell"))

	// Số
	fmt.Println("\n\nBÀI TẬP - SỐ")
	fmt.Println("Kiểm tra 7 có phải là số nguyên tố:", isPrime(7))
	fmt.Println("Tính giai thừa của 5:", factorial(5))
	fmt.Println("Số Fibonacci thứ 7:", fibonacci(7))
	fmt.Println("Kiểm tra 6 có phải là số hoàn hảo:", isPerfectNumber(6))
	fmt.Println("Tính tổng các chữ số của 123:", sumOfDigits(123))

	// Mảng
	fmt.Println("\n\nBÀI TẬP - MảNG")
	arr := []int{3, 5, 7, 2, 8}
	max, min := findMaxAndMin(arr)
	fmt.Println("Giá trị lớn nhất và nhỏ nhất trong mảng [3, 5, 7, 2, 8]:", max, min)
	fmt.Println("Sắp xếp mảng [5, 3, 8, 4, 2]:", sortArray([]int{5, 3, 8, 4, 2}))
	fmt.Println("Số lần xuất hiện của 2 trong mảng:", countOccurrences([]int{1, 2, 3, 2, 2, 4}, 2))
	fmt.Println("Xóa phần tử 2 khỏi mảng:", removeElement([]int{1, 2, 3, 4}, 2))
	fmt.Println("Gộp hai mảng [1,3] và [2,4]:", mergeArrays([]int{1, 3}, []int{2, 4}))
}
