def canConstruct(ransomNote: str, magazine: str) -> bool:
    letterCount = {}

    for c in magazine:
        letterCount[c] = letterCount.get(c, 0) + 1

    for c in ransomNote:
        letterCount[c] = letterCount.get(c, 0) - 1
        if letterCount[c] < 0:
            return False
    return True

print(canConstruct("aabb", "aabbc"))