""" This program can be used to calculate value of a/b
to arbitrary precision for any two given integers a and b provided
that gcd(10,b/gcd(a mod b,b)) = 1  If gcd(a mod b,b) = 1 then this
is equivalent to requiring that gcd(10,b) = 1
There are two main functions in this program:
(1) divide(a,b,n1) which calculates a/b to n1 decimal digits of precision
(2) digits(a,b,n1,n2) which calculates the n1 to n2 digits in the fractional
part of the decimal expansion of a/b
The function digits is efficient even if n1 is a large integer provided that
n2 - n1 is reasonably small.  Other division algortihms such as subtractive
methods or multiplicative (iterative) methods  
are not efficient or computationally feasible for large values of n1.  
The output of both functions is returned as a string.
See the comments in front of the definitions of both of the functions divide
and digits below for more information """

""" This calculates g^e (mod n) for integers e,g, and n """
def exp1(e,g,n):
    t = 1
    sq = g
    e1 = e
    while(e1!=0):
        if (e1%2)==1:
            t = (sq*t)%n
            e1 = (e1-1)//2
        else:
            e1 = e1//2
        sq = (sq*sq)%n
    return(t)

""" Primality test """
def is_prime(p):
    p = abs(p)
    if (p==0):
        p = 1
    m = 50
    y = 0
    b = False
    for i in range(2,m):
        t1 = i%p
        t = exp1((p-1)//2,i,p)
        if ((t==1) or (t==(p-1)) or (t1==0)):
            y = y + 1
    if (y==(m-2)):
        b = True
    if (p==1):
        b = False
    return(b)

""" Finds a prime p such that p = 1 (mod c) and p > c*m
If c is odd then the prime p = 3 (mod 4)
If 2^a divides c and 2^(a+1) does not divide c where a>0
then p = (2^a)+1 mod (2^(a+1)) """
def findprime(c,m):
    if (c<1):
        c = 1
    if (m<1):
        m = 1
    if (c%2==1):
        c = 2*c
        m = (m+(m%2))//2
    t = False
    i = m+(m+1)%2
    while(t==False):
        t1 = c*i+1
        t = is_prime(t1)
        i = i+2
    return(t1)

""" Calculates the greatest commmon divisor of a and b """
def gcd(a,b):
    a = abs(a)
    b = abs(b)
    if (a==0 or b==0):
        s = a+b
        return(s)
    if (a==1 or b==1):
        return(1)
    s = gcd(b%a,a)
    return(s)

""" Recursive algorithm for calculating a^-1 (mod b) assuming gcd(a,b) = 1"""
def inverse2(a,b):
    if (a==0 or b==0):
        return(0)
    if (a==1):
        return(1)
    else:
        t = (b-((b*inverse2(b%a,a))//a))
        return(t)

def f10v3(n,a,p,q):
    p1 = p*q
    g = 10
    a1 = inverse2(a,p1)
    n1 = abs(n)
    t = exp1(n1,g,p1)
    if (n<0):
        t = inverse2(t,p1)
    t1 = a1*t%p
    t2 = inverse2(t1,p)
    s = t2
    s = (s*a*t-1)//p
    s = s%q
    return(s)

def f10v4(n,a,p):
    n = -n
    t = (11,13,17,19,23,29,31,37,41,43,47)
    q = p
    for q1 in t:
        if a%q1!=0:
            q = q1
            break
    q = min(q,p)
    if (q<11):
        q = 11
    s1 = f10v3(n,a,p,q)
    s2 = f10v3(n+1,a,p,q)
    s1 = (q-s1)%q
    s = (s1+s2)%q
    n1 = abs(n)
    g = 10
    if (n>0):
        g = inverse2(10,q)        
    s1 = exp1(n1,g,q) 
    s = s*s1%q
    a1 = inverse2(a,q)
    s = a1*s%q
    return(s)

""" Given two integers a and b of arbitrary size, this calculates the
n1 to the n2 digits in the base 10 decimal expansion of the fractional part
of a/b which is |a|/|b| if |a| < |b| and is (|a| mod |b|)/|b| if |a| > |b|.
For example 1/17 = 0.058823529... so digits(1,17,1,20) = '05882352941176470588'
which is the first 20 digits in the decimal expansion of 1/17.  The output is
returned as a string.  This is based on a different division algorithm based on
exponentiation modulo b which is different from previous methods that are
referred to as either (1) subtractive methods or digit recurrence methods or
(2) multiplicative methods.  If n = n2 - n1 then this
algorithm runs O(n log n (log b)^2 + n (log b)^3) time.  Also it should be
noted that algorithm runs efficiently even if n1 is a very large integer.
For example suppose that b is a 100 digit integer and that n1 is a 1000 digit
integer then digits(a,b,n1,n1+n) could be quickly and efficiently calculated in
O(n log n1 (log b)^2 + n (log b)^3) time.  With previous subtractive or
multipicative division algorithms this would be computationally infeasible.
This algorithm can be improved to run in O(n log b + (log n1)(log b)^2) time.
Also this algorithm requires that gcd(10,b/gcd(a,b)) = 1
If gcd(a,b) = 1 then this is equivalent to requiring that gcd(10,b) = 1 """
def digits(a,b,n1,n2):
    a = int(a)
    b = int(b)
    n1 = int(n1)
    n2 = int(n2)
    t = gcd(a,b)
    if (t>1):
        a = a//t
        b = b//t
    flag = 0
    if (b!=0) and ((b%2==0) or (b%5==0)):
        flag = 2
    if (b==0):
        a = 0
        b = 1
        flag = 1
    a = abs(a)
    b = abs(b)
    if (a>b):
        a = a%b
    string1 = ""
    for i in range(n1,n2+1):
        s = f10v4(i,a,b)
        string1 = string1+str(s)
    if (flag==1):
        string1 = "division by 0 error"
    if (flag==2):
        string1 = "error gcd(10,b/gcd(a,b)) > 1"
    return(string1)

""" Given two integers a and b of arbitrary size, this calculates the value
of a/b to n1 digits of decimal precision in base 10.  Most previously used
methods for division fall into one of two categories: (1) subtractive methods
or digit recurrence methods or (2) multiplicative methods such as
Newtwon-Raphson.  This is a different method based on exponentiation modulo b
and this algorithm runs in O(n log n (log b)^2 + n (log b)^3) time.  It can
be improved to run in O(n log b) time if a < b which in many cases is actually
faster than previously used methods.  The output is returned as a string.
Also this algorithm requires that gcd(10,b/gcd(a,b)) = 1
If gcd(a,b) = 1 then this is equivalent to requiring that gcd(10,b) = 1 """
def divide(a,b,n1):
    a = int(a)
    b = int(b)
    n1 = int(n1)
    t = gcd(a,b)
    if (t>1):
        a = a//t
        b = b//t
    flag = 0
    if (b!=0) and ((b%2==0) or (b%5==0)):
        flag = 2
    if (b==0):
        a = 0
        b = 1
        flag = 1
    sign = ""
    if (a<0) and (b>0):
        sign = "-"
    if (a>0) and (b<0):
        sign = "-"
    a = abs(int(a))
    b = abs(int(b))
    s = 0
    if (a>b):
        s = a//b
        a = a%b
    string1 = str(s)
    string2 = ""
    for i in range(1,n1+1):
        s = f10v4(i,a,b)
        string2 = string2+str(s)
    s = sign+string1+"."+string2
    if (flag==1):
        s = "division by 0 error"
    if (flag==2):
        s = "error gcd(10,b/gcd(a,b)) > 1"
    return(s)


""" Examples """

""" Example 1 """
a = 34
b = 89
c = a/b
print("\n\nExample 1")
print("\nThe following is an example of division using Python's / operator")
print("\n            {0}/{1} = {2}".format(a,b,c))

print("\n\nThis is an example of division using the function divide(a,b,n)")
print("to calculate {0}/{1} to a precision of 50 decimal digits".format(a,b))
s = divide(a,b,50)
print("\n divide({0},{1},50) = {2}".format(a,b,s))
print("\n\n")

""" Example 2 """
a = 113253454212386751
b = 65423454543432347
n = 100
print("\n\nExample 2")
print("\nHere is a larger example of using the divide function to")
print("calculate {0} / {1} to {2} digits of precision".format(a,b,n))
s = divide(a,b,n)
print("{0} / {1} = {2}".format(a,b,s))
print("\n\n")

""" Example 3 """
a = 12
p = 17
n1 = p-1
print("\n\nExample 3")
s1 = divide(a,p,n1)
print("\ndivide({0},{1},{2}) = {3}".format(a,p,n1,s1))
s1 = digits(a,p,1,10)
s2 = digits(a,p,-9,0)
s3 = digits(a,p,6,11)
print("\nUsing the function digits the first 10 digits of the decimal")
print("expansion of {0} / {1} are {2}".format(a,p,s1),end=" ")
print("since digits({0},{1},1,10) = {2}".format(a,p,s1))
print("The last 10 digits of the decimal expansion are",s2,end=" ")
print("\nsince digits({0},{1},-9,0) = {2}".format(a,p,s2))
print("The 6th through the 11th digits are",s3,end=" ")
print("since digits({0},{1},6,11) = {2}".format(a,p,s3))
print("\n\n")

""" Example 4 """
a = 2**520+2**519+2**401+123543525356364532132432112223267
p = 2**521-1
n1 = 120
print("\n\nExample 4")
s1 = divide(a,p,n1)
print()
print("Here is an example of using the divide function to calculate",end="")
print(" a / p to",n1,"decimal digits of accuracy where")
print("a =",a)
print("p =",p)
print("\ndivide(a,p,{0}) = {1}".format(n1,s1))
n = 2**256+1432432432
n1 = n+10
s2 = digits(a,p,n,n1)
print("\nIf n =",n)
print("then the 11 digits starting from the nth digit",end="")
print(" to the (n+10)th digit are")
print("\ndigits(a,p,n,n+10) = {0}".format(s2))
print("\nNotice that the previous function digits was calculated efficiently")
print("in polynomial time even though n is a very large integer")
print("(n is approximately 2^256) and this would be computationally")
print("infeasible with the two other commonly used division methods usually")
print("referred to as (1) subtractive methods or digit recurrence methods or")
print("(2) multiplicative methods such as Newton-Raphson")
