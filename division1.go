// integer division program
// calculates a/b to arbitrary precision
package main

import "fmt"
import "strconv"
import "reflect"

func main() {
	var a,b,n1,n2 int
	var string1 string

	a1 := reflect.TypeOf(a).Size()
	fmt.Print("\n The size of int type is ",a1," bytes \n\n")

// Example 1
	a = 4342321
	b = 13213127
	n1 = 100
	string1 = divide(a,b,n1)
	fmt.Println("\n",a," / ",b," to ",n1," decimial digits of precision is ")
	fmt.Print(string1,"\n\n")


// Example 2
	a = 1231237
	b = 42312319
	n1 = 5
	n2 = 25
	string1 = digits(a,b,n1,n2)
	fmt.Print("The ",n1," to ",n2," digits in the expansion of ",a)
	fmt.Print(" / ",b," is ",string1,"\n\n")

	}


func abs(a int) int {
	var t int
	t = a
	if a<0{
		t = -a
	}
	return t
}


func exp1(e int, g int, n int) int {
	var t,sq,e1 int
	var t1 bool

	t = 1
	sq = g
	e1 = e
	for t1 = true; t1; t1 = (e1!=0){
		if(e1%2==1){
			t = (sq*t)%n
			e1 = (e1-1)/2
		}	else{
				e1 = e1/2
			}
		sq = (sq*sq)%n
	}
	result := t
	return result
}

func is_prime(p int) bool {
	var t1,t,i,m,y int
	var b bool

	p = abs(p)
	if (p==0){
		p = 1
	}
	m = 50
	y = 0
	b = false
	for i=2; i<m; i++ {
		t1 = i%p
		t = exp1((p-1)/2,i,p)
		if ((t==1)||(t==(p-1))||(t1==0)){
			y = y+1
		}
	}
	if (y==(m-2)){
		b = true
	}
	if (p==1){
		b = false
	}
	return b
}

func findprime(c, m int) int {
	var t bool
	var i,t1 int

	if c<1{
		c = 1
	}
	if m<1{
		m = 1
	}
	t = false
	if c%2==1{
		c = 2*c
		m = (m+m%2)/2
	}
	i = m+(m+1)%2
	t1 = c*i+1
	for t=true; t; t=!(is_prime(t1)){
		t1 = c*i+1
		i = i+2
	}
	return t1
}


func gcd(a, b int) int {
	var s int
	a = abs(a)
	b = abs(b)
	if a==0||b==0{
		s = a+b
		return s
	}
	if a==1||b==1{
		return 1
	}
	s = gcd(b%a,a)
	return s
}

func inverse2(a, b int) int {
	var t int
	if a==0||b==0{
		return 0
	}
	if a==1{
		return 1
	}
	t = b-((b*inverse2(b%a,a))/a)
	return t

}

func f10v3(n, a, p, q int) int {
	var p1,g,a1,n1,t,t1,t2,s int
	p1 = p*q
	g = 10
	a1 = inverse2(a,p1)
	n1 = abs(n)
	t = exp1(n1,g,p1)
	if n<0{
		t = inverse2(t,p1)
	}
	t1 = (a1*t)%p
	t2 = inverse2(t1,p)
	s = t2
	s = (s*a)%p1
	s = (s*t)%p1
	s = (s+(p-1))/p
	s = s%q
	return s
}


func f10v4(n, a, p int) int{
	var q,q1,s1,s2,s,g,n1,a1,i int
	t := [11]int {11,13,17,19,23,29,31,37,41,43,47}
	n = -n
	q = p
	for i=0; i<11; i++ {
		q1 = t[i]
		if a%q1!=0{
			q = q1
			break
		}
	}
	if p<q{
		q = p
	}
	if q<11{
		q = 11
	}
	s1 = f10v3(n,a,p,q)
	s2 = f10v3(n+1,a,p,q)
	s1 = (q-s1)%q
	s = (s1+s2)%q
	n1 = abs(n)
	g = 10
	if n>0{
		g = inverse2(10,q)
	}
	s1 = exp1(n1,g,q)
	s = (s*s1)%q
	a1 = inverse2(a,q)
	s = (a1*s)%q
	return s
}

func digits(a, b, n1, n2 int) string{
	var t,i,s int
	var flag byte
	var string1 string

	t = gcd(a,b)
	if t>1{
		a = a/t
		b = b/t
	}
	flag = 0
	if (b!=0) && (b%2==0 || b%5==0){
		flag = 2
	}
	if b==0{
		a = 0
		b = 1
		flag = 1
	}
	a = abs(a)
	b = abs(b)
	if a>b{
		a = a%b
	}
	string1 = ""
	for i = n1; i<(n2+1); i++{
		s = f10v4(i,a,b)
		string1 = string1 + strconv.Itoa(s)
	}
	if flag==1{
		string1 = "division by 0 error"
	}
	if flag==2{
		string1 = "error gcd(10,b/gcd(a,b)) > 1"
	}
	return string1
}

func divide(a, b, n1 int) string{
	var t,i,s int
	var flag byte
	var sign,string1,string2,string3 string

	t = gcd(a,b)
	if t>1{
		a = a/t
		b = b/t
	}
	flag = 0
	if (b!=0) && (b%2==0 || b%5==0){
		flag = 2
	}
	if b==0{
		a = 0
		b = 1
		flag = 1
	}
	sign = ""
	if a<0 && b>0{
		sign = "-"
	}
	if a>0 && b<0{
		sign = "-"
	}
	a = abs(a)
	b = abs(b)
	s = 0
	if a>b{
		s = a/b
		a = a%b
	}
	string1 = strconv.Itoa(s)
	string2 = ""
	for i=1; i<(n1+1); i++{
		s = f10v4(i,a,b)
		string2 = string2 + strconv.Itoa(s)
	}
	string3 = sign + string1 + "." + string2
	if flag==1{
		string3 = "division by 0 error"
	}
	if flag==2{
		string3 = "error gcd(10,b/gcd(a,b)) > 1"
	}
	return string3
}
