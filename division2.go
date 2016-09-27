/* This program can be used to calculate value of a/b
to arbitrary precision for any two given integers a and b provided
that gcd(10,b/gcd(a,b)) = 1  If gcd(a,b) = 1 then this is
equivalent to requiring that gcd(10,b) = 1
There are three main functions in this program:
(1) divide(a,b,n1) which calculates a/b to n1 decimal digits of precision
(2) digits(a,b,n1,n2) which calculates the n1 to n2 digits in the fractional
part of the decimal expansion of a/b
(3) digits2(a,b,n1,n2) which calculates the n1 to n2 digits in the repeating
part of the fractional part of the decimal expansion of a/b.  This function
works even if gcd(10,b/gcd(a,b)) > 1.
The functions digits and digits2 are efficient even if n1 is a large integer
provided that n2 - n1 is reasonably small.  Other division algortihms such as
subtractive methods (digit recurrence methods) or multiplicative (iterative)
methods such as Newton-Raphson are not efficient or computationally feasible
for large values of n1.
The output of the divide and digits functions is returned as a string.
The output of digits2 is returned as the struct type dt.
See the comments in front of the definitions of both of the functions
divide and digits below for more information */

package main

import "fmt"
import "strconv"
import "reflect"
import "os"
import "os/exec"

type dt struct{
	c int
	digits string
}

func main() {
	var a,b,n1,n2,t1 int
	var string1 string
	var y dt

// The following line clears the screen in windows
// It should be deleted or commented out if a different os is being used
	clear()

	a1 := reflect.TypeOf(a).Size()
	fmt.Print("\n The size of int type is ",a1," bytes \n\n")

// Example 1
	a = 4342321
	b = 13213127
	n1 = 100
	string1 = divide(a,b,n1)
	fmt.Println(a,"/",b," to ",n1," decimal digits of precision is ")
	fmt.Print(string1,"\n\n")


// Example 2
	a = 1231237
	b = 42312319
	n1 = 5
	n2 = 25
	string1 = digits(a,b,n1,n2)
	fmt.Print("The ",n1," to ",n2," digits in the expansion of ",a," / ",b)
	fmt.Print(" is ",string1,"\n\n")

// Example 3
	a = 127
	t1 = 625
	b = t1*7
	n1 = 1
	n2 = 6
	y = digits2(a,b,n1,n2)
	fmt.Print("The first ",n2," digits in the expansion of the repeating")
	fmt.Println(" part of ",a,"/",b," is ",y.digits)
	fmt.Print("You would have to skip the first ",y.c," digits ")
	fmt.Println("to get to the part of the decimal expansion that repeats since")
	c1 := float64(a)
	c2 := float64(b)
	c := c1/c2
	fmt.Println(a,"/",b," = ",c,"\n\n")

	}

// This clears the screen in windows
// This function should be deleted or commented out
// if a different os is being used
func clear(){
	  cmd := exec.Command("cmd", "/c", "cls")
	  cmd.Stdout = os.Stdout
	  cmd.Run()
}

// This calculates the absolute value of the integer a
func abs(a int) int {
	var t int
	t = a
	if a<0{
		t = -a
	}
	return t
}

// This calculates g^e (mod n) for integers e,g, and n
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

// Primality test
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

/* Finds a prime p such that p = 1 (mod c) and p > c*m
If c is odd then the prime p = 3 (mod 4)
If 2^a divides c and 2^(a+1) does not divide c where a>0
then p = (2^a)+1 mod (2^(a+1)) */
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

// Calculates the greatest commmon divisor of a and b
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

// Recursive algorithm for calculating a^-1 (mod b) assuming gcd(a,b) = 1
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
	s = (s+(p1-1))/p
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


/* Given two integers a and b of type int, this calculates the
n1 to the n2 digits in the decimal expansion of the fractional part
of a/b which is |a|/|b| if |a| < |b| and is (|a| mod |b|)/|b| if |a| > |b|.
For example 1/17 = 0.058823529... so digits(1,17,1,20) = '05882352941176470588'
which is the first 20 digits in the decimal expansion of 1/17.  The output is
returned as a string.  This is based on a different division algorithm based on
exponentiation modulo b which is different from previous methods that are
referred to as either (1) subtractive methods such as trial division or
(2) multiplicative methods such as Newton-Raphson.  If n = n2 - n1 then this
algorithm runs O(n log n (log b)^2 + n (log b)^3) time.  Also it should be
noted that algorithm runs efficiently even if n1 is a very large integer.
For example suppose that b is a 100 digit integer and that n1 is a 1000 digit
integer then digits(a,b,n1,n1+n) could be quickly and efficiently calculated in
O(n log n1 (log b)^2 + n (log b)^3) time.  With previous subtractive or
multipicative division algorithms this would be computationally infeasible.
This algorithm can be improved to run in O(n log b + (log n1)(log b)^2) time.
Also this algorithm requires that gcd(10,b/gcd(a,b)) = 1
If gcd(a,b) = 1 then this is equivalent to requiring that gcd(10,b) = 1 */
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


/* Given two integers a and b of arbitrary size, this calculates the value
of a/b to n1 digits of decimal precision.  Most previously used
methods for division fall into one of two categories: (1) subtractive methods
or digit recurrence methods or (2) multiplicative methods such as
Newtwon-Raphson.  This is a different method based on exponentiation modulo b
and this algorithm runs in O(n log n (log b)^2 + n (log b)^3) time.  It can
be improved to run in O(n log b) time if a < b which in many cases is actually
faster than previously used methods.  The output is returned as a string.
Also this algorithm requires that gcd(10,b/gcd(a,b)) = 1
If gcd(a,b) = 1 then this is equivalent to requiring that gcd(10,b) = 1 */
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

/* digits2 can be used in the case that gcd(10, b/gcd(a,b)) > 1
since digits only works in the case that gcd(10, b/gcd(a,b)) = 1
If gcd(10, b/gcd(a,b)) = 1 then digits2 returns the output y
with y.string1 identical to the output of digits and with y.c = 0.
If gcd(10, b,gcd(a,b)) > 1 then y.c will equal how many digits
have to be skipped in the decimal expansion a/b to get to the first
digit of the repeating part and y.string1 will equal the n1 to n2 digits
in the repeating part of the decimal expansion of a/b
 */
func digits2(a, b, n1, n2 int) dt{
	var t,i,s,c,c1,c2,e int
	var flag byte
	var string1 string
	var y dt

	a = abs(a)
	b = abs(b)
	t = gcd(a,b)
	if t>1{
		a = a/t
		b = b/t
	}
	flag = 0
	if b==0{
		a = 0
		b = 1
		flag = 1
	}
	c1 = 0
	c2 = 0
	for i = 1; i < 100; i++{
		if (b%2)==0{
		c1 = c1+1
		b = b/2
		}
		if (b%2)!=0{
			break
		}
	}
	for i = 1; i < 100; i++{
		if (b%5)==0{
			c2 = c2+1
			b = b/5
		}
		if (b%5)!=0{
			break
		}
	}
		c = c1
		if c2>c1{
			c = c2
		}
		e = abs(c1-c2)
		if c1>c2{
			t = exp1(e,5,b)
			a = (a*t)%b
		}
		if c1<c2{
			t = exp1(e,2,b)
			a = (a*t)%b
		}
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
		y.c = c
		y.digits = string1
		return y

}
