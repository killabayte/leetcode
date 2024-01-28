def canConstruct(self, ransomNote: str, magazine: str) -> bool:
    letterCount = {}

    for char in magazine:
        letterCount[char] += 1