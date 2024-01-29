from typing import List


def summaryRanges(nums: List[int]) -> List[str]:
    result = []

    for i in range(len(nums)):
        start = nums[i]
        end = start

        while i < len(nums)-1 and nums[i+1] == nums[i]+1:
            end = nums[i+1]
            i += 1

        if start == end:
            result.append(str(start))
        else:
            result.append(str(start) + "->" + str(end))

    return result

print(summaryRanges([0,1,2,4,5,7]))