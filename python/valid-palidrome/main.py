def isPalindrome(s: str) -> bool:
    s = cleanString(s)
    left, right = 0, len(s) - 1
    
    
def cleanString(s: str) -> str:
    return ''.join([c.lower() for c in s if isAlphaNumeric(c)])

def isAlphaNumeric(c: str) -> bool:
    return c.isalnum()