class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:     
        if numerator == 0:
            return "0"
        
        sign = ""
        if (numerator < 0 and denominator > 0) or (numerator > 0 and denominator < 0):
            sign = "-"
        
        numerator = abs(numerator)
        denominator = abs(denominator)
        pre_dec = numerator // denominator
        pre_dec = str(pre_dec)
        res = sign + pre_dec
        
        remainder = numerator % denominator
        if remainder == 0:
            return res
        
        res += "."
        
        # LOGIC: For the post decimal part we are doing lon division.
        # We multiply by 10 and then get remainder after dividing
        # For repeating - if we reach a number we have already seen then it will start looping
        # That's why using the map to see if we have encountered that number before
        post_denom = ""
        map={}
        while (remainder != 0):
            if (remainder in map):
                # Only putting the repeating number in the brackets
                post_denom = post_denom[:map[remainder]] + "(" + post_denom[map[remainder]:] + ")"
                break
            map[remainder] = len(post_denom)
            
            remainder = remainder * 10
            
            res_part = remainder // denominator
            
            post_denom += str(res_part)
            remainder = remainder % denominator
        
        res += post_denom
        return res
    
sol = Solution()

test1 = [1,2]
test2 = [50,22]
test3 = [4,333]
test4 = [-50,8]

print(sol.fractionToDecimal(test1[0],test1[1]))
print(sol.fractionToDecimal(test2[0],test2[1]))
print(sol.fractionToDecimal(test3[0],test3[1]))
print(sol.fractionToDecimal(test4[0],test4[1]))