# Generate valid IP addresses from given digits

## Daily Coding Problem: Problem #1085 [Medium]

This problem was asked by Snapchat.

Given a string of digits, generate all possible valid IP address combinations.

IP addresses must follow the format A.B.C.D,
where A, B, C, and D are numbers between 0 and 255.
Zero-prefixed numbers, such as 01 and 065, are not allowed, except for 0 itself.

For example, given "2542540123", you should return ['254.25.40.123', '254.254.0.123'].

## Analysis

The problem statement is too loose.
The example clarifies that you leave the string of digit characters
in the order given.

I treated this as a
"generate all plausible strings, then display the valid strings"
problem.

Go's easy concurrency lends itself to generating answers
that might fit the requirements,
then sorting out the valid answers.
I have a goroutine running a recursive semi-plausible IP-address-construction
algorithm, then writing all the semi-plausible IP-addresses to a channel.
The main goroutine reads from the channel,
runs a validation function on the semi-plausible addresses,
printing only the strings that fit the definition of "dotted quad IP address".

The semi-plausible IP-address generation is facilitated by using the input
digits in the order given.
Validating the semi-plausible IP-address candidates is harder than generating them.

You could do this by backtracking, too.
The same 
