def fractionToDecimal(numr, denr):
 
    # Initialize result
    res = ""
 
    # Create a map to store already seen
    # remainders. Remainder is used as key
    # and its position in result is stored
    # as value. Note that we need position
    # for cases like 1/6.  In this case,
    # the recurring sequence doesn't start
    # from first remainder.
    mp = {}
 
    # Find first remainder
    rem = numr % denr
 
    print(rem)
    # Keep finding remainder until either
    # remainder becomes 0 or repeats
    while ((rem != 0) and (rem not in mp)):
        print("before:", mp, rem, res)
        # Store this remainder
        mp[rem] = len(res)
 
        # Multiply remainder with 10
        rem = rem * 10
 
        # Append rem / denr to result
        res_part = rem // denr
        print("res_part: ", res_part)
        res += str(res_part)
 
        # Update remainder
        rem = rem % denr
        
        print("after:", mp, rem, res)
    
    print("outside:", mp, rem)
    if (rem == 0):
        return ""
    else:
        return res[mp[rem]:]
 
 
# Driver code
numr, denr = 4, 333
res = fractionToDecimal(numr, denr)
 
if (res == ""):
    print("No recurring sequence")
else:
    print("Recurring sequence is", res)