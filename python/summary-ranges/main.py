def summaryRanges(nums: List[int]) -> List[str]:
    result = []

    for i in range(len(nums)):
        start = nums[i]
        end = start
        