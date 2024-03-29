def isPalindrome(s: str) -> bool:
    s = cleanString(s)
    left, right = 0, len(s) - 1

    while left < right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1

    return True
    
def cleanString(s: str) -> str:
    return ''.join([c.lower() for c in s if isAlphaNumeric(c)])

def isAlphaNumeric(c: str) -> bool:
    return c.isalnum()

print(isPalindrome("A man, a plan, a canal: Panama"))
print(isPalindrome("race a car"))
print(isPalindrome(" "))
