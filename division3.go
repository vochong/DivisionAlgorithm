/* This program can be used to calculate the n1 to n2 digits in the
decimal expansion of a/b for any two given integers a and b.
The function digits works provided that gcd(10,b/gcd(a,b)) = 1
If gcd(a,b) = 1 then this is equivalent to requiring that gcd(10,b) = 1
There are three main functions in this program:
(1) digits(a,b,n1,n2) which calculates the n1 to n2 digits in the fractional
part of the decimal expansion of a/b using a new and faster method based on
exponentiation modulo b.
(2) digits2(a,b,n1,n2) which calculates the n1 to n2 digits in the repeating
part of the fractional part of the decimal expansion of a/b.  This function
works even if gcd(10,b/gcd(a,b)) > 1.  The function digits2 is based on a new
and faster method that is based on exponetiation modulo b.
(3) digits3(a,b,n1,n2) calculates the n1 to n2 digits in the fractional part
of the decimal expansion of a/b using the standard method based on division.
The functions digits and digits2 are efficient even if n1 is a large integer
provided that n2 - n1 is reasonably small.  Other division algortihms such as
subtractive methods (digit recurrence methods) or multiplicative (iterative)
methods such as Newton-Raphson are not efficient or computationally feasible
for large values of n1.  digits3 is a much less efficient version (for large
values of n1) based on division and is given and is given here to compare how
much faster the functions digits and digits2 are.
The output of the digits and digits3 functions is returned as a string.
The output of digits2 is returned as the struct type dt2.
See the comments in front of the definitions of both of the functions
divide and digits below for more information */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"math/big"
	"strings"
	"time"
)


type dt2 struct{
	c int64
	digits string
}

func main() {
	var n1,n2 uint
	var string1 string
	var y dt2
	a := big.NewInt(0)
	b := big.NewInt(0)

// The following line clears the screen in windows
// It should be deleted or commented out if a different os is being used

	clear()

	// Example 1
	time1 := time.Now()
	prec := uint(4*100)
	a1 := "121211"
	b1 := "1233211"
	a.SetString(a1,10)
	b.SetString(b1,10)
	d2 := new(big.Int).SetInt64(90)
	d3 := new(big.Int).SetInt64(100)
	fmt.Printf("\n\n Example 1  a = %d  b = %d",a,b)

	t4 := new(big.Float).SetPrec(prec)
	t5 := new(big.Float).SetPrec(prec)
	t6 := new(big.Float).SetPrec(prec)
	t4.SetString(a1)
	t5.SetString(b1)
	t6.Quo(t4,t5)
	fmt.Printf("\n\n a/b = %.100f\n",t6)

	string1 = digits(a,b,d2,d3)
	fmt.Printf("\n digits(%d, %d, %d, %d) = %s",a,b,d2,d3,string1)
	time2 := time.Now()
	time3 := time2.Sub(time1)
	fmt.Println("\n\n time to calculate example 1 is",time3)


	// Example 2
	time1 = time.Now()
	prec = uint(4*100)
	a1 = "121211126"
	b1 = "123321125"
	a.SetString(a1,10)
	b.SetString(b1,10)
	d2.SetInt64(1)
	d3.SetInt64(20)
	fmt.Printf("\n\n\n\n Example 2  a = %d  b = %d",a,b)

	t4.SetPrec(prec)
	t5.SetPrec(prec)
	t6.SetPrec(prec)
	t4.SetString(a1)
	t5.SetString(b1)
	t6.Quo(t4,t5)
	fmt.Printf("\n\n a/b = %.100f\n",t6)

	y = digits2(a,b,d2,d3)
	fmt.Printf("\n digits2(%d, %d, %d, %d) = (%d, %s)",a,b,d2,d3,y.c,y.digits)
	time2 = time.Now()
	time3 = time2.Sub(time1)
	fmt.Println("\n\n time to calculate example 2 is",time3)


	// Example 3
	time1 = time.Now()
	a1 = "12111243767123489121487"
	b1 = "341232378975123713776517"
	a.SetString(a1,10)
	b.SetString(b1,10)
	n1 = 30000
	n2 = n1+80
	fmt.Printf("\n\n\n\n")
	fmt.Printf(" Example 3  a = %d  b = %d  n1 = %d  n2 = %d\n\n",a,b,n1,n2)
	string1 = digits3(a,b,n1,n2)
	fmt.Println(" digits3(a,b,n1,n2) = ",string1)
	time2 = time.Now()
	time3 = time2.Sub(time1)
	fmt.Println("\n time to calculate example 3 is",time3)


	// Example 4
	time1 = time.Now()
	a.SetString(a1,10)
	b.SetString(b1,10)
	d2.SetString("30000",10)
	d3.SetString("30080",10)
	fmt.Printf("\n\n\n\n")
	fmt.Printf(" Example 4  a = %d  b = %d  n1 = %d  n2 = %d \n\n",a,b,n1,n2)
	string1 = digits(a,b,d2,d3)
	fmt.Println(" digits(a,b,n1,n2) = ",string1)
	time2 = time.Now()
	time3 = time2.Sub(time1)
	fmt.Println("\n time to calculate example 4 is",time3,"\n\n")

	}


// This clears the screen in windows
// This function should be deleted or commented out
// if a different os is being used
func clear(){
	  cmd := exec.Command("cmd", "/c", "cls")
	  cmd.Stdout = os.Stdout
	  cmd.Run()
}

func f10v3(n, a, p, q *big.Int) *big.Int{
	g := big.NewInt(10)
	p1 := big.NewInt(0)
	a1 := big.NewInt(0)
	n1 := big.NewInt(0)
	t := big.NewInt(0)
	t1 := big.NewInt(0)
	t2 := big.NewInt(0)
	t3 := big.NewInt(1)
	s := big.NewInt(0)

	p1.Mul(p,q)
	a1.ModInverse(a,p1)
	n1.Abs(n)
	t.Exp(g,n1,p1)
	t2.SetString("0",10)
	if n.Cmp(t2)==-1{
		t.ModInverse(t,p1)
	}
	t1.Mod(t1.Mul(a1,t),p)
	s.ModInverse(t1,p)
	s.Mod(s.Mul(s,a),p1)
	s.Mod(s.Mul(s,t),p1)
	t2.Sub(p1,t3)
	s.Div(s.Add(s,t2),p)
	s.Mod(s,q)
	return s
}

func f10v4(n, a, p *big.Int) string{
	var i int
	t := [11]int64 {11,13,17,19,23,29,31,37,41,43,47}
	q := big.NewInt(0)
	q1 := big.NewInt(0)
	s1 := big.NewInt(0)
	s2 := big.NewInt(0)
	s := big.NewInt(0)
	g := big.NewInt(10)
	n1 := big.NewInt(0)
	n2 := big.NewInt(0)
	a1 := big.NewInt(0)

	n2.Neg(n)
	q.Set(p)
	for i=0; i<11; i++ {
		q1.SetInt64(t[i])
		s1.Mod(a,q1)
		if s1.Cmp(s2)!=0{
			q.Set(q1)
			break
		}
	}
	if p.Cmp(q)==-1{
		q.Set(p)
	}
	s2.SetInt64(11)
	if q.Cmp(s2)==-1{
		q.SetInt64(11)
	}
	s1 = f10v3(n2,a,p,q)
	s2.SetInt64(1)
	s.Add(n2,s2)
	s2 = f10v3(s,a,p,q)
	s1.Sub(q,s1)
	s1.Mod(s1,q)
	s.Add(s1,s2)
	s.Mod(s,q)
	n1.Abs(n2)
	s2.SetInt64(0)
	if n.Cmp(s2)==-1{
		g.ModInverse(g,q)
	}
	s1.Exp(g,n1,q)
	s.Mod(s.Mul(s,s1),q)
	a1.ModInverse(a,q)
	s.Mod(s.Mul(a1,s),q)
  	t1 := s.String()
	return t1
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

func digits(a, b, n1, n2 *big.Int) string{
  var flag byte
  var string1 string
  i := big.NewInt(0)
  t := big.NewInt(0)
  s := big.NewInt(0)
  s1 := big.NewInt(0)
  s2 := big.NewInt(0)
  a1 := big.NewInt(0)
  b1 := big.NewInt(0)

  a1.Set(a)
  b1.Set(b)
  a1.Abs(a1)
  b1.Abs(b1)
  t.GCD(nil,nil,a1,b1)
  s.SetInt64(1)
  if t.Cmp(s)==1{
    a1.Div(a1,t)
    b1.Div(b1,t)
  }
  s.SetInt64(2)
  s1.Mod(b1,s)
  s.SetInt64(5)
  s2.Mod(b1,s)
  s.SetInt64(0)
  flag = 0
  if (b1.Cmp(s)!=0) && (s1.Cmp(s)==0 || s2.Cmp(s)==0){
    flag = 2
    }
  if (b1.Cmp(s)==0){
    a1.SetInt64(0)
    b1.SetInt64(1)
    flag = 1
  }
  if (a1.Cmp(b1)==1){
    a1.Mod(a1,b1)
  }
  string1 = ""
	i.Set(n1)
	s1.SetInt64(1)
	for ; i.Cmp(n2)!=1 ; {
		string1 = string1 + f10v4(i,a1,b1)
		i.Add(i,s1)
	}
	if flag==1{
		string1 = "division by 0 error"
	}
	if flag==2{
		string1 = "error gcd(10,b/gcd(a,b)) > 1"
	}
	return string1
}


/* digits2 can be used in the case that gcd(10, b/gcd(a,b)) > 1
since digits only works in the case that gcd(10, b/gcd(a,b)) = 1
If gcd(10, b/gcd(a,b)) = 1 then digits2 returns the output y
with y.string1 identical to the output of digits and with y.c = 0.
If gcd(10, b,gcd(a,b)) > 1 then y.c will equal how many digits
have to be skipped in the decimal expansion a/b to get to the first
digit of the repeating part and y.string1 will equal the n1 to n2 digits
in the repeating part of the decimal expansion of a/b  */

func digits2(a, b, n1, n2 *big.Int) dt2{
	var i1,c,c1,c2,e int64
	var flag byte
	var string1 string
	var y dt2
	i := big.NewInt(0)
  	t := big.NewInt(0)
	s := big.NewInt(0)
  	s1 := big.NewInt(0)
  	s2 := big.NewInt(0)
	a1 := big.NewInt(0)
	b1 := big.NewInt(0)

	a1.Set(a)
	b1.Set(b)
	a1.Abs(a1)
	b1.Abs(b1)
	t.GCD(nil,nil,a1,b1)
	s.SetInt64(1)
	if t.Cmp(s)==1{
	  a1.Div(a1,t)
	  b1.Div(b1,t)
	}
	if (b1.Cmp(s)==0){
	  a1.SetInt64(0)
	  b1.SetInt64(1)
	  flag = 1
	}
	if (a1.Cmp(b1)==1){
	  a1.Mod(a1,b1)
	}
	c1 = 0
	c2 = 0
	s.SetInt64(0)
	s1.SetInt64(2)
	for i1 = 1; i1 < 100; i1++{
		s2.Mod(b1,s1)
		if s2.Cmp(s)==0{
			c1 = c1+1
			b1.Div(b1,s1)
		}
		s2.Mod(b1,s1)
		if s2.Cmp(s)!=0{
			break
		}
}
s1.SetInt64(5)
for i1 = 1; i1 < 100; i1++{
	s2.Mod(b1,s1)
	if s2.Cmp(s)==0{
		c2 = c2+1
		b1.Div(b1,s1)
	}
	s2.Mod(b1,s1)
	if s2.Cmp(s)!=0{
		break
	}
}
c = c1
if c2>c1{
	c = c2
}
e = 0
if c1>c2{
	e = c1-c2
}
if c1<c2{
	e = c2-c1
}
s2.SetInt64(e)
if c1>c2{
	s1.SetInt64(5)
	t.Exp(s1,s2,b1)
	a1.Mul(a1,t)
	a1.Mod(a1,b1)
}
if c1<c2{
	s1.SetInt64(2)
	t.Exp(s1,s2,b1)
	a1.Mul(a1,t)
	a1.Mod(a1,b1)
}

string1 = ""
i.Set(n1)
s1.SetInt64(1)
for ; i.Cmp(n2)!=1 ; {
	string1 = string1 + f10v4(i,a1,b1)
	i.Add(i,s1)
}
if flag==1{
	string1 = "division by 0 error"
}
y.c = c
y.digits = string1
return y
}


/* digits3 can be used to calculate the n1 to n2 digits in the decimal
expansion of a/b.  This algorithm can be considered to be the standard
method for doing this based on large floating point arithmetic.
Unlike digits or digits2, it requires that n1 and n2 both be positive integers.
For large values of n1 it is much less efficient than digits or digits2 since
all decimal digits from 1 to n2 have to be calculated in the decimal expansion
of a/b unlike digits and digits2 which only have to calculate the n1 to n2
digits.  This algorithm is given here to compare how much faster the
functions digits and digits2 are. */

func digits3(a, b *big.Int, n1, n2 uint) string{
	var string1 string
	prec := uint(4*n2)
	t1 := new(big.Float).SetPrec(prec)
	t2 := new(big.Float).SetPrec(prec)
	t3 := new(big.Float).SetPrec(prec)
	a1 := big.NewInt(0)
	b1 := big.NewInt(0)

	a1.Set(a)
	b1.Set(b)
	a1.Abs(a1)
	b1.Abs(b1)

	if (a1.Cmp(b1)==1){
	  a1.Mod(a1,b1)
	}

	string1 = a1.String()
	t1.SetString(string1)
	string1 = b1.String()
	t2.SetString(string1)

	if n1<1{
		n1 = 1
	}

	t3.Quo(t1, t2)
	string1 = t3.Text('f',int(prec))
	s1 := strings.Split(string1, ".")
	string1 = s1[1]
	string1 = string1[int(n1-1):int(n2)]

	return string1
}
