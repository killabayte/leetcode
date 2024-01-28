def canConstruct(self, ransomNote: str, magazine: str) -> bool:
    letterCount = {}

    for char in magazine:
        letterCount[char] += 1

    for char in ransomNote:
        letterCount[char] -= 1
        if letterCount[char] < 0:
            return False